package data

import (
	"time"

	"github.com/ALTA-BE17/Rest-API-Clean-Arch-Test/features/user"
	"gorm.io/gorm"
)

type User struct {
	UserID    uint `gorm:"primaryKey; autoIncrement"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Username  string         `gorm:"type:VARCHAR(100);not null;unique"`
	Phone     string         `gorm:"type:VARCHAR(15);not null;unique"`
	Email     string         `gorm:"type:VARCHAR(100);not null;unique"`
	Password  string         `gorm:"type:VARCHAR(225);not null"`
}

// User-model to user-core
func userModels(u User) user.Core {
	return user.Core{
		UserID:    u.UserID,
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
	return User{
		UserID:    u.UserID,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		Username:  u.Username,
		Phone:     u.Phone,
		Email:     u.Email,
		Password:  u.Password,
	}
}
