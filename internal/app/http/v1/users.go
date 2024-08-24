package v1

import (
	"github.com/OnlineShop/internal/app/middlewares"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/router"
	"github.com/OnlineShop/internal/app/services"
	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/app/validation"
)

func RegisterUsersRoutes(r router.IRouter, RepositoryFactory repositories.RepositoryFactory, ) {

	usersValidator := validation.NewUserValidator()
	usersUsecase := usecases.NewUserUsecase(RepositoryFactory.UserRepository, usersValidator)
	usersService := services.NewUserService(usersUsecase)

	usersRouter := r.CreateSubRouter("/users")
	usersRouter.RegisterRoute("GET", "", middlewares.AuthProtect(middlewares.GuardRoute(usersService.GetAll, []string{"readPermissions"}), RepositoryFactory.UserRepository))
	usersRouter.RegisterRoute("POST", "", middlewares.AuthProtect(middlewares.GuardRoute(usersService.Create, []string{"createPermission"}), RepositoryFactory.UserRepository))
	usersRouter.RegisterRoute("PUT", "", middlewares.AuthProtect(middlewares.GuardRoute(usersService.Update, []string{"updatePermission"}), RepositoryFactory.UserRepository))
	usersRouter.RegisterRoute("GET", "/{id}", middlewares.AuthProtect(middlewares.GuardRoute(usersService.GetById, []string{"readPermissions"}), RepositoryFactory.UserRepository))
	usersRouter.RegisterRoute("DELETE", "/{id}", middlewares.AuthProtect(middlewares.GuardRoute(usersService.Delete, []string{"deletePermission"}), RepositoryFactory.UserRepository))
}
