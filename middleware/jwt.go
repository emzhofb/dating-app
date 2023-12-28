package middleware

import (
	"dating-app/constants"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
)

func GenerateJWT(customerId int) (string, error) {
	claims := jwt.MapClaims{}
	claims["customerId"] = customerId
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(constants.JWT_SECRET))
}

func GetUserIdFromJWT(c echo.Context) int {
	customer := c.Get("user").(*jwt.Token)
	if customer.Valid {
		claims := customer.Claims.(jwt.MapClaims)
		customerId := claims["customerId"].(float64)
		return int(customerId)
	}
	return 0
}
