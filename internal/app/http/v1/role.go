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

func RegisterRoleRoutes(r router.IRouter, RepositoryFactory repositories.RepositoryFactory, l logger.Ilogger) {

	rolesValidator := validation.NewRolesValidation(l)
	rolesUsecase := usecases.NewRolesUsecases(RepositoryFactory.RolesRepository, rolesValidator, l)
	rolesService := services.NewRolesService(rolesUsecase, l)

	rolesRouter := r.CreateSubRouter("/roles")
	//	rolesRouter.RegisterRoute("GET", "", middlewares.AuthProtect(middlewares.GuardRoute(rolesService.GetAll, []string{"readPermissions"}, l), RepositoryFactory.UserRepository, l))
	rolesRouter.RegisterRoute("POST", "", middlewares.AuthProtect(middlewares.GuardRoute(rolesService.Create, []string{"createPermission"}, l), RepositoryFactory.UserRepository, l))
	//	rolesRouter.RegisterRoute("PUT", "", middlewares.AuthProtect(middlewares.GuardRoute(rolesService.Update, []string{"updatePermission"}, l), RepositoryFactory.UserRepository, l))
	//	rolesRouter.RegisterRoute("GET", "/{id}", middlewares.AuthProtect(middlewares.GuardRoute(rolesService.GetById, []string{"readPermissions"}, l), RepositoryFactory.UserRepository, l))
	//	rolesRouter.RegisterRoute("DELETE", "/{id}", middlewares.AuthProtect(middlewares.GuardRoute(rolesService.Delete, []string{"deletePermission"}, l), RepositoryFactory.UserRepository, l))
}
