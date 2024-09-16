package main

import (
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
	"gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	e := echo.New()

	e.Use(middleware.RequestIDWithConfig(middleware.RequestIDConfig{
		Generator: func() string {
			return uuid.New().String()
		},
	}))

	setupLogging(e)

	e.Use(middleware.Recover())

	db := infrastructure.ConnectDB()

	userHandler := _userDi.InitUser(db)
	internal.SetUserRouter(e, userHandler)

	pokemonHandler := _pokemonDi.InitPokemon(db)
	internal.SetPokemonRouter(e, pokemonHandler)

	e.HTTPErrorHandler = handler.CustomErrorHandler

	e.Logger.Info("Starting server on %s\n", config.Config.Port)
	if err := e.Start(":8080"); err != http.ErrServerClosed {
		e.Logger.Fatal(err)
	}
}

func setupLogging(e *echo.Echo) {
	accessLogFile := newLoggerFile("./log/app/access.log")
	systemLogFile := newLoggerFile("./log/app/system.log")

	e.Logger.SetOutput(systemLogFile)
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: accessLogFile,
	}))
}

func newLoggerFile(filename string) *lumberjack.Logger {
	return &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    1,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   true,
	}
}
