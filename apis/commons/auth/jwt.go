package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/luancpereira/APICheckout/apis/commons/token"
	coreError "github.com/luancpereira/APICheckout/core/errors"
)

type JWT struct{}

func (j JWT) Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		payload, err := j.payload(ctx)
		if err != nil {
			return
		}

		ctx.Set(AUTH_MIDDLEWARE_AUTHORIZATION_PAYLOAD, payload)
		ctx.Next()
	}
}

func (JWT) payload(ctx *gin.Context) (payload *token.UserPayload, err error) {
	authorizationHeader := ctx.GetHeader(AUTH_MIDDLEWARE_HEADER_AUTHORIZATION)
	if len(authorizationHeader) == 0 {
		err = coreError.New("error.auth.header.invalid")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	fields := strings.Fields(authorizationHeader)
	if len(fields) < 2 {
		err = coreError.New("error.auth.header.invalid")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	authorizationType := strings.ToLower(fields[0])
	if authorizationType != AUTH_MIDDLEWARE_HEADER_BEARER {
		err = coreError.New("error.auth.header.invalid")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	accessToken := fields[1]
	payload, err = token.JWT{}.VerifyToken(accessToken)
	if err != nil {
		err = coreError.New("error.auth.login.unauthorized")
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, err)
		return
	}

	return
}

func (JWT) Info(ctx *gin.Context) (userPayload *token.UserPayload) {
	payload, exists := ctx.Get(AUTH_MIDDLEWARE_AUTHORIZATION_PAYLOAD)
	if !exists {
		return
	}

	userPayload, _ = payload.(*token.UserPayload)

	return
}
