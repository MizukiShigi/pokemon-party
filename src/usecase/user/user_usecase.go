package user

import (
	"errors"
	"time"

	"github.com/MizukiShigi/go_pokemon/domain"
	"github.com/MizukiShigi/go_pokemon/validator/user"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	ur domain.IUserRepository
	uv user.IUserValidator
}

func NewUserUsecase(ur domain.IUserRepository, uv user.IUserValidator) domain.IUserUsecase {
	return &UserUsecase{ur, uv}
}

func generateJwt(userId int) string {
	claims := jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, _ := token.SignedString([]byte("SECRET"))
	return tokenString
}

func (uu *UserUsecase) Login(reqUser *domain.User) (string, error) {
	dbUser, err := uu.ur.GetUserByEmail(reqUser.Email)
	if err != nil {
		return "", err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(reqUser.Password)); err != nil {
		return "", err
	}

	token := generateJwt(dbUser.ID)

	return token, nil
}

func (uu *UserUsecase) Register(user *domain.User) error {
	if err := uu.uv.CreateUserValidate(*user); err != nil {
		return err
	}

	isDuplicate, err := uu.ur.CheckDuplicateEmail(user.Email)
	if err != nil {
		return err
	}

	if isDuplicate {
		return errors.New("既に登録済みのユーザーです")
	}

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}

	user.Password = string(hashPass)

	return uu.ur.CreateUser(user)
}
