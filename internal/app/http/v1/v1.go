package v1

import (
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/router"
)

func RegisterRoutes(r router.IRouter, RepositoryFactory *repositories.RepositoryFactory) {

	RegisterAuthRoutes(r, *RepositoryFactory)

	v1Router := r.CreateSubRouter("/v1/api")

	RegisterUsersRoutes(v1Router, *RepositoryFactory)
	RegisterRoleRoutes(v1Router, *RepositoryFactory)
	RegisterPermissionRoutes(v1Router, *RepositoryFactory)

}
