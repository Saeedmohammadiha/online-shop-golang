package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/utils"
	internaljwt "github.com/OnlineShop/internal/pkg/jwt"
	"github.com/OnlineShop/internal/pkg/logger"
)

func AuthProtect(f func(w http.ResponseWriter, r *http.Request), u repositories.IUserRepository, l logger.Ilogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//check the header if there is a token
		authorizationHeader := r.Header.Get("Authorization")
		l.Debug("authorizationHeader", authorizationHeader)
		if authorizationHeader == "" {
			l.Debug("authorizationHeader", authorizationHeader)

		}
		authorizationHeaderArray := strings.Split(authorizationHeader, " ")
		l.Debug("authorizationHeaderArray", authorizationHeaderArray)
		token := authorizationHeaderArray[1]
		l.Debug("authorizationHeader", token)

		// decode token
		ctx := r.Context()

		claims, err := internaljwt.New(l).DecodeToken(token)
		if err != nil {
			utils.SendErrorResponse(ctx, err, w, l)
			return
		}

		// get the user data from the db
		user, err := u.GetById(claims.UserID)
		if err != nil {
			utils.SendErrorResponse(ctx, err, w, l)
			return
		}
		// check if the token
		ctx = context.WithValue(ctx, utils.USER, user)

		r = r.WithContext(ctx)
		l.Debug("the value of the user is set in the context successfully ",
			"user", user,
		)
		f(w, r)
	}

}
