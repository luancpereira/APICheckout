package server

import (
	"strconv"
	"strings"

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

func (Param) GetPathParamString(ctx *gin.Context, key string, required bool) (value string, err error) {
	value = ctx.Param(key)

	if required && len(strings.TrimSpace(value)) == 0 {
		err = coreError.New("error.request.path.param.invalid", key)
		Response{}.ResponseBadRequest(ctx, err)
		return
	}

	return
}

func (s Param) GetPathParamInt64(ctx *gin.Context, key string, required bool) (value int64, err error) {
	param, err := s.GetPathParamString(ctx, key, required)
	if err != nil {
		return
	}

	value, err = strconv.ParseInt(param, 10, 64)
	if err != nil {
		err = coreError.New("error.request.path.param.invalid", key)
		Response{}.ResponseBadRequest(ctx, err)
	}

	return
}

func (Param) GetQueryParam(ctx *gin.Context) (filters map[string]string, sorts map[string]string, limit, offset int64) {
	filters = make(map[string]string)
	sorts = make(map[string]string)
	limit = 10
	offset = 0

	for parameter, value := range ctx.Request.URL.Query() {
		if strings.HasPrefix(strings.ToLower(parameter), "sort_") {
			sorts[strings.ReplaceAll(parameter, "sort_", "")] = value[0]
		}

		if strings.HasPrefix(strings.ToLower(parameter), "filter_") {
			filters[strings.ReplaceAll(parameter, "filter_", "")] = value[0]
		}

		if parameter == "limit" {
			limitInt, _ := strconv.Atoi(value[0])
			limit = int64(limitInt)
		}

		if parameter == "offset" {
			offsetInt, _ := strconv.Atoi(value[0])
			offset = int64(offsetInt)
		}
	}

	return
}

func (Param) GetQueryParamRaw(ctx *gin.Context) (parameters string) {
	queryUrl := strings.Split(ctx.Request.URL.String(), "?")
	parameters = ""
	if len(queryUrl) > 1 {
		parameters = queryUrl[1]
	}

	return
}

func (Param) GetQueryParamString(ctx *gin.Context, key string, required bool) (value string, err error) {
	value = ctx.Query(key)

	if required && len(strings.TrimSpace(value)) == 0 {
		err = coreError.New("error.request.path.param.invalid", key)
		Response{}.ResponseBadRequest(ctx, err)
	}

	return
}

func (Param) GetQueryParamInt64(ctx *gin.Context, key string, required bool) (value int64, err error) {
	rawValue := ctx.Query(key)

	if len(strings.TrimSpace(rawValue)) == 0 {
		if required {
			err = coreError.New("error.request.path.param.required", key)
			Response{}.ResponseBadRequest(ctx, err)
		}

		return
	}

	value, err = strconv.ParseInt(rawValue, 10, 64)
	if err != nil {
		err = coreError.New("error.request.path.param.invalid", key)
		Response{}.ResponseBadRequest(ctx, err)
	}

	return
}
