package utils

import (
	"context"
	"encoding/json"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	"github.com/OnlineShop/internal/pkg/logger"
)

type requestBody string
type requestParams string
type user string
type token string

const REQUEST_PARAMS requestParams = "params"
const REQUEST_BODY requestBody = "requestBody"
const USER user = "user"
const TOKEN token = "token"

type ContextKeys interface {
	requestParams | requestBody | user | token
}

func GetValueFromCtx[T ContextKeys, U interface{}](ctx context.Context, key T, valueContainer *U) *apperrors.AppError {
	l := logger.Logger()
	rawValue := ctx.Value(key)
	if rawValue == nil {
		l.Error("there is no value in the context with this key",
			"key", key,
			"value", rawValue,
		)
		return apperrors.NewInternalError("there is no value in context", nil)
	}
	if err := ConvertCtxValueToStruct(&rawValue, valueContainer); err != nil {
		return err
	}
	return nil
}

func ConvertCtxValueToStruct[T interface{}](value *any, variable *T) *apperrors.AppError {
	l := logger.Logger()
	jsonRequestBody, err := json.Marshal(value)
	if err != nil {
		l.Error("unable to convert the ctx data to json:",
			"value: ", value,
			"to variable:", variable,
			"error", err)
		return apperrors.NewInternalError("failed to convert ctx data to json", err)
	}
	l.Info("the value from ctx is converted to json", "value:", jsonRequestBody)

	err = json.Unmarshal([]byte(jsonRequestBody), variable)
	if err != nil {
		l.Error("error unmarshaling json ctx value to struct",
			"value: ", value,
			"to variable:", variable,
			"error", err)

		return apperrors.NewInternalError("failed to convert the json to struct", err)
	}
	l.Info("successfully converted the ctx json to  struct",
		"value:", value,
		"struct result: ", variable)
	return nil
}
