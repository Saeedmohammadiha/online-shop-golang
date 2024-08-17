package middlewares

import (
	"context"
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
	"time"

	"github.com/OnlineShop/internal/pkg/logger"
)

type BodyKey string

const KEYCON BodyKey = "requestBody"

// SanitizeMiddleware sanitizes incoming JSON requests
func (m *Middlewares) SanitizeBodyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodPut {

			var data map[string]interface{}

			if r.Body == nil {
				m.log.Info("body is empty pass next handler")
				next.ServeHTTP(w, r)
				return
			}

			// Decode the request body into the struct
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				m.log.Error("invalid json, sent error to user",
					"status code", http.StatusBadRequest,
					"error", err,
				)
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			sanitizedData := sanitizeData(&data, m.log)

			// Create a new context with sanitized data

			ctx := context.WithValue(r.Context(), KEYCON, sanitizedData)
			r = r.WithContext(ctx)
			m.log.Info("add the sanitized data to the context and pass to next handler", "body data", sanitizedData)

			// Continue with the next handler
			next.ServeHTTP(w, r)
		} else {

			next.ServeHTTP(w, r)
		}

	})
}

// sanitizeMap sanitizes values in a map.
func sanitizeData(data *map[string]interface{}, l logger.Ilogger) map[string]interface{} {
	sanitized := make(map[string]interface{})
	for key, value := range *data {
		sanitizedKey := template.HTMLEscapeString(key)
		sanitizedValue := sanitizeValue(value, l)
		sanitized[sanitizedKey] = sanitizedValue
		l.Info("sanitizing the property", key, sanitizedKey)
		l.Info("sanitizing the value", value, sanitizedValue)
	}
	return sanitized
}

// sanitizeValue sanitizes individual values in the map.
func sanitizeValue(value interface{}, l logger.Ilogger) interface{} {
	switch v := value.(type) {
	case string:
		return sanitizeString(v, l)
	case map[string]interface{}:
		return sanitizeData(&v, l)
	case []interface{}:
		sanitizedArray := make([]interface{}, len(v))
		for i, elem := range v {
			sanitizedArray[i] = sanitizeValue(elem, l)
		}
		return sanitizedArray
	case int:
		return sanitizeInt(v, l)
	case float64: // JSON numbers are decoded as float64
		return sanitizeInt(int(v), l)
	case time.Time:
		return sanitizeTime(v, l)
	default:
		// Return the value as is if it's not a type we handle
		return value
	}
}

// sanitizeString escapes HTML and trims whitespace from the input string
func sanitizeString(input string, l logger.Ilogger) string {
	// Escape HTML entities
	escaped := template.HTMLEscapeString(input)

	l.Info("sanitizing string", input, escaped)
	// Trim leading and trailing whitespace
	return strings.TrimSpace(escaped)
}

// sanitizeInt sanitizes integer values.
func sanitizeInt(input int, l logger.Ilogger) int {
	// Example: Clamp negative values to 0 and handle other validations as needed
	l.Info("sanitizing int", input)
	if input < 0 {
		return 0
	}
	l.Info("sanitized int", input)
	return input
}

// sanitizeTime formats and validates date/time values.
func sanitizeTime(input time.Time, l logger.Ilogger) time.Time {
	// Example: Clamp future dates to the current time
	if input.After(time.Now()) {
		return time.Now()
	}
	l.Info("sanitized time", input)
	return input
}
