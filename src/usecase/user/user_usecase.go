package user

import (
	"github.com/MizukiShigi/go_pokemon/domain"
	"github.com/MizukiShigi/go_pokemon/validator/user"
)

type UserUsecase struct {
	ur domain.IUserRepository
	uv user.IUserValidator
}

func NewUserUsecase(ur domain.IUserRepository, uv user.IUserValidator) domain.IUserUsecase {
	return &UserUsecase{ur, uv}
}

func (uu *UserUsecase) GetUser(user *domain.User) error {
	if err := uu.uv.GetUserValidate(*user); err != nil {
		return err
	}
	err := uu.ur.GetUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (uu *UserUsecase) CreateUser(user *domain.User) (error) {
	if err := uu.ur.CreateUser(user); err != nil {
		return err
	}
	return nil
}
