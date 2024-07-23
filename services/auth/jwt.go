package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/muhammadderic/ecomrest/configs"
)

func CreateJWT(secret []byte, userId int) (string, error) {
	// Set the token expiration time
	expiration := time.Second * time.Duration(configs.Envs.JWTExpirationInSeconds)

	// Create a new JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId":    strconv.Itoa(userId),
		"expiresAt": time.Now().Add(expiration).Unix(),
	})

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
