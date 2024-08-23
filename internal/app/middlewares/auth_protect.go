package middlewares

import (
	"context"

	"net/http"
	"strings"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/utils"
	internaljwt "github.com/OnlineShop/internal/pkg/jwt"
	"github.com/OnlineShop/internal/pkg/logger"
)

func AuthProtect(f func(w http.ResponseWriter, r *http.Request), u repositories.IUserRepository, l logger.Ilogger) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		//check the header if there is a token
		ctx := r.Context()
		authorizationHeader := r.Header.Get("Authorization")
		l.Debug("authorizationHeader", authorizationHeader)
		if authorizationHeader == "" {
			l.Debug("authorizationHeader", authorizationHeader)
			err := apperrors.NewAuthenticationError("the token is not provided", nil)
			utils.SendErrorResponse(ctx, err, w, l)
			return
		}
		authorizationHeaderArray := strings.Split(authorizationHeader, " ")
		l.Debug("authorizationHeaderArray", authorizationHeaderArray)
		token := authorizationHeaderArray[1]
		l.Debug("authorizationHeader", token)
		ctx = context.WithValue(ctx, utils.TOKEN, token)
		// decode token

		claims, err := internaljwt.New(l).DecodeToken(token)
		if err != nil {
			utils.SendErrorResponse(ctx, err, w, l)
			return
		}

		// get the user data from the db
		user, err := u.GetById(ctx, claims.UserID)
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
