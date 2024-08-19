package v1

import (
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/router"
	"github.com/OnlineShop/internal/app/services"
	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/app/validation"
	"github.com/OnlineShop/internal/pkg/logger"
)

func RegisterAuthRoutes(r router.IRouter, RepositoryFactory repositories.RepositoryFactory, l logger.Ilogger) {
	authValidator := validation.NewAuthValidation(l)
	authUsecase := usecases.NewAuthUsecase(RepositoryFactory.UserRepository, authValidator, l)
	authService := services.NewAuthService(authUsecase, l)

	r.RegisterRoute("POST", "/login", authService.Login)

}
