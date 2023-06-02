package data

import (
	"time"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID    string `gorm:"type:varchar(100);primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"type:varchar(100);not null;unique"`
	Phone     string         `gorm:"type:varchar(15);not null;unique"`
	Email     string         `gorm:"type:varchar(100);not null;unique"`
	Password  string         `gorm:"type:varchar(225);not null"`
}

// User-model to user-core
func userModels(u User) user.Core {
	User_ID, _ := uuid.NewUUID()
	return user.Core{
		UserID:    User_ID.String(),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Username:  u.Username,
		Phone:     u.Phone,
		Email:     u.Email,
		Password:  u.Password,
	}
}

// User-core to user-model
func userEntities(u user.Core) User {
	User_ID, _ := uuid.NewUUID()
	return User{
		UserID:    User_ID.String(),
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Username:  u.Username,
		Phone:     u.Phone,
		Email:     u.Email,
		Password:  u.Password,
	}
}
