package data

import (
	"errors"
	"log"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/app/dependency"
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
	result := User{}
	hashed, err := helper.HashPassword(request.Password)
	if err != nil {
		q.dep.Logger.Warn("error while hashing password", zap.Error(err))
		return user.Core{}, errors.New("password processing error")
	}

	request.Password = hashed
	req := userEntities(request)

	query := q.dep.DB.Table("users").Create(&req)
	if query.Error != nil {
		q.dep.Logger.Warn("error insert data, duplicated", zap.Error(query.Error))
		return user.Core{}, errors.New("error insert data, duplicated")
	}

	rowAffect := query.RowsAffected
	if rowAffect == 0 {
		q.dep.Logger.Warn("no user has been created", zap.Error(query.Error))
		return user.Core{}, errors.New("insert failed, row affected = 0")
	}

	q.dep.Logger.Sugar().Infof("new user has been created")
	return userModels(result), nil
}

func (q *Query) Login(request user.Core) (user.Core, string, error) {
	result := User{}
	query := q.dep.DB.Where("username = ?", request.Username).First(&result)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		q.dep.Logger.Warn("user not found", zap.Error(query.Error))
		return user.Core{}, "", errors.New("invalid username or password")
	}

	rowAffect := query.RowsAffected
	if rowAffect == 0 {
		q.dep.Logger.Warn("no user has been created", zap.Error(query.Error))
		return user.Core{}, "", errors.New("insert failed, row affected = 0")
	}

	if !helper.MatchPassword(request.Password, result.Password) {
		q.dep.Logger.Warn("password does not match")
		return user.Core{}, "", errors.New("password does not match")
	}

	token, err := middlewares.CreateToken((result.UserID))
	if err != nil {
		q.dep.Logger.Warn("error while creating jwt token", zap.Error(err))
		return user.Core{}, "", errors.New("internal server error")
	}

	q.dep.Logger.Sugar().Infof("user login: %s, %s", result.Username, result.Email)
	return userModels(result), token, nil
}

func (q *Query) Profile(userId string) (user.Core, error) {
	userModel := User{}
	query := q.dep.DB.Raw("SELECT * FROM users WHERE user_id = ? AND deleted_at IS NULL", userId).Scan(&userModel)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		q.dep.Logger.Warn("user not found", zap.Error(query.Error))
		return user.Core{}, errors.New("error while retrieving user profile")
	}

	q.dep.Logger.Sugar().Infof("user has been retrieved profile: %s, %s", userModel.Username, userModel.Email)
	return userModels(userModel), nil
}

func (q *Query) SearchUsers(userId string, pattern string) ([]user.Core, error) {
	users := []User{}
	query := q.dep.DB.Raw("SELECT * FROM users WHERE username OR email LIKE ? AND deleted_at IS NULL;", "%"+pattern+"%").Scan(&users)
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

	q.dep.Logger.Sugar().Infof("Username : %s searching for total users: %d of %s", user.Username, count, pattern)
	return res, nil
}

func (q *Query) UpdateProfile(userId string, request user.Core) (user.Core, error) {
	result := User{}
	hashed, errHash := helper.HashPassword(request.Password)
	if errHash != nil {
		q.dep.Logger.Warn("error while hashing password", zap.Error(errHash))
		return user.Core{}, errors.New("password processing error")
	}

	request.Password = hashed
	req := userEntities(request)

	log.Printf("success hashing new password: %s", hashed)

	query := q.dep.DB.Table("users").Where("user_id = ? AND deleted_at IS NULL", userId).Updates(&req)
	if errors.Is(query.Error, gorm.ErrRecordNotFound) {
		q.dep.Logger.Sugar().Infof("user id : %d not found", userId)
		return user.Core{}, errors.New("error while retrieving user profile")
	}

	affrows := query.RowsAffected
	if affrows == 0 {
		q.dep.Logger.Sugar().Infoln("failed to update user")
		return user.Core{}, errors.New("failed to update user, row affected = 0")
	}

	err1 := query.Error
	if err1 != nil {
		q.dep.Logger.Sugar().Infoln("update user query error, duplicate data entry")
		return user.Core{}, errors.New("duplicate data entry")
	}

	user, _ := q.Profile(userId)
	q.dep.Logger.Sugar().Infof("user: %s, %s has been updated a profile", user.Username, user.Email)
	return userModels(result), nil
}

func (q *Query) Deactive(userId string) (user.Core, error) {
	userModel := User{}
	query := q.dep.DB.Table("users").Where("user_id = ? AND deleted_at IS NULL", userId).Delete(&userModel)
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
