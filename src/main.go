package main

import (
	initDb "github.com/OnlineShop/src/database"
	v1 "github.com/OnlineShop/src/http/v1"
	"github.com/OnlineShop/src/repository"
	"github.com/OnlineShop/src/router"
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
