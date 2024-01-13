package user

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/MizukiShigi/go_pokemon/domain"
)

type IUserHandler interface {
	HandleUser(w http.ResponseWriter, r *http.Request)
	GetUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
}

type UserHandler struct {
	uu domain.IUserUsecase
}

type UserResponse struct {
	Status     string `json:"status"`
	StatusCode int    `json:"status_code"`
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
}

func NewUserResponse(id int, name string, email string) *UserResponse {
	return &UserResponse{"ok", 200, id, name, email}
}

func NewUserHandler(uu domain.IUserUsecase) IUserHandler {
	return &UserHandler{uu}
}

func (uh *UserHandler) HandleUser(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		uh.GetUser(w, r)
	case "POST":
		uh.CreateUser(w, r)
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}

func (uh *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/users/"))
	if err != nil {
		myError := domain.NewMyError(domain.InvalidInput, "user_id")
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
	user := domain.User{ID: userId}
	err = uh.uu.GetUser(&user)
	if err != nil {
		if myError, ok := err.(domain.MyError); ok {
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
		return
	}
	resUser := NewUserResponse(user.ID, user.Name, user.Email)
	jsonRes, err := json.Marshal(resUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}

func (uh *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var newUser domain.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		myError := domain.NewMyError(domain.InvalidInput, "post data")
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
	err = uh.uu.CreateUser(&newUser)
	if err != nil {
		if myError, ok := err.(domain.MyError); ok {
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
		return
	}
	resUser := NewUserResponse(newUser.ID, newUser.Email, newUser.Name)
	jsonRes, err := json.Marshal(resUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonRes)
}
