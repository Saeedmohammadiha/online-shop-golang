package utils

import (
	"encoding/json"
	"fmt"

	"github.com/OnlineShop/internal/pkg/logger"
)

func ConvertCtxValueToStruct[T interface{}](value *any, variable *T, l logger.Ilogger) error {
	jsonRequestBody, err := json.Marshal(value)
	if err != nil {
		l.Error("unable to convert the ctx data to json:",
			"value: ", value,
			"to variable:", variable,
			"error", err)
		return fmt.Errorf("failed to convert json: %w", ErrConvert)
	}
	l.Info("the value from ctx is converted to json", "value:", jsonRequestBody)

	err = json.Unmarshal([]byte(jsonRequestBody), variable)
	if err != nil {
		l.Error("error unmarshaling json ctx value to struct",
			"value: ", value,
			"to variable:", variable,
			"error", err)

		return fmt.Errorf("failed to convert json: %w", ErrConvert)
	}
	l.Info("successfully converted the ctx json to  struct",
		"value:", value,
		"struct result: ", variable)
	return nil
}
