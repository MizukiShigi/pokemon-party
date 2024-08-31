package pokemon

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/MizukiShigi/go_pokemon/domain"
)

type IPokemonHandler interface {
	GetPokemon(w http.ResponseWriter, r *http.Request)
}

type PokemonHandler struct {
	pu domain.IPokemonUsecase
}

type PokemonDetail struct {
	PokemonNumber  int     `json:"pokemon_number"`
	Name           string  `json:"name"`
	Type1          string  `json:"type1"`
	Type2          string  `json:"type2"`
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

func (ph *PokemonHandler) GetPokemon(w http.ResponseWriter, r *http.Request) {
	pokemonNumber, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/pokemons/"))
	if err != nil {
		myError := domain.NewMyError(domain.InvalidInput, "pokemon_number")
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
	pokemon := domain.Pokemon{PokemonNumber: pokemonNumber}
	err = ph.pu.GetPokemon(&pokemon)
	if err != nil {
		writeError(w, err)
		return
	}
	resPokemon := NewPokemonResponse(pokemon)
	jsonRes, err := json.Marshal(resPokemon)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
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
