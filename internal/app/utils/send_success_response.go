package utils

import (
	"encoding/json"
	"net/http"

	"github.com/OnlineShop/internal/pkg/logger"
)

func SendSuccessResponse(w http.ResponseWriter, data interface{}, l logger.Ilogger) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(data); err != nil {
		l.Error("failed to parse json to send response", "status code", http.StatusInternalServerError)
		http.Error(w, "failed to parse json", http.StatusInternalServerError)
	}
	l.Info("the response has sent to user",
		"status code", http.StatusOK,
		"response", data,
	)
}
