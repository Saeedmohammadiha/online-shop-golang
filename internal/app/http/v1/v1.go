package v1

import (
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/router"
	"github.com/OnlineShop/internal/app/services"
	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/app/validation"
	"github.com/OnlineShop/internal/pkg/logger"
)

func RegisterRoutes(r router.IRouter, RepositoryFactory *repositories.RepositoryFactory, log logger.Ilogger) {
	v1Router := r.CreateSubRouter("/v1/api")
	permissionValidator := validation.NewPermissionValidation(log)
	permissionUsecase := usecases.NewPermissionUsecase(RepositoryFactory.PermissionRepository, permissionValidator, log)
	permissionService := services.NewPermissionService(permissionUsecase, log)
	RegisterPermissionRoutes(v1Router, permissionService)
}
