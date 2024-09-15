package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/MizukiShigi/go_pokemon/domain"
	"github.com/labstack/echo/v4"
)

func CustomErrorHandler(err error, c echo.Context) {
	var myError domain.MyError

	log.Printf("Error: %v", err)

	//  独自エラーの場合
	if errors.As(err, &myError) {
		errorRes := domain.NewErrorResponse(myError)
		c.JSON(http.StatusInternalServerError, errorRes)
		return
	}

	// echoのエラーの場合
	if he, ok := err.(*echo.HTTPError); ok {
		c.JSON(he.Code, err)
		return
	}

	// それ以外のエラーの場合
	c.JSON(http.StatusInternalServerError, "Internal Server Error")
}
