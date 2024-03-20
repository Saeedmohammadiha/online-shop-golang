package main

import (
	"fmt"
	"log"
	"net/http"

	approuters "github.com/OnlineShop/appRouters"
	"github.com/OnlineShop/database"
	models "github.com/OnlineShop/models"
	"github.com/OnlineShop/repository"
	"github.com/OnlineShop/services"
	"github.com/gorilla/mux"
)

func main() {

	db := databaseConfig.MysqlDatabaseConnection()
	db.AutoMigrate(&models.User{},
		&models.Product{},
		&models.Order{},
		&models.Address{},
		&models.Comment{},
		&models.Discount{},
		&models.OrderItem{},
		&models.OrderStatus{},
		&models.Permission{},
		&models.Role{},
		&models.Score{},
		&models.TransactionStatus{},
		&models.Transaction{},
	)
	userRepo := repository.NewUserRepository(db)
	userService:=services.NewUserService(userRepo)
	router := mux.NewRouter()

	api := router.PathPrefix("/v1/api")
	approuters.UserRouter(api, userService)

	fmt.Println("server is listening on port 5000")

	log.Fatal(http.ListenAndServe(":5000", router))
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
