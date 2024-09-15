package internal

import (
	"github.com/labstack/echo/v4"

	_pokemonHandler "github.com/MizukiShigi/go_pokemon/handler/pokemon"
	_userHandler "github.com/MizukiShigi/go_pokemon/handler/user"
)

func SetUserRouter(e *echo.Echo, uh _userHandler.IUserHandler) {
	e.POST("/register", uh.Register)
	e.POST("/login", uh.Login)
	// http.HandleFunc("/register", methodHandler("POST", uh.Register))
	// http.HandleFunc("/login", methodHandler("POST", uh.Login))
	// http.HandleFunc("/users/", Auth(methodHandler("POST", uh.HandlePokemon)))

}

func SetPokemonRouter(e *echo.Echo, ph _pokemonHandler.IPokemonHandler) {
	m := e.Group("", Auth)
	m.GET("/pokemons/:pokemon_number", ph.GetPokemon)
}
