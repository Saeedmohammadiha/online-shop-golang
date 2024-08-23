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

func RegisterUsersRoutes(r router.IRouter, RepositoryFactory repositories.RepositoryFactory, l logger.Ilogger) {

	usersValidator := validation.NewUserValidator(l)
	usersUsecase := usecases.NewUserUsecase(RepositoryFactory.UserRepository, usersValidator, l)
	usersService := services.NewUserService(usersUsecase, l)

	usersRouter := r.CreateSubRouter("/users")
	//	usersRouter.RegisterRoute("GET", "", middlewares.AuthProtect(middlewares.GuardRoute(usersService.GetAll, []string{"readPermissions"}, l), RepositoryFactory.UserRepository, l))
	usersRouter.RegisterRoute("POST", "", middlewares.AuthProtect(middlewares.GuardRoute(usersService.Create, []string{"createPermission"}, l), RepositoryFactory.UserRepository, l))
	//	usersRouter.RegisterRoute("PUT", "", middlewares.AuthProtect(middlewares.GuardRoute(usersService.Update, []string{"updatePermission"}, l), RepositoryFactory.UserRepository, l))
	//	usersRouter.RegisterRoute("GET", "/{id}", middlewares.AuthProtect(middlewares.GuardRoute(usersService.GetById, []string{"readPermissions"}, l), RepositoryFactory.UserRepository, l))
	//	usersRouter.RegisterRoute("DELETE", "/{id}", middlewares.AuthProtect(middlewares.GuardRoute(usersService.Delete, []string{"deletePermission"}, l), RepositoryFactory.UserRepository, l))
}
