package main

import (
	"github.com/OnlineShop/database"
	v1 "github.com/OnlineShop/http/v1"
	"github.com/OnlineShop/repository"
	"github.com/OnlineShop/router"
)

func main() {

	db := initDb.MysqlDatabaseConnection()

	RepositoryFactory := repository.NewRepositoryFactory(db)

	// register routes
	newRouter := router.New()
	v1.RegisterRoutes(newRouter, RepositoryFactory)

	//router.RegisterPermissionRoutes(appRouter, permissionService)

	// TODO: needs refactor
	// userRepo := repository.NewUserRepository(db)
	// roleRepo := repository.NewRoleRepository(db)

	// userService := services.NewUserService(userRepo)
	// authService := services.NewAuthService(&userRepo)
	// roleService := services.NewRoleService(roleRepo)

	// appRouter.Get("/users", userService.FindAll)
	// appRouter.Post("/users", userService.Create)
	// appRouter.Post("/roles", roleService.Create)
	// appRouter.Get("/users/{id}", userService.FindById)
	// appRouter.Put("/users/{id}", userService.Updata)
	// appRouter.Delete("/users/{id}", userService.Delete)

	// appRouter.Post("/auth/login", authService.Login)
	newRouter.Serve(":5000")

}
