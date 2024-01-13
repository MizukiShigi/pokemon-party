package domain

import (
	"time"
)

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name" validate:"min=1,max=30"`
	Email      string    `json:"email" validate:"email"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
}

type IUserUsecase interface {
	GetUser(user *User) error
	CreateUser(user *User) error
}

type IUserRepository interface {
	GetUser(user *User) error
	CreateUser(user *User) error
}
