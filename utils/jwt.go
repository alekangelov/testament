package utils

import (
	"os"

	"github.com/golang-jwt/jwt/v5"
)

func getJwtAlgorithm(input string) jwt.SigningMethod {
	switch input {
	case "HS256":
		return jwt.SigningMethodHS256
	case "HS384":
		return jwt.SigningMethodHS384
	case "HS512":
		return jwt.SigningMethodHS512
	default:
		return jwt.SigningMethodHS256
	}

}
func SignJWT(c jwt.Claims) string {
	secret := os.Getenv("JWT_SECRET")
	algorithm := os.Getenv("JWT_ALGORITHM")
	signingMethod := getJwtAlgorithm(algorithm)
	token := jwt.NewWithClaims(signingMethod, c)
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		panic(err)
	}
	return tokenString
}
