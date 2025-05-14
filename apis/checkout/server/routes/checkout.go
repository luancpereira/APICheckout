package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/luancpereira/APICheckout/apis/checkout/server/model/request"
	"github.com/luancpereira/APICheckout/apis/checkout/server/model/response"
	commonsServer "github.com/luancpereira/APICheckout/apis/commons/server"
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
//	@Security	JWT
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
funcs for deletes
******/

// godoc
//
//	@Tags		Checkout Orders
//	@Produce	json
//	@Security	JWT
//	@Param		transactionID	path	int64	true	"transactionID"
//	@Success	204
//	@Failure	400	{object}	response.Exception
//	@Router		/api/checkout/transactions/{transactionID} [delete]
func (Checkout) DeleteTransaction(ctx *gin.Context) {
	transactionID, err := commonsServer.Param{}.GetPathParamInt64(ctx, "transactionID", true)
	if err != nil {
		return
	}

	err = service.Checkout{}.DeleteTransactionByID(int32(transactionID))
	if err != nil {
		commonsServer.Response{}.ResponseBadRequest(ctx, err)
		return
	}

	commonsServer.Response{}.ResponseNoContent(ctx)
}

/*****
funcs for deletes
******/

/*****
funcs for gets
******/

// godoc
//
//	@Tags		Checkout Orders
//	@Produce	json
//	@Security	JWT
//	@Param		transactionID	path		int64	true	"transactionID"
//	@Param		country			path		string	true	"country"
//	@Success	200				{object}	response.GetTransactionsByID
//	@Failure	400				{object}	response.Exception
//	@Router		/api/checkout/transactions/{transactionID}/country/{country} [get]
func (Checkout) GetByID(ctx *gin.Context) {
	transactionID, err := commonsServer.Param{}.GetPathParamInt64(ctx, "transactionID", true)
	if err != nil {
		return
	}

	country, err := commonsServer.Param{}.GetPathParamString(ctx, "country", true)
	if err != nil {
		return
	}

	model, err := service.Checkout{}.GetByID(transactionID, country)
	if err != nil {
		commonsServer.Response{}.ResponseBadRequest(ctx, err)
		return
	}

	var res response.GetTransactionsByID
	err = copier.Copy(&res, model)
	if err != nil {
		commonsServer.Response{}.ResponseBadRequest(ctx, err)
		return
	}

	commonsServer.Response{}.ResponseOK(ctx, res)
}

// godoc
//
//	@Tags		Checkout Orders
//	@Produce	json
//	@Security	JWT
//	@Param		country					path		string	true	"country"
//	@Param		limit					query		int32	false	"limit min 1"	default(10)
//	@Param		offset					query		int32	false	"offset min 0"	default(0)
//	@Param		filter_transaction_date	query		string	true	"filter_transaction_date"
//	@Success	200						{object}	response.List{data=[]response.GetTransactions}
//	@Failure	400						{object}	response.Exception
//	@Router		/api/checkout/transactions/country/{country} [get]
func (Checkout) GetList(ctx *gin.Context) {
	country, err := commonsServer.Param{}.GetPathParamString(ctx, "country", true)
	if err != nil {
		return
	}

	filters, _, limit, offset := commonsServer.Param{}.GetQueryParam(ctx)
	filterTransactionDate := filters["transaction_date"]
	if filterTransactionDate == "" {
		commonsServer.Response{}.ResponseBadRequest(ctx, coreError.New("error.transaction.date.required"))
		return
	}

	models, total, err := service.Checkout{}.GetList(filters, limit, offset, country)
	if err != nil {
		commonsServer.Response{}.ResponseBadRequest(ctx, err)
		return
	}

	var res []response.GetTransactions

	for _, model := range models {
		res = append(res, response.GetTransactions{
			ID:                                      model.ID,
			Description:                             model.Description,
			TransactionDate:                         model.TransactionDate,
			TransactionValue:                        model.TransactionValue,
			ExchangeRate:                            model.ExchangeRate,
			TransactionValueConvertedToWishCurrency: model.TransactionValueConvertedToWishCurrency,
		})
	}

	commonsServer.Response{}.ResponseListOk(ctx, res, total)
}

/*****
funcs for gets
******/

/*****
other funcs
******/

/*****
other funcs
******/
