package routes

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/luancpereira/APICheckout/apis/checkout/server/model/request"
	"github.com/luancpereira/APICheckout/apis/checkout/server/model/response"
	commonsServer "github.com/luancpereira/APICheckout/apis/commons/server"
	commonsToken "github.com/luancpereira/APICheckout/apis/commons/token"
	"github.com/luancpereira/APICheckout/core/database/sqlc"
	"github.com/luancpereira/APICheckout/core/enums"
	coreError "github.com/luancpereira/APICheckout/core/errors"
	"github.com/luancpereira/APICheckout/core/service"
)

type User struct{}

/*****
funcs for posts
******/

// godoc
//
//	@Tags		User
//	@Produce	json
//	@Param		body	body		request.PostUser	true	"Body JSON"
//	@Success	201		{object}	response.PostUser
//	@Failure	400		{object}	[]response.PostUserException
//	@Router		/api/users [post]
func (User) PostUser(ctx *gin.Context) {
	var req request.PostUser
	err := commonsServer.Param{}.ParseBody(ctx, &req)
	if err != nil {
		var errorFields []coreError.CoreErrorField
		coreError.MakeErrorField(err, enums.USER_VALIDATION_FIELD_FORM, &errorFields)
		commonsServer.Response{}.ResponseBadRequestBody(ctx, errorFields)
		return
	}

	model := sqlc.InsertUserParams{
		Email:      req.Email,
		Name:       req.Name,
		Password:   req.Password,
		Permission: enums.USER_PERMISSION_TYPE_NORMAL,
		CreatedAt:  time.Now(),
	}

	ID, errorFields := service.User{}.Create(model, req.RepeatPassword)
	if len(errorFields) > 0 {
		commonsServer.Response{}.ResponseBadRequestBody(ctx, errorFields)
		return
	}

	accessToken, err := commonsToken.JWT{}.CreateToken(model.Email, ID)
	if err != nil {
		coreError.MakeErrorField(err, enums.USER_VALIDATION_FIELD_FORM, &errorFields)
		commonsServer.Response{}.ResponseBadRequestBody(ctx, errorFields)
		return
	}

	res := response.PostUser{
		Token: accessToken,
	}
	commonsServer.Response{}.ResponseCreatedBody(ctx, res)
}

/*****
funcs for posts
******/
