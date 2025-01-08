package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/luancpereira/APICheckout/apis/public/server/model/request"
	"github.com/luancpereira/APICheckout/apis/public/server/model/response"
	coreError "github.com/luancpereira/APICheckout/core/errors"
	"github.com/luancpereira/APICheckout/core/service"
)

type Checkout struct{}

/*****
funcs for posts
******/

// godoc
//
//	@Tags		Checkout Orders
//	@Produce	json
//	@Param		body	body		request.InsertTransaction	true	"Body JSON"
//	@Success	201		{object}	response.Created
//	@Failure	400		{object}	response.Exception
//	@Router		/api/checkout [post]
func (Checkout) InsertTransaction(ctx *gin.Context) {
	var req request.InsertTransaction
	err := GetBody(ctx, &req)
	if err != nil {
		return
	}

	ID, err := service.Checkout{}.CreateTransaction(req.Description, req.TransactionDate, req.TransactionValue)
	if err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	ResponseCreated(ctx, ID)
}

/*****
funcs for posts
******/

/*****
funcs for gets
******/

// godoc
//
//	@Tags		Checkout Orders
//	@Produce	json
//	@Success	200	{object}	response.List{data=[]response.GetOrders}
//	@Failure	400	{object}	response.Exception
//	@Router		/api/checkout/transactions [get]
func (Checkout) GetList(ctx *gin.Context) {
	models, total, err := service.Checkout{}.GetList()
	if err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	var res []response.GetOrders

	for _, model := range models {
		res = append(res, response.GetOrders{
			ID:               model.ID,
			Description:      model.Description,
			TransactionDate:  model.TransactionDate,
			TransactionValue: model.TransactionValue,
		})
	}

	ResponseListOk(ctx, res, total)
}

/*****
funcs for gets
******/

/*****
other funcs
******/

func GetBody(ctx *gin.Context, obj any) (err error) {
	err = ParseBody(ctx, obj)
	if err != nil {
		ResponseBadRequest(ctx, err)
		return
	}

	return
}

func ParseBody(ctx *gin.Context, obj any) (err error) {
	err = ctx.ShouldBindJSON(obj)
	if err != nil {
		err = coreError.New("error.request.body.invalid", err.Error())
		return
	}

	return
}

func ResponseListOk(ctx *gin.Context, bodyResponse any, total int64) {
	var list response.List

	list.Pagination = response.Pagination{Total: total}
	list.Data = bodyResponse

	ctx.JSON(http.StatusOK, list)
}

func ResponseCreated(ctx *gin.Context, ID int64) {
	bodyResponse := response.Created{ID: ID}

	ResponseCreatedBody(ctx, bodyResponse)
}

func ResponseCreatedBody(ctx *gin.Context, bodyResponse any) {
	ctx.JSON(http.StatusCreated, bodyResponse)
}

func ResponseBadRequest(ctx *gin.Context, err interface{}) {
	errOut := coreError.ConvertTo(err)

	ctx.AbortWithStatusJSON(http.StatusBadRequest, errOut)
}

/*****
other funcs
******/
