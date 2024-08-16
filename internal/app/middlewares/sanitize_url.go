package middlewares

import (
	"context"
	"html/template"
	"net/http"
)

// SanitizeURLParamsMiddleware sanitizes URL query parameters and adds them to the context.
func SanitizeURLParamsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Sanitize query parameters
		sanitizedQueryParams := make(map[string][]string)
		for key, values := range r.URL.Query() {
			sanitizedKey := template.HTMLEscapeString(key)
			sanitizedValues := make([]string, len(values))
			for i, value := range values {
				sanitizedValues[i] = template.HTMLEscapeString(value)
			}
			sanitizedQueryParams[sanitizedKey] = sanitizedValues
		}

		// Enrich context with sanitized parameters
		type ParamsKey string
		const paramsKey ParamsKey = "params"
		ctx := context.WithValue(r.Context(), paramsKey, sanitizedQueryParams)
		r = r.WithContext(ctx)

		// Continue with the next handler
		next.ServeHTTP(w, r)
	})
}
