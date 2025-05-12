package errors

import (
	"github.com/jellydator/ttlcache/v3"
	commonsUtils "github.com/luancpereira/APICheckout/apis/commons/utils"
	log "github.com/sirupsen/logrus"
)

type CoreError struct {
	Key        string               `json:"key"`
	Message    string               `json:"message"`
	Attributes []CoreAttributeError `json:"attributes"`
}

type CoreAttributeError struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type CoreErrorField struct {
	Field   string `json:"field"`
	Key     string `json:"key"`
	Message string `json:"message"`
}

func New(keys ...string) *CoreError {
	var cacheMsg *ttlcache.Item[string, string]
	msgKey := keys[0]
	var message string

	if len(keys) > 1 {
		for i := 1; i < len(keys); i++ {
			message = commonsUtils.ConcatenateStrings(message, keys[i])
		}
	}
	cacheMsg = C.Get(msgKey)

	if cacheMsg == nil {
		log.Errorf("%s", commonsUtils.ConcatenateStrings("error not in errors.json please contact the dev team:", msgKey))
		return &CoreError{Key: msgKey}
	}

	if commonsUtils.StringIsNotEmpty(message) {
		fullMessageError := commonsUtils.ConcatenateStrings(cacheMsg.Value(), " ", message)
		return &CoreError{Key: cacheMsg.Key(), Message: fullMessageError}

	}
	return &CoreError{Key: cacheMsg.Key(), Message: cacheMsg.Value()}
}

func ConvertTo(err interface{}) *CoreError {
	errOut, ok := err.(*CoreError)
	if !ok {
		errDefault, _ := err.(error)
		errOut = New("error.unmapped", errDefault.Error())
	}

	return errOut
}

func (e *CoreError) Error() string {
	return commonsUtils.ConcatenateStrings(e.Key, " | ", e.Message)
}

func MakeErrorField(err error, field string, errorFields *[]CoreErrorField) (hasError bool) {
	if err == nil {
		return
	}

	errOut := ConvertTo(err)
	if errOut == nil {
		return
	}

	*errorFields = append(*errorFields, CoreErrorField{
		Field:   field,
		Key:     errOut.Key,
		Message: errOut.Message,
	})

	return len(*errorFields) > 0
}
