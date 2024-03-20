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
	router.Post("/user", userService.Create)

	router.Serve(":5000")

}
