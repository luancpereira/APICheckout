package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/luancpereira/APICheckout/apis/checkout/server/model/request"
	"github.com/luancpereira/APICheckout/apis/checkout/server/model/response"
	commonsServer "github.com/luancpereira/APICheckout/apis/commons/server"
	commonsToken "github.com/luancpereira/APICheckout/apis/commons/token"
	"github.com/luancpereira/APICheckout/core/service"
)

type Auth struct{}

/*****
funcs for posts
******/

// godoc
//
//	@Tags		Auth
//	@Produce	json
//	@Param		body	body		request.PostAuthLogin	true	"Body JSON"
//	@Failure	400		{object}	response.Exception
//	@Failure	401		{object}	response.Exception
//	@Success	200		{object}	response.PostAuthLogin
//	@Router		/api/auth/login [post]
func (Auth) PostLogin(ctx *gin.Context) {
	var req request.PostAuthLogin
	err := commonsServer.Param{}.GetBody(ctx, &req)
	if err != nil {
		return
	}

	model, err := service.User{}.Login(req.Email, req.Password)
	if err != nil {
		commonsServer.Response{}.ResponseUnauthorized(ctx)
		return
	}

	token, err := commonsToken.JWT{}.CreateToken(model.Email, model.ID)
	if err != nil {
		commonsServer.Response{}.ResponseUnauthorized(ctx)
		return
	}

	res := response.PostAuthLogin{
		Token: token,
	}

	commonsServer.Response{}.ResponseOK(ctx, res)
}

/*****
funcs for posts
******/
