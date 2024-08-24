package v1

import (
	"github.com/OnlineShop/internal/app/middlewares"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/router"
	"github.com/OnlineShop/internal/app/services"
	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/app/validation"
	internaljwt "github.com/OnlineShop/internal/pkg/jwt"
)

func RegisterAuthRoutes(r router.IRouter, RepositoryFactory repositories.RepositoryFactory) {
	j := internaljwt.New()
	authValidator := validation.NewAuthValidation()
	authUsecase := usecases.NewAuthUsecase(RepositoryFactory.UserRepository, authValidator, j)
	authService := services.NewAuthService(authUsecase)

	r.RegisterRoute("POST", "/login", authService.Login)
	r.RegisterRoute("POST", "/logout", middlewares.AuthProtect(authService.Logout, RepositoryFactory.UserRepository))
	r.RegisterRoute("POST", "/refresh", authService.Refresh)

}
