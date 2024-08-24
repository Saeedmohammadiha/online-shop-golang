package middlewares

import (
	"net/http"
)

type Imiddlewares interface {
	SanitizeBodyMiddleware(next http.Handler) http.Handler
	SanitizeURLParamsMiddleware(next http.Handler) http.Handler
}

type Middlewares struct{}

func New() Imiddlewares {
	return &Middlewares{}
}
