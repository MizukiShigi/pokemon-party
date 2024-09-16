package domain

import (
	"time"
)

type Pokemon struct {
	Id             int       `json:"id"`
	PokemonNumber  int       `json:"pokemon_number"`
	Name           string    `json:"name"`
	Type1          string    `json:"type1"`
	Type2          *string   `json:"type2"`
	Height         float64   `json:"height"`
	Weight         float64   `json:"weight"`
	BaseExperience int       `json:"base_experience"`
	ImageUrl       string    `json:"image_url"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type IPokemonUsecase interface {
	GetPokemon(pokemon *Pokemon) error
}

type IPokemonRepository interface {
	InsertPokeomons(pokemons []Pokemon) error
	GetPokemonByPokemonNumber(pokemonNumber int) (Pokemon, error)
}
