package v1

import (
	"github.com/OnlineShop/repository"
	"github.com/OnlineShop/router"
	"github.com/OnlineShop/services"
)

func RegisterRoutes(r router.IRouter, RepositoryFactory *repository.RepositoryFactory) {
	v1Router := r.CreateSubRouter("/v1/api")
	permissionService := services.NewPermissionService(RepositoryFactory.PermissionRepository)
	RegisterPermissionRoutes(v1Router, permissionService)
}
