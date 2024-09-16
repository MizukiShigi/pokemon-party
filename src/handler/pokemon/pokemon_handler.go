package pokemon

import (
	"net/http"
	"strconv"

	"github.com/MizukiShigi/go_pokemon/domain"
	"github.com/labstack/echo/v4"
)

type IPokemonHandler interface {
	GetPokemon(c echo.Context) error
}

type PokemonHandler struct {
	pu domain.IPokemonUsecase
}

type PokemonDetail struct {
	PokemonNumber  int     `json:"pokemon_number"`
	Name           string  `json:"name"`
	Type1          string  `json:"type1"`
	Type2          *string  `json:"type2,omitempty"`
	Height         float64 `json:"height"`
	Weight         float64 `json:"weight"`
	BaseExperience int     `json:"base_experience"`
	ImageUrl       string  `json:"image_url"`
}

type PokemonResponse struct {
	Status     string        `json:"status"`
	StatusCode int           `json:"status_code"`
	Pokemon    PokemonDetail `json:"pokemon"`
}

func NewPokemonResponse(pokemon domain.Pokemon) *PokemonResponse {
	return &PokemonResponse{"ok", 200, PokemonDetail{
		PokemonNumber:  pokemon.PokemonNumber,
		Name:           pokemon.Name,
		Type1:          pokemon.Type1,
		Type2:          pokemon.Type2,
		Height:         pokemon.Height,
		Weight:         pokemon.Weight,
		BaseExperience: pokemon.BaseExperience,
		ImageUrl:       pokemon.ImageUrl,
	}}
}

func NewPokemonHandler(pu domain.IPokemonUsecase) IPokemonHandler {
	return &PokemonHandler{pu}
}

func (ph *PokemonHandler) GetPokemon(c echo.Context) error {
	pokemonNumber, _ := strconv.Atoi(c.Param("pokemon_number"))
	pokemon := domain.Pokemon{PokemonNumber: pokemonNumber}
	err := ph.pu.GetPokemon(&pokemon)
	if err != nil {
		return err
	}
	resPokemon := NewPokemonResponse(pokemon)
	return c.JSON(http.StatusOK, resPokemon)
}
