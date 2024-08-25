package user

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/MizukiShigi/go_pokemon/domain"
)

type IUserHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	Register(w http.ResponseWriter, r *http.Request)
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
	Status     string     `json:"status"`
	StatusCode int        `json:"status_code"`
	Jwt        string     `json:"jwt"`
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

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jwt, err := uh.uu.Login(&user)
	if err != nil {
		writeError(w, err)
		return
	}
	if len(jwt) == 0 {
		writeError(w, errors.New("failed to login"))
		return
	}
	resUser := NewLoginResponse(jwt)
	jsonRes, err := json.Marshal(resUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func (uh *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var newUser domain.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = uh.uu.Register(&newUser)
	if err != nil {
		writeError(w, err)
		return
	}
	resUser := NewUserCreateResponse(newUser.ID, newUser.Email)
	jsonRes, err := json.Marshal(resUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonRes)
}

func writeError(w http.ResponseWriter, err error) {
	var myError domain.MyError
	if errors.As(err, &myError) {
		errorRes := domain.NewErrorResponse(myError)
		jsonErrorRes, err := json.Marshal(errorRes)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonErrorRes)
		return
	}
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
