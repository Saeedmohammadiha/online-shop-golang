package utils

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	dto "github.com/OnlineShop/internal/app/dto/error"
	"github.com/OnlineShop/internal/pkg/logger"
)

var ErrValidation = errors.New("validationError")
var ErrDatabase = errors.New("databaseError")
var ErrConvert = errors.New("convertJsonError")
var ErrAlreadyExists = errors.New("recordAlreadyExists")

func responseGenerator(err error) *dto.Error {
	//TODO: add a check for the environment and add or cleat the error inside this function
	res := &dto.Error{
		Data: dto.ErrorData{
			Error:   err,
			Message: "",
			Status:  0,
		},
	}

	//TODO: add more cases if needed and maybe change it to a switch case
	if errors.Is(err, ErrDatabase) {
		res.Data.Message = "internal error, please try again later"
		res.Data.Status = http.StatusInternalServerError
	}
	if errors.Is(err, ErrValidation) {
		res.Data.Message = err.Error()
		res.Data.Status = http.StatusBadRequest
	}
	if errors.Is(err, ErrConvert) {
		res.Data.Message = err.Error()
		res.Data.Status = http.StatusBadRequest
	}
	if errors.Is(err, ErrAlreadyExists) {
		res.Data.Message = "the permission is already exists"
		res.Data.Status = http.StatusBadRequest
	}

	return res
}
func SendErrorResponse(ctx context.Context, err error, w http.ResponseWriter, l logger.Ilogger) {
	response := responseGenerator(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Data.Status)

	//TODO: prevent to send stack trace to the user in the production environment
	json.NewEncoder(w).Encode(err)
	l.Error("an error has sent to user as a response",
		"response error:", response,
		"error", err,
	)
}
