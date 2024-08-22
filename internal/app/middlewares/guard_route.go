package middlewares

import (
	"net/http"

	"github.com/OnlineShop/internal/app/models"
	"github.com/OnlineShop/internal/app/utils"
	"github.com/OnlineShop/internal/pkg/logger"
)

func GuardRoute(f func(w http.ResponseWriter, r *http.Request), routePernissions []string, l logger.Ilogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		user := models.User{}
		utils.GetValueFromCtx(ctx, utils.USER, &user, l)

		// check the roles of the user and the permissions of that role
		//compare the provided permissions to the permissions of the user's role

		// if the there was the permission pass the function
		//if there is no permission just send the 403 response
		f(w, r)
	}
}
