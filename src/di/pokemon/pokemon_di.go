package pokemon

import (
	"database/sql"
	_pokemonHandler "github.com/MizukiShigi/go_pokemon/handler/pokemon"
	"github.com/MizukiShigi/go_pokemon/repository"
	_pokemonUsecase "github.com/MizukiShigi/go_pokemon/usecase/pokemon"
)

func InitPokemon(db *sql.DB) _pokemonHandler.IPokemonHandler {
	pokemonRepository := repository.NewPokemonRepository(db)
	pokemonUsecase := _pokemonUsecase.NewPokemonUsecase(pokemonRepository)
	return _pokemonHandler.NewPokemonHandler(pokemonUsecase) 
}