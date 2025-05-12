package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserPayload struct {
	Email  string
	UserID int64
	jwt.RegisteredClaims
}

func (UserPayload) New(email string, userID int64, duration time.Duration) (payload UserPayload) {
	jwtClaims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	payload = UserPayload{
		email,
		userID,
		jwtClaims,
	}

	return
}
