package v1

import (
	"github.com/OnlineShop/internal/app/middlewares"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/router"
	"github.com/OnlineShop/internal/app/services"
	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/app/validation"
	"github.com/OnlineShop/internal/pkg/logger"
)

func RegisterPermissionRoutes(r router.IRouter, RepositoryFactory repositories.RepositoryFactory, l logger.Ilogger) {

	permissionValidator := validation.NewPermissionValidation(l)
	permissionUsecase := usecases.NewPermissionUsecase(RepositoryFactory.PermissionRepository, permissionValidator, l)
	permissionService := services.NewPermissionService(permissionUsecase, l)

	permissionRouter := r.CreateSubRouter("/permissions")
	permissionRouter.RegisterRoute("GET", "", middlewares.AuthProtect(middlewares.GuardRoute(permissionService.GetAll, []string{"readPermissions"}, l), RepositoryFactory.UserRepository, l))
	permissionRouter.RegisterRoute("POST", "", middlewares.AuthProtect(middlewares.GuardRoute(permissionService.Create, []string{"createPermission"}, l), RepositoryFactory.UserRepository, l))
	permissionRouter.RegisterRoute("PUT", "", middlewares.AuthProtect(middlewares.GuardRoute(permissionService.Update, []string{"updatePermission"}, l), RepositoryFactory.UserRepository, l))
	permissionRouter.RegisterRoute("GET", "/{id}", middlewares.AuthProtect(middlewares.GuardRoute(permissionService.GetById, []string{"readPermissions"}, l), RepositoryFactory.UserRepository, l))
	permissionRouter.RegisterRoute("DELETE", "/{id}", middlewares.AuthProtect(middlewares.GuardRoute(permissionService.Delete, []string{"deletePermission"}, l), RepositoryFactory.UserRepository, l))
}
