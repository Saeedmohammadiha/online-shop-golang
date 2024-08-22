package utils

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/OnlineShop/internal/app/models"
	"github.com/OnlineShop/internal/pkg/logger"
)

type requestBody string
type requestParams string
type user *models.User

var REQUEST_PARAMS requestParams = "params"
var REQUEST_BODY requestBody = "requestBody"
var USER user = &models.User{}

type ContextKeys interface {
	requestParams | requestBody
}

func GetValueFromCtx[T ContextKeys, U interface{}](ctx context.Context, key T, valueContainer *U, l logger.Ilogger) error {
	rawValue := ctx.Value(key)
	if err := ConvertCtxValueToStruct(&rawValue, valueContainer, l); err != nil {
		return err
	}
	return nil
}

func ConvertCtxValueToStruct[T interface{}](value *any, variable *T, l logger.Ilogger) error {
	jsonRequestBody, err := json.Marshal(value)
	if err != nil {
		l.Error("unable to convert the ctx data to json:",
			"value: ", value,
			"to variable:", variable,
			"error", err)
		return fmt.Errorf("%s%w", ErrConvertTag, err)
	}
	l.Info("the value from ctx is converted to json", "value:", jsonRequestBody)

	err = json.Unmarshal([]byte(jsonRequestBody), variable)
	if err != nil {
		l.Error("error unmarshaling json ctx value to struct",
			"value: ", value,
			"to variable:", variable,
			"error", err)

		return fmt.Errorf("%s%w", ErrConvertTag, err)
	}
	l.Info("successfully converted the ctx json to  struct",
		"value:", value,
		"struct result: ", variable)
	return nil
}
