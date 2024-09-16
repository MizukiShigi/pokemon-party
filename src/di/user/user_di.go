package user

import (
	"database/sql"

	_userHandler "github.com/MizukiShigi/go_pokemon/handler/user"
	"github.com/MizukiShigi/go_pokemon/repository"
	_userUsecase "github.com/MizukiShigi/go_pokemon/usecase/user"
	_userValidator "github.com/MizukiShigi/go_pokemon/validator/user"
)

func InitUser(db *sql.DB) _userHandler.IUserHandler {
	userRepository := repository.NewUserRepository(db)
	userValidator := _userValidator.NewUserValidator()
	userUsecase := _userUsecase.NewUserUsecase(userRepository, userValidator)
	userHandler := _userHandler.NewUserHandler(userUsecase)
	return userHandler
}
