package middlewares

import (
	"net/http"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	"github.com/OnlineShop/internal/app/models"
	"github.com/OnlineShop/internal/app/utils"
)

func containsPermission(routePermissions map[string]struct{}, permission string) bool {
	_, exists := routePermissions[permission]
	return exists
}

func GuardRoute(f func(w http.ResponseWriter, r *http.Request), routePermissions []string) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		user := models.User{}
		err := utils.GetValueFromCtx(ctx, utils.USER, &user)
		if err != nil {
			err := apperrors.NewAuthorizationError("you do not have access", err)
			utils.SendErrorResponse(ctx, err, w)
			return
		}
		// Convert routePermissions to a map for O(1) lookup
		routePermissionsMap := make(map[string]struct{}, len(routePermissions))
		for _, perm := range routePermissions {
			routePermissionsMap[perm] = struct{}{}
		}

		for _, role := range user.Roles {
			for _, perm := range role.Permissions {
				if containsPermission(routePermissionsMap, perm.Title) {
					f(w, r)
				} else {
					err := apperrors.NewAuthorizationError("you do not have access", nil)
					utils.SendErrorResponse(ctx, err, w)
					return
				}
			}
		}

	}
}
