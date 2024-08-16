package middlewares

import (
	"net/http"

	"github.com/OnlineShop/internal/pkg/logger"
)

type Imiddlewares interface {
	SanitizeBodyMiddleware(next http.Handler) http.Handler
	SanitizeURLParamsMiddleware(next http.Handler) http.Handler
}

type Middlewares struct {
	log logger.Ilogger
}

func New(l logger.Ilogger) Imiddlewares {
	return &Middlewares{
		log: l,
	}
}
