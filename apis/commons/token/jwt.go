package token

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	commonsConfig "github.com/luancpereira/APICheckout/apis/commons/config"
	coreError "github.com/luancpereira/APICheckout/core/errors"
)

type JWT struct{}

func (JWT) CreateToken(email string, userID int64) (token string, err error) {
	duration, err := time.ParseDuration(commonsConfig.EXPIRATION_TOKEN_DEFAULT)
	if err != nil {
		return "", coreError.New("error.login.create.token")
	}

	payload := UserPayload{}.New(email, userID, duration)
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err = jwtToken.SignedString([]byte(commonsConfig.JWT_SECRET))
	if err != nil {
		return "", coreError.New("error.login.create.token")
	}

	return
}

func (JWT) VerifyToken(token string) (payload *UserPayload, err error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, coreError.New("error.auth.login.unauthorized")
		}
		return []byte(commonsConfig.JWT_SECRET), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &UserPayload{}, keyFunc)
	if err != nil {
		return
	}

	payload, ok := jwtToken.Claims.(*UserPayload)
	if !ok {
		err = coreError.New("error.auth.login.unauthorized")
		return
	}

	return
}
