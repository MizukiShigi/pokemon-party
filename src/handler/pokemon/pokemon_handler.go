package pokemon

// import (
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// 	"strconv"
// 	"strings"

// 	"github.com/MizukiShigi/go_pokemon/domain"
// 	"github.com/shopspring/decimal"
// )

// type IPokemonHandler interface {
// 	HandlePokemon(w http.ResponseWriter, r *http.Request)
// 	GetPokemon(w http.ResponseWriter, r *http.Request)
	
// }

// type PokemonHandler struct {
// 	pu domain.IPokemonUsecase
// }
// type PokemonDetail struct {
// 	ID    int    `json:"id"`
// 	Name  string `json:"name"`
// 	Type1 string `json:"type1"`
// 	Type2 string `json:"type2"`
// 	Height string `json:"height"`
// 	Weight string `json:"weight"`
// 	BaseExperience int `json:"base_experience"`
// 	ImageUrl string `json:"image_url"`
// }

// type PokemonResponse struct {
// 	Status     string     `json:"status"`
// 	StatusCode int        `json:"status_code"`
// 	Pokemon       PokemonDetail `json:"pokemon"`
// }

// func NewPokemonResponse(id int, name string, email string) *PokemonResponse {
// 	return &PokemonResponse{"ok", 200, PokemonDetail{ID: id, Name: name, Email: email}}
// }

// func NewPokemonHandler(pu domain.IPokemonUsecase) IPokemonHandler {
// 	return &PokemonHandler{pu}
// }

// func (ph *PokemonHandler) HandlePokemon(w http.ResponseWriter, r *http.Request) {
// 	if (r.Method != "GET") {
// 		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
// 		return
// 	}
// 	if strings.HasPrefix(r.URL.Path, "/pokemons/") {
// 		ph.GetPokemon(w, r)
// 			return
// 	}
// 	if r.URL.Path == "/pokemons" {
// 		ph.GetPokemons(w, r)
// 		return
// 	}
// 	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		
// }

// func (ph *PokemonHandler) GetPokemon(w http.ResponseWriter, r *http.Request) {
// 	pokemonId, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/pokemons/"))
// 	if err != nil {
// 		myError := domain.NewMyError(domain.InvalidInput, "pokemon_id")
// 		errorRes := domain.NewErrorResponse(myError)
// 		jsonErrorRes, err := json.Marshal(errorRes)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(jsonErrorRes)
// 		return
// 	}
// 	pokemon := domain.Pokemon{ID: pokemonId}
// 	err = ph.pu.GetPokemon(&pokemon)
// 	if err != nil {
// 		writeError(w, err)
// 		return
// 	}
// 	resPokemon := NewPokemonResponse(pokemon.ID, pokemon.Name, pokemon.Email)
// 	jsonRes, err := json.Marshal(resPokemon)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(jsonRes)
// }

// func (ph *PokemonHandler) GetPokemons(w http.ResponseWriter, r *http.Request) {
// 	pokemonId, err := strconv.Atoi(strings.TrimPrefix(r.URL.Path, "/pokemons/"))
// 	if err != nil {
// 		myError := domain.NewMyError(domain.InvalidInput, "pokemon_id")
// 		errorRes := domain.NewErrorResponse(myError)
// 		jsonErrorRes, err := json.Marshal(errorRes)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(jsonErrorRes)
// 		return
// 	}
// 	pokemon := domain.Pokemon{ID: pokemonId}
// 	err = ph.pu.GetPokemon(&pokemon)
// 	if err != nil {
// 		writeError(w, err)
// 		return
// 	}
// 	resPokemon := NewPokemonResponse(pokemon.ID, pokemon.Name, pokemon.Email)
// 	jsonRes, err := json.Marshal(resPokemon)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(jsonRes)
// }

// func writeError(w http.ResponseWriter, err error) {
// 	var myError domain.MyError
// 	if errors.As(err, &myError) {
// 		errorRes := domain.NewErrorResponse(myError)
// 		jsonErrorRes, err := json.Marshal(errorRes)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(jsonErrorRes)
// 		return
// 	}
// 	http.Error(w, err.Error(), http.StatusInternalServerError)
// }