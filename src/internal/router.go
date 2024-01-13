package internal

import (
	"net/http"

	_userHandler "github.com/MizukiShigi/go_pokemon/handler/user"
)

func SetUserRouter(uh _userHandler.IUserHandler) {
	//  ユーザー情報取得
	http.HandleFunc("/users/", uh.HandleUser)
	// ユーザー情報作成
	http.HandleFunc("/users", uh.HandleUser)
	// // マスタポケモン取得
	// http.HandleFunc("/pokemons/", r.pokemonHandler)
	// // 手持ちポケモン取得
	// http.HandleFunc("/pokemons/party/", r.partyHandler)
	// // ポケモン交換
	// http.HandleFunc("/pokemons/exchange", r.exchangeHandler)
}