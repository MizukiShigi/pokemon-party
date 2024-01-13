package user

import (
	"github.com/MizukiShigi/go_pokemon/domain"

	"github.com/go-playground/validator/v10"
)

type IUserValidator interface {
	GetUserValidate(user domain.User) error
	CreateUserValidate(user domain.User) error
}

type userValidator struct{}

func NewUserValidator() IUserValidator {
	return &userValidator{}
}

func (uv *userValidator) GetUserValidate(user domain.User) error {
	validate := validator.New()
	if err := validate.Var(user.ID, "required"); err != nil {
		myError := domain.NewMyError(domain.InvalidInput, "user_id")
		return myError
	}
	return nil
}

func (uv *userValidator) CreateUserValidate(user domain.User) error {
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		var errFields string
		for _, err := range err.(validator.ValidationErrors) {
			errFields += err.Field()
		}
		myError := domain.NewMyError(domain.InvalidInput, errFields)
		return myError
	}
	return nil
}
