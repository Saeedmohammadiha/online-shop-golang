package main

import (
	"github.com/OnlineShop/database"
	"github.com/OnlineShop/repository"
	"github.com/OnlineShop/router"
	"github.com/OnlineShop/services"
)

func main() {

	db := databaseConfig.MysqlDatabaseConnection()

	userRepo := repository.NewUserRepository(db)
	roleRepo := repository.NewRoleRepository(db)

	userService := services.NewUserService(userRepo)
	authService := services.NewAuthService(userRepo)
	roleService := services.NewRoleService(roleRepo)
	router := router.NewRouter()

	router.Get("/users", userService.FindAll)
	router.Post("/users", userService.Create)
	router.Post("/roles", roleService.Create)
	router.Get("/users/{id}", userService.FindById)
	router.Put("/users/{id}", userService.Updata)
	router.Delete("/users/{id}", userService.Delete)

	router.Post("/auth/login", authService.Login)
	router.Serve(":5000")

}
