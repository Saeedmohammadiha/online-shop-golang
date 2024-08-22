package utils

import (
	"context"
	"encoding/json"
	"net/http"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	dto "github.com/OnlineShop/internal/app/dto/error"
	"github.com/OnlineShop/internal/pkg/logger"
)

func responseGenerator(err *apperrors.AppError) *dto.Error {
	//TODO: add a check for the environment and add or cleat the error inside this function
	res := dto.Error{
		Data: dto.ErrorData{
			Error:   nil,
			Message: "",
			Status:  0,
		},
	}

	//TODO: add more cases if needed and maybe change it to a switch case
	if err.Tag == apperrors.AuthenticationErrorTag {
		res.Data.Message = err.Error()
		res.Data.Status = http.StatusUnauthorized
		return &res
	}
	if err.Tag == apperrors.InternalErrorTag {
		res.Data.Message = "internal error, please try again later"
		res.Data.Status = http.StatusInternalServerError
		res.Data.Error = err.Err
		return &res
	}
	if err.Tag == apperrors.ValidationErrorTag {
		res.Data.Message = err.Error()
		res.Data.Status = http.StatusBadRequest
		res.Data.Error = err.Err
		return &res
	}
	if err.Tag == apperrors.AuthorizationErrorTag {
		res.Data.Message = err.Error()
		res.Data.Status = http.StatusForbidden
		res.Data.Error = err.Err
		return &res
	}
	// if err.Tag == apperrors. {
	// 	res.Data.Message = "the permission is already exists"
	// 	res.Data.Status = http.StatusBadRequest
	// }

	return &res
}

func SendErrorResponse(ctx context.Context, err *apperrors.AppError, w http.ResponseWriter, l logger.Ilogger) {
	response := responseGenerator(err)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Data.Status)

	//TODO: prevent to send stack trace to the user in the production environment
	if encodeErr := json.NewEncoder(w).Encode(*response); encodeErr != nil {
		l.Debugf("failed to encode error response: %v", encodeErr)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	l.Error("an error has sent to user as a response",
		"response error:", response,
		"error", err,
	)
}
