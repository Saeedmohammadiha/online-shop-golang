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
	permissionRouter.RegisterRoute("GET", "", middlewares.AuthProtect(permissionService.GetAll, RepositoryFactory.UserRepository, l))
	permissionRouter.RegisterRoute("POST", "", middlewares.AuthProtect(permissionService.Create, RepositoryFactory.UserRepository, l))
	permissionRouter.RegisterRoute("PUT", "", middlewares.AuthProtect(permissionService.Update, RepositoryFactory.UserRepository, l))
	permissionRouter.RegisterRoute("GET", "/{id}", middlewares.AuthProtect(permissionService.GetById, RepositoryFactory.UserRepository, l))
	permissionRouter.RegisterRoute("DELETE", "/{id}", middlewares.AuthProtect(permissionService.Delete, RepositoryFactory.UserRepository, l))
}
