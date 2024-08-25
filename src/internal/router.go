package internal

import (
	"net/http"

	_userHandler "github.com/MizukiShigi/go_pokemon/handler/user"
	// _pokemonHandler "github.com/MizukiShigi/go_pokemon/handler/pokemon"
)

func methodHandler(method string, h http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if r.Method != method {
            http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
            return
        }
        h(w, r)
    }
}

func SetUserRouter(uh _userHandler.IUserHandler) {
	http.HandleFunc("/register", methodHandler("POST", uh.Register))
	http.HandleFunc("/login", methodHandler("POST", uh.Login))
}

// func SetPokemonRouter(ph _pokemonHandler.IPokemonHandler) {
// 	// マスタポケモン1匹取得
// 	http.HandleFunc("/pokemons/", ph.HandlePokemon)
// 	// マスタポケモン複数匹取得
// 	http.HandleFunc("/pokemons", ph.HandlePokemon)
// }