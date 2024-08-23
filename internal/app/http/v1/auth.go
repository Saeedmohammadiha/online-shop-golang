package v1

import (
	"github.com/OnlineShop/internal/app/middlewares"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/router"
	"github.com/OnlineShop/internal/app/services"
	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/app/validation"
	internaljwt "github.com/OnlineShop/internal/pkg/jwt"
	"github.com/OnlineShop/internal/pkg/logger"
)

func RegisterAuthRoutes(r router.IRouter, RepositoryFactory repositories.RepositoryFactory, l logger.Ilogger) {
	j := internaljwt.New(l)
	authValidator := validation.NewAuthValidation(l)
	authUsecase := usecases.NewAuthUsecase(RepositoryFactory.UserRepository, authValidator, j, l)
	authService := services.NewAuthService(authUsecase, l)

	r.RegisterRoute("POST", "/login", authService.Login)
	r.RegisterRoute("POST", "/logout", middlewares.AuthProtect(authService.Logout, RepositoryFactory.UserRepository, l))
	r.RegisterRoute("POST", "/refresh", authService.Refresh)

}
