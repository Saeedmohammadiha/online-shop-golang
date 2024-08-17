package utils

import (
	"context"
	"encoding/json"
	"net/http"

	dto "github.com/OnlineShop/internal/app/dto/error"
	"github.com/OnlineShop/internal/pkg/logger"
)

func SendErrorResponse(ctx context.Context, w http.ResponseWriter, err *dto.Error, l logger.Ilogger) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(err.Data.Status)

	//TODO: prevent to send stack trace to the user in the production environment
	json.NewEncoder(w).Encode(err)
	l.Error("an error has sent to user as a response", "error", err)
}
