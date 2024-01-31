package internal

import (
	"net/http"

	_userHandler "github.com/MizukiShigi/go_pokemon/handler/user"
	_pokemonHandler "github.com/MizukiShigi/go_pokemon/handler/pokemon"
)

func SetUserRouter(uh _userHandler.IUserHandler) {
	//  ユーザー情報取得
	http.HandleFunc("/users/", uh.HandleUser)
	// ユーザー情報作成
	http.HandleFunc("/users", uh.HandleUser)
}

func SetPokemonRouter(ph _pokemonHandler.IPokemonHandler) {
	// マスタポケモン1匹取得
	http.HandleFunc("/pokemons/", ph.HandlePokemon)
	// マスタポケモン複数匹取得
	http.HandleFunc("/pokemons", ph.HandlePokemon)
}