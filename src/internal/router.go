package internal

import (
	"github.com/labstack/echo/v4"

	_pokemonHandler "github.com/MizukiShigi/go_pokemon/handler/pokemon"
	_userHandler "github.com/MizukiShigi/go_pokemon/handler/user"
	_partyHandler "github.com/MizukiShigi/go_pokemon/handler/party"
)

func SetUserRouter(e *echo.Echo, uh _userHandler.IUserHandler) {
	e.POST("/register", uh.Register)
	e.POST("/login", uh.Login)
}

func SetPokemonRouter(e *echo.Echo, ph _pokemonHandler.IPokemonHandler) {
	m := e.Group("", Auth)
	m.GET("/pokemons/:pokemon_number", ph.GetPokemon)
}

func SetPartyRouter(e *echo.Echo, ph _partyHandler.IPartyHandler) {
	e.POST("/party", ph.AddParty)
}
