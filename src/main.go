package main

import (
	"log"
	"net/http"

	"github.com/MizukiShigi/go_pokemon/config"
	"github.com/MizukiShigi/go_pokemon/infrastructure"
	"github.com/MizukiShigi/go_pokemon/internal"
	_userDi "github.com/MizukiShigi/go_pokemon/di/user"
)

func main() {
	db := infrastructure.ConnectDB()
	userHandler := _userDi.InitUser(db)
	internal.SetUserRouter(userHandler)

	log.Printf("Starting server on %s\n", config.Config.Port)
	err := http.ListenAndServe(":"+config.Config.Port, nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
