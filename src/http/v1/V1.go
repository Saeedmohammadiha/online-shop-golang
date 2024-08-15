package v1

import (
	"github.com/OnlineShop/src/repository"
	"github.com/OnlineShop/src/router"
	"github.com/OnlineShop/src/services"
)

func RegisterRoutes(r router.IRouter, RepositoryFactory *repository.RepositoryFactory) {
	v1Router := r.CreateSubRouter("/v1/api")
	permissionService := services.NewPermissionService(RepositoryFactory.PermissionRepository)
	RegisterPermissionRoutes(v1Router, permissionService)
}
