package v1

import (
	"github.com/OnlineShop/internal/app/middlewares"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/router"
	"github.com/OnlineShop/internal/app/services"
	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/app/validation"
)

func RegisterRoleRoutes(r router.IRouter, RepositoryFactory repositories.RepositoryFactory) {

	rolesValidator := validation.NewRolesValidation()
	rolesUsecase := usecases.NewRolesUsecases(RepositoryFactory.RolesRepository, RepositoryFactory.PermissionRepository, rolesValidator)
	rolesService := services.NewRolesService(rolesUsecase, )

	rolesRouter := r.CreateSubRouter("/roles")
	rolesRouter.RegisterRoute("GET", "", middlewares.AuthProtect(middlewares.GuardRoute(rolesService.GetAll, []string{"readPermissions"}), RepositoryFactory.UserRepository))
	rolesRouter.RegisterRoute("POST", "", middlewares.AuthProtect(middlewares.GuardRoute(rolesService.Create, []string{"createPermission"}), RepositoryFactory.UserRepository))
	rolesRouter.RegisterRoute("PUT", "", middlewares.AuthProtect(middlewares.GuardRoute(rolesService.Update, []string{"updatePermission"}), RepositoryFactory.UserRepository))
	rolesRouter.RegisterRoute("GET", "/{id}", middlewares.AuthProtect(middlewares.GuardRoute(rolesService.GetById, []string{"readPermissions"}), RepositoryFactory.UserRepository))
	rolesRouter.RegisterRoute("DELETE", "/{id}", middlewares.AuthProtect(middlewares.GuardRoute(rolesService.Delete, []string{"deletePermission"}), RepositoryFactory.UserRepository))
}
