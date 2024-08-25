package domain

import (
	"time"
)

type User struct {
	ID         int       `json:"id"`
	Email      string    `json:"email" validate:"email"`
	Password   string    `json:"password" validate:"min=8,max=20"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type IUserUsecase interface {
	Login(user *User) (jwt string, err error)
	Register(user *User) error
}

type IUserRepository interface {
	CheckDuplicateEmail(email string) (bool, error)
	GetUserByEmail(email string) (User, error)
	CreateUser(user *User) error
}
