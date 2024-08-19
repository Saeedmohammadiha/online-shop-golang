package v1

import (
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
	permissionRouter.RegisterRoute("GET", "", permissionService.FindAll)
	permissionRouter.RegisterRoute("POST", "", permissionService.Create)
	permissionRouter.RegisterRoute("PUT", "", permissionService.Update)
	permissionRouter.RegisterRoute("GET", "/{id}", permissionService.FindById)
	permissionRouter.RegisterRoute("DELETE", "/{id}", permissionService.Delete)
}
