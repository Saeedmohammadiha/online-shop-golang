package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/OnlineShop/database"
	models "github.com/OnlineShop/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type test struct {
	Test string
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	data := test{Test: "a"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func userHandler(db *gorm.DB, userRouter *mux.Router) *gorm.DB {
	result := db.Find(&models.User{})
	userRouter.HandleFunc("", GetUsers).Methods("GET")

	return result
}

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

	router := mux.NewRouter()

	api := router.PathPrefix("/v1/api")
	userRouter := api.PathPrefix("/users").Subrouter()
	userHandler(db, userRouter)

	fmt.Println("server is listening on port 5000")

	log.Fatal(http.ListenAndServe(":5000", router))
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
