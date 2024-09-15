package user

import (
	"net/http"

	"github.com/MizukiShigi/go_pokemon/domain"
	"github.com/labstack/echo/v4"
)

type IUserHandler interface {
	Login(c echo.Context) error
	Register(c echo.Context) error
}

type UserHandler struct {
	uu domain.IUserUsecase
}
type UserDetail struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

type UserResponse struct {
	Status     string     `json:"status"`
	StatusCode int        `json:"status_code"`
	User       UserDetail `json:"user"`
}

type LoginUserResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	Jwt        string `json:"jwt"`
}

func NewUserCreateResponse(id int, email string) *UserResponse {
	return &UserResponse{"ok", 201, UserDetail{ID: id, Email: email}}
}

func NewLoginResponse(jwt string) *LoginUserResponse {
	return &LoginUserResponse{"ok", 200, jwt}
}

func NewUserHandler(uu domain.IUserUsecase) IUserHandler {
	return &UserHandler{uu}
}

func (uh *UserHandler) Login(c echo.Context) error {
	var user domain.User
	err := c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	jwt, err := uh.uu.Login(&user)
	if err != nil {
		return err
	}
	if len(jwt) == 0 {
		return c.String(http.StatusBadRequest, "failed to login")
	}
	resUser := NewLoginResponse(jwt)
	return c.JSON(http.StatusOK, resUser)
}

func (uh *UserHandler) Register(c echo.Context) error {
	var newUser domain.User
	err := c.Bind(&newUser)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	err = uh.uu.Register(&newUser)
	if err != nil {
		return err
	}
	resUser := NewUserCreateResponse(newUser.ID, newUser.Email)
	return c.JSON(http.StatusCreated, resUser)
}
