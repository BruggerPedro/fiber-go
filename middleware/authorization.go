package middleware

import (
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func SplitToken(headerToken string) string {
	parsToken := strings.SplitAfter(headerToken, " ")
	tokenString := parsToken[1]
	return tokenString
}

func AuthenticateToken(tokenString string) error {
	//token check
	_, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil

}

func Auth(headerToken string) error {
	if headerToken == "" {
		return fiber.NewError(401, "Unauthorized")
	}
	if err := AuthenticateToken(SplitToken(headerToken)); err != nil {
		return fiber.NewError(401, "Unauthorized")
	}

	return nil
}
