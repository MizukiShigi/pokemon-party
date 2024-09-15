package internal

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "Authorization header is required")
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token format")
		}

		strToken := bearerToken[1]

		token, err := jwt.Parse(strToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte("SECRET"), nil
		})
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token claims")
		}

		now := time.Now().Unix()

		// 有効期限チェック
		if exp, ok := claims["exp"].(float64); ok {
			if now > int64(exp) {
				return echo.NewHTTPError(http.StatusUnauthorized, "Expired token")
			}
		}

		return next(c)
	}
}
