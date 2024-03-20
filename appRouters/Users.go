package approuters

import (
	"github.com/OnlineShop/services"
	"github.com/gorilla/mux"
)

func UserRouter(api *mux.Route, userService *services.UserService) {
	userRouter := api.PathPrefix("/users").Subrouter()

	userRouter.HandleFunc("", userService.Create).Methods("POST")

}
