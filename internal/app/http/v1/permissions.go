package v1

import (
	"github.com/OnlineShop/internal/app/middlewares"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/router"
	"github.com/OnlineShop/internal/app/services"
	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/app/validation"
)

func RegisterPermissionRoutes(r router.IRouter, RepositoryFactory repositories.RepositoryFactory) {

	permissionValidator := validation.NewPermissionValidation()
	permissionUsecase := usecases.NewPermissionUsecase(RepositoryFactory.PermissionRepository, permissionValidator)
	permissionService := services.NewPermissionService(permissionUsecase)

	permissionRouter := r.CreateSubRouter("/permissions")
	permissionRouter.RegisterRoute("GET", "", middlewares.AuthProtect(middlewares.GuardRoute(permissionService.GetAll, []string{"readPermissions"}), RepositoryFactory.UserRepository))
	permissionRouter.RegisterRoute("POST", "", middlewares.AuthProtect(middlewares.GuardRoute(permissionService.Create, []string{"createPermission"}), RepositoryFactory.UserRepository))
	permissionRouter.RegisterRoute("PUT", "", middlewares.AuthProtect(middlewares.GuardRoute(permissionService.Update, []string{"updatePermission"}), RepositoryFactory.UserRepository))
	permissionRouter.RegisterRoute("GET", "/{id}", middlewares.AuthProtect(middlewares.GuardRoute(permissionService.GetById, []string{"readPermissions"}), RepositoryFactory.UserRepository))
	permissionRouter.RegisterRoute("DELETE", "/{id}", middlewares.AuthProtect(middlewares.GuardRoute(permissionService.Delete, []string{"deletePermission"}), RepositoryFactory.UserRepository))
}
