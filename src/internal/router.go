package internal

import (
	"net/http"

	_userHandler "github.com/MizukiShigi/go_pokemon/handler/user"
)

func SetUserRouter(uh _userHandler.IUserHandler) {
	//  ユーザー情報取得
	http.HandleFunc("/users/", uh.HandleUser)
	// // マスタポケモン取得
	// http.HandleFunc("/pokemons/", r.pokemonHandler)
	// // 手持ちポケモン取得
	// http.HandleFunc("/pokemons/party/", r.partyHandler)
	// // ポケモン交換
	// http.HandleFunc("/pokemons/exchange", r.exchangeHandler)
}

// func (router *Router) userHandler(w http.ResponseWriter, r *http.Request) {
// 	switch r.Method {
//     case "GET":
//         router.uc.GetUser(w, r)
//     default:
//         http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
//     }
// }

// func pokemonHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	log.Println("pokemon")
// }

// func partyHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	log.Println("party")
// }

// func exchangeHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	log.Println("exchange")
// }
