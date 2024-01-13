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

func (uu *UserUsecase) GetUser(user domain.User) (domain.User, error) {
	if err := uu.uv.GetUserValidate(user); err != nil {
		return domain.User{}, err
	}
	storedUser := domain.User{}
	err := uu.ur.GetUser(&storedUser, user.ID)
	if err != nil {
		return domain.User{}, err
	}
	return storedUser, nil
}

func (uu *UserUsecase) CreateUser(user domain.User) (domain.User, error) {
	// newUser := domain.User{
	// 	Name: user.Name,
	// 	Email: user.Email,
	// }
	if err := uu.ur.CreateUser(&user); err != nil {
		return domain.User{}, err
	}
	return user, nil
}
