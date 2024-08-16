package middlewares

import (
	"context"
	"encoding/json"
	"html/template"
	"net/http"
	"strings"
	"time"
)

// SanitizeMiddleware sanitizes incoming JSON requests
func SanitizeBodyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost || r.Method == http.MethodPut {

			var data map[string]interface{}

			if r.Body == nil {
				next.ServeHTTP(w, r)
				return
			}

			// Decode the request body into the struct
			if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
				http.Error(w, "Invalid JSON", http.StatusBadRequest)
				return
			}

			sanitizedData := sanitizeData(&data)

			// Create a new context with sanitized data
			type BodyKey string
			const bodyKey BodyKey = "requestBody"
			ctx := context.WithValue(r.Context(), bodyKey, sanitizedData)
			r = r.WithContext(ctx)

			// Continue with the next handler
			next.ServeHTTP(w, r)
		}

		next.ServeHTTP(w, r)
	})
}

// sanitizeData sanitizes fields of the struct.
// sanitizeMap sanitizes values in a map.
func sanitizeData(data *map[string]interface{}) map[string]interface{} {
	sanitized := make(map[string]interface{})
	for key, value := range *data {
		sanitizedKey := template.HTMLEscapeString(key)
		sanitizedValue := sanitizeValue(value)
		sanitized[sanitizedKey] = sanitizedValue
	}
	return sanitized
}

// sanitizeValue sanitizes individual values in the map.
func sanitizeValue(value interface{}) interface{} {
	switch v := value.(type) {
	case string:
		return sanitizeString(v)
	case map[string]interface{}:
		return sanitizeData(&v)
	case []interface{}:
		sanitizedArray := make([]interface{}, len(v))
		for i, elem := range v {
			sanitizedArray[i] = sanitizeValue(elem)
		}
		return sanitizedArray
	case int:
		return sanitizeInt(v)
	case float64: // JSON numbers are decoded as float64
		return sanitizeInt(int(v))
	case time.Time:
		return sanitizeTime(v)
	default:
		// Return the value as is if it's not a type we handle
		return value
	}
}

// sanitizeString escapes HTML and trims whitespace from the input string
func sanitizeString(input string) string {
	// Escape HTML entities
	escaped := template.HTMLEscapeString(input)
	// Trim leading and trailing whitespace
	return strings.TrimSpace(escaped)
}

// sanitizeInt sanitizes integer values.
func sanitizeInt(input int) int {
	// Example: Clamp negative values to 0 and handle other validations as needed
	if input < 0 {
		return 0
	}
	return input
}

// sanitizeTime formats and validates date/time values.
func sanitizeTime(input time.Time) time.Time {
	// Example: Clamp future dates to the current time
	if input.After(time.Now()) {
		return time.Now()
	}
	return input
}
