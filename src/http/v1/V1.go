package v1

import (
	"github.com/OnlineShop/src/logger"
	"github.com/OnlineShop/src/repository"
	"github.com/OnlineShop/src/router"
	"github.com/OnlineShop/src/services"
)

func RegisterRoutes(r router.IRouter, RepositoryFactory *repository.RepositoryFactory, log logger.Ilogger) {
	v1Router := r.CreateSubRouter("/v1/api")
	permissionService := services.NewPermissionService(RepositoryFactory.PermissionRepository, log)
	RegisterPermissionRoutes(v1Router, permissionService)
}
