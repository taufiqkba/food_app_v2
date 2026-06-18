package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Id          string        `json:"id"`
	Role        string        `json:"role"`
	ExpiredTime time.Duration `json:"expired_time"`
	jwt.RegisteredClaims
}

func GenerateJWT(data Claims, secretKey string) (string, error) {
	var token string
	var err error

	tok := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	token, err = tok.SignedString([]byte(secretKey))
	return token, err
}
