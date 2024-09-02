package auth

import (
	"time"
	"vou/pkg/db/coredb"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("your-secret-key")

func GenerateToken(user coredb.User) (string, error) {
	claims := &jwt.StandardClaims{
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		Issuer:    user.ID.Hex(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
