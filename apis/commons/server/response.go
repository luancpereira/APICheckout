package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luancpereira/APICheckout/apis/commons/server/model/response"
	coreError "github.com/luancpereira/APICheckout/core/errors"
)

type Response struct{}

func (Response) ResponseListOk(ctx *gin.Context, bodyResponse any, total int64) {
	var list response.List

	list.Pagination = response.Pagination{Total: total}
	list.Data = bodyResponse

	ctx.JSON(http.StatusOK, list)
}

func (r Response) ResponseCreated(ctx *gin.Context, ID int64) {
	bodyResponse := response.Created{ID: ID}

	r.ResponseCreatedBody(ctx, bodyResponse)
}

func (Response) ResponseCreatedBody(ctx *gin.Context, bodyResponse any) {
	ctx.JSON(http.StatusCreated, bodyResponse)
}

func (Response) ResponseBadRequest(ctx *gin.Context, err interface{}) {
	errOut := coreError.ConvertTo(err)

	ctx.AbortWithStatusJSON(http.StatusBadRequest, errOut)
}
