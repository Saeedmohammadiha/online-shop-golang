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

	userService := services.NewUserService(userRepo)
	router := router.NewRouter()

	router.Get("/users", userService.FindAll)
	router.Post("/users", userService.Create)
	router.Get("/users/{id}", userService.FindById)
	router.Put("/users/{id}", userService.Updata)
	router.Delete("/users/{id}", userService.Delete)
	router.Serve(":5000")

}
