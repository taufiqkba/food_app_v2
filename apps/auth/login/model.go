package login

import (
	"time"

	"github.com/taufiqkba/food_app_v2/internal/utils/encryption"
	token "github.com/taufiqkba/food_app_v2/internal/utils/jwt"
)

type Auth struct {
	Id        int         `json:"id"`
	PublicID  string      `json:"public_id"`
	Email     string      `json:"email"`
	Password  string      `json:"password"`
	Role      string      `json:"role"`
	IsActive  bool        `json:"is_active"`
	CreatedAt time.Time   `json:"created_at"`
	UpdatedAt time.Ticker `json:"updated_at"`
}

func (a Auth) ValidatePassowrd(password string) error {
	if err := encryption.ValidatePassword(a.Password, password); err != nil {
		return errEmailOrPasswordIsNotMatched
	}
	return nil
}

func (a Auth) GenerateToken(data token.Claims, secretKey string) (string, error) {
	token, err := token.GenerateJWT(data, secretKey)
	if err != nil {
		return "", err
	}

	return token, nil
}
