package middlewares

import (
	"context"
	"html/template"
	"net/http"

	"github.com/OnlineShop/internal/app/utils"
)

// SanitizeURLParamsMiddleware sanitizes URL query parameters and adds them to the context.
func (m *Middlewares) SanitizeURLParamsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Sanitize query parameters
		sanitizedQueryParams := make(map[string][]string)
		for key, values := range r.URL.Query() {
			sanitizedKey := template.HTMLEscapeString(key)
			sanitizedValues := make([]string, len(values))
			for i, value := range values {
				sanitizedValues[i] = template.HTMLEscapeString(value)
				m.log.Infof("sanitize the param %s , to the %s", value, sanitizedValues[i])
			}
			sanitizedQueryParams[sanitizedKey] = sanitizedValues
		}

		// Enrich context with sanitized parameters
		ctx := context.WithValue(r.Context(), utils.REQUEST_PARAMS, sanitizedQueryParams)
		r = r.WithContext(ctx)
		m.log.Info("add the sanitized data to the context and pass to next handler", "params", sanitizedQueryParams)

		// Continue with the next handler
		next.ServeHTTP(w, r)
	})
}
