package server

import (
	"github.com/gin-gonic/gin"

	coreError "github.com/luancpereira/APICheckout/core/errors"
)

type Param struct{}

func (p Param) GetBody(ctx *gin.Context, obj any) (err error) {
	err = p.ParseBody(ctx, obj)
	if err != nil {
		Response{}.ResponseBadRequest(ctx, err)
		return
	}

	return
}

func (Param) ParseBody(ctx *gin.Context, obj any) (err error) {
	err = ctx.ShouldBindJSON(obj)
	if err != nil {
		err = coreError.New("error.request.body.invalid", err.Error())
		return
	}

	return
}
