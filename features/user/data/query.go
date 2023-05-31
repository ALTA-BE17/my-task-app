package data

import (
	"errors"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/dependency"
	bookCore "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/book"
	book "github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/book/data"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/helper"
	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/middlewares"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Query struct {
	dep dependency.Dependency
}

func New(dep dependency.Dependency) user.UserData {
	return &Query{dep: dep}
}

func (q *Query) Register(request user.Core) (user.Core, error) {
	hashed, err := helper.HashPassword(request.Password)
	if err != nil {
		q.dep.Logger.Warn("error while hashing password", zap.Error(err))
		return user.Core{}, errors.New("password processing error")
	}
	createdUser := user.Core{}
	request.Password = hashed
	req := userEntities(request)
	query := q.dep.DB.Table("users").Create(&req)
	if query.Error != nil {
		q.dep.Logger.Warn("error insert data, duplicated", zap.Error(query.Error))
		return user.Core{}, errors.New("error insert data, duplicated")
	}

	rowAffect := query.RowsAffected
	if rowAffect == 0 {
		q.dep.Logger.Warn("no user has been created")
		return user.Core{}, errors.New("insert failed, row affected = 0")
	}

	q.dep.Logger.Sugar().Infof("new user has been created: %s, %s", createdUser.Name, createdUser.Email)
	return createdUser, nil
}

func (q *Query) Login(request user.Core) (user.Core, string, error) {
	result := User{}
	query := q.dep.DB.Where("name = ?", request.Name).First(&result)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		q.dep.Logger.Info("Username not found", zap.Error(query.Error))
		return user.Core{}, "", errors.New("invalid username or password")
	}

	rowAffect := query.RowsAffected
	if rowAffect == 0 {
		q.dep.Logger.Warn("no user has been created")
		return user.Core{}, "", errors.New("insert failed, row affected = 0")
	}

	if !helper.MatchPassword(request.Password, result.Password) {
		q.dep.Logger.Info("password does not match")
		return user.Core{}, "", errors.New("password does not match")
	}

	token, err := middlewares.CreateToken(int(result.ID))
	if err != nil {
		return user.Core{}, "", errors.New("internal server error")
	}

	q.dep.Logger.Sugar().Infof("user login: %s", result.Password)
	q.dep.Logger.Sugar().Infof("user login: %s, %s", result.Name, result.Email)
	return userModels(result), token, nil
}

func (q *Query) Profile(userId uint) (user.Core, error) {
	userModel := User{}
	query := q.dep.DB.Raw("SELECT * FROM users WHERE id = ? AND deleted_at IS NULL", userId).Scan(&userModel)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		q.dep.Logger.Sugar().Infof("user id : %d not found", userId)
		return user.Core{}, errors.New("error while retrieving user profile")
	}

	q.dep.Logger.Sugar().Infof("user has been retrieved profile: %s, %s", userModel.Name, userModel.Email)
	return userModels(userModel), nil
}

func (q *Query) SearchUsers(userId uint, pattern string) ([]user.Core, error) {
	users := []User{}
	query := q.dep.DB.Raw("SELECT * FROM users WHERE name OR email LIKE ? AND deleted_at IS NULL;", "%"+pattern+"%").Scan(&users)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		q.dep.Logger.Sugar().Errorf("users with quote %s not found.", pattern)
		return nil, errors.New("not found, error while retrieving list users")
	}

	res := make([]user.Core, len(users))
	for i, user := range users {
		res[i] = userModels(user)
	}

	user, _ := q.Profile(userId)
	count := len(res)

	q.dep.Logger.Sugar().Infof("Username : %s searching for total users: %d of %s", user.Name, count, pattern)
	return res, nil
}

func (q *Query) UpdateProfile(userId uint, request user.Core) (user.Core, error) {
	userModel := userEntities(request)
	query := q.dep.DB.Table("users").Where("id = ? AND deleted_at IS NULL", userId).Updates(&userModel)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		q.dep.Logger.Sugar().Infof("user id : %d not found", userId)
		return user.Core{}, errors.New("error while retrieving user profile")
	}

	affrows := query.RowsAffected
	if affrows == 0 {
		q.dep.Logger.Sugar().Infoln("failed to update user")
		return user.Core{}, errors.New("failed to update user, row affected = 0")
	}

	err := query.Error
	if err != nil {
		q.dep.Logger.Sugar().Infoln("update user query error, duplicate data entry")
		return user.Core{}, errors.New("duplicate data entry")
	}

	user, _ := q.Profile(userId)
	q.dep.Logger.Sugar().Infof("user: %s, %s has been updated a profile", user.Name, user.Email)
	return userModels(userModel), nil
}

func (q *Query) Deactive(userId uint) (user.Core, error) {
	userModel := User{}
	query := q.dep.DB.Table("users").Where("id = ? AND deleted_at IS NULL", userId).Delete(&userModel)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		q.dep.Logger.Sugar().Infof("user id : %d not found", userId)
		return user.Core{}, errors.New("error while retrieving user profile")
	}

	affrows := query.RowsAffected
	if affrows == 0 {
		q.dep.Logger.Sugar().Infoln("delete user query error")
		return user.Core{}, errors.New("insert failed, row affected = 0")
	}

	err := query.Error
	if err != nil {
		q.dep.Logger.Sugar().Infoln("update user query error, duplicate data entry")
		return user.Core{}, errors.New("duplicate data entry")
	}

	q.dep.Logger.Sugar().Infof("user has been deleted a profile")
	return userModels(userModel), nil
}

func (q *Query) GetAllUserHasBooks() ([]user.Core, error) {
	users := []User{}
	query := q.dep.DB.Preload("Books").Find(&users)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		q.dep.Logger.Sugar().Infof("list users not found")
		return []user.Core{}, errors.New("internal server error, list users not found")
	}

	userHasBooks := make(map[uint][]book.Book)
	for _, user := range users {
		// user.Books... digunakan untuk menggabungkan semua elemen dalam slice user.Books
		// menjadi argumen yang bisa di-append ke slice yang ada pada userBooksMap[user.ID].
		userHasBooks[user.ID] = append(userHasBooks[user.ID], user.Books...)

		// for _, book := range user.Books {
		// 	userBooksMap[user.ID] = append(userBooksMap[user.ID], book)
		// }
	}

	res := make([]user.Core, len(users))
	for i, user := range users {
		result := userModels(user)
		result.Books = make([]bookCore.Core, len(userHasBooks[user.ID]))
		for j, book := range userHasBooks[user.ID] {
			result.Books[j] = bookModels(book)
		}
		res[i] = result
	}

	return res, nil
}
