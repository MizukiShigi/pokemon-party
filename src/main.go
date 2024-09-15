package main

import (
	"log"
	"net/http"

	"github.com/MizukiShigi/go_pokemon/config"
	_pokemonDi "github.com/MizukiShigi/go_pokemon/di/pokemon"
	_userDi "github.com/MizukiShigi/go_pokemon/di/user"
	"github.com/MizukiShigi/go_pokemon/handler"
	"github.com/MizukiShigi/go_pokemon/infrastructure"
	"github.com/MizukiShigi/go_pokemon/internal"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.New().String()
		},
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `{"time":"${time_rfc3339_nano}","id":"${id}","remote_ip":"${remote_ip}",
		"host":"${host}","method":"${method}","uri":"${uri}","user_agent":"${user_agent}",
		"status":${status},"error":"${error}","latency":${latency},"latency_human":"${latency_human}"` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))
	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		requestID := c.Response().Header().Get(echo.HeaderXRequestID)
		log.Printf(
			`{"id":"%s", "method":"%s", "uri":"%s", "request_body":"%s", "response_body":"%s"}` + "\n",
			requestID,
			c.Request().Method,
			c.Request().RequestURI,
			string(reqBody),
			string(resBody),
		)
	}))

	db := infrastructure.ConnectDB()

	userHandler := _userDi.InitUser(db)
	internal.SetUserRouter(e, userHandler)

	pokemonHandler := _pokemonDi.InitPokemon(db)
	internal.SetPokemonRouter(e, pokemonHandler)

	e.HTTPErrorHandler = handler.CustomErrorHandler

	log.Printf("Starting server on %s\n", config.Config.Port)
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
