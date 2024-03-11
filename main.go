package main

import (
	"fmt"
	"log"
	"net/http"


	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gorilla/mux"
)






func getUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("users")
}

func main() {
	//create a server
	//crearte a router
	//add db
	//
	dsn := "root:S@eed1372144@127.0.0.1:3306/rrr?charset=utf8mb4&parseTime=True&loc=Local"
	_, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("errrr", err)
	}

	router := mux.NewRouter()

	api := router.PathPrefix("/v1/api")
	userRouter := api.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("", getUsers).Methods("GET")

	fmt.Println("server is listening on port 5000")

	log.Fatal(http.ListenAndServe(":5000", router))
	// if err != nil {
	// 	fmt.Println(err)
	// }
}
