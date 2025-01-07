package routes

import (
	"github.com/gin-gonic/gin"
	commonsServer "github.com/luancpereira/APICheckout/apis/commons/server"
	"github.com/luancpereira/APICheckout/apis/public/server/model/request"
	"github.com/luancpereira/APICheckout/apis/public/server/model/response"
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
	err := commonsServer.Param{}.GetBody(ctx, &req)
	if err != nil {
		return
	}

	ID, err := service.Checkout{}.CreateTransaction(req.Description, req.TransactionDate, req.TransactionValue)
	if err != nil {
		commonsServer.Response{}.ResponseBadRequest(ctx, err)
		return
	}

	commonsServer.Response{}.ResponseCreated(ctx, ID)
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
		commonsServer.Response{}.ResponseBadRequest(ctx, err)
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

	commonsServer.Response{}.ResponseListOk(ctx, res, total)
}

/*****
funcs for gets
******/
