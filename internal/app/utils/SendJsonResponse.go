package utils

import (
	"encoding/json"
	"net/http"

	"github.com/OnlineShop/internal/pkg/logger"
)

// func SendJsonResponse(w http.ResponseWriter, data interface{}, l logger.Ilogger) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)

// 	if err := json.NewEncoder(w).Encode(data); err != nil {
// 		l.Error("failed to parse json to send response", "status code", http.StatusInternalServerError)
// 		http.Error(w, "failed to parse json", http.StatusInternalServerError)
// 	}
// 	l.Info("the response has sent to user",
// 		"status code", http.StatusOK,
// 		"response", data,
// 	)
// }

func SendJsonResponse(w http.ResponseWriter, data interface{}, l logger.Ilogger) {
	// Set the content type for the response
	w.Header().Set("Content-Type", "application/json")

	// Encode the response data to JSON
	if err := json.NewEncoder(w).Encode(data); err != nil {
		// Log the error
		l.Error("failed to parse json to send response", "error", err)

		// Since we can't send another response, just log the situation
		return
	}

	// Log that the response was successfully sent
	l.Info("the response has been sent to the user",
		"status code", http.StatusOK,
		"response", data,
	)
}