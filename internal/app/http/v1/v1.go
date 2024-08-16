package v1

import (
	"github.com/OnlineShop/internal/app/logger"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/router"
	"github.com/OnlineShop/internal/app/services"
)

func RegisterRoutes(r router.IRouter, RepositoryFactory *repositories.RepositoryFactory, log logger.Ilogger) {
	v1Router := r.CreateSubRouter("/v1/api")
	permissionService := services.NewPermissionService(RepositoryFactory.PermissionRepository, log)
	RegisterPermissionRoutes(v1Router, permissionService)
}
