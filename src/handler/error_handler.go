package handler

import (
	"errors"
	"net/http"

	"github.com/MizukiShigi/go_pokemon/domain"
	"github.com/labstack/echo/v4"
)

func CustomErrorHandler(err error, c echo.Context) {
	var myError domain.MyError

	c.Logger().Error(err)

	//  独自エラーの場合
	if errors.As(err, &myError) {
		errorRes := domain.NewErrorResponse(myError)
		err := c.JSON(http.StatusInternalServerError, errorRes)
		if err != nil {
			c.Logger().Error(err)
		}
		return
	}

	// echoのエラーの場合
	if he, ok := err.(*echo.HTTPError); ok {
		err := c.JSON(he.Code, he.Message)
		if err != nil {
			c.Logger().Error(err)
		}
		return
	}

	// それ以外のエラーの場合
	err = c.JSON(http.StatusInternalServerError, "Internal Server Error")
	if err != nil {
		c.Logger().Error(err)
	}
}
