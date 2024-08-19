package v1

import (
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/router"

	"github.com/OnlineShop/internal/pkg/logger"
)

func RegisterRoutes(r router.IRouter, RepositoryFactory *repositories.RepositoryFactory, log logger.Ilogger) {

	RegisterAuthRoutes(r, *RepositoryFactory, log)

	v1Router := r.CreateSubRouter("/v1/api")

	RegisterPermissionRoutes(v1Router, *RepositoryFactory, log)

}
