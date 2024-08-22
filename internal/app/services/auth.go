package services

import (
	"net/http"

	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/app/utils"

	"github.com/OnlineShop/internal/pkg/logger"
)

type IAuthService interface {
	Login(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
}
type AuthService struct {
	authUsecase usecases.IAuthUsecases

	log logger.Ilogger
}

func NewAuthService(u usecases.IAuthUsecases, l logger.Ilogger) IAuthService {
	l.Info("new Auth service is created")
	return &AuthService{
		authUsecase: u,
		log:         l,
	}
}

func (a *AuthService) Login(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	responseData, err := a.authUsecase.Login(ctx)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w, a.log)
		return
	}

	utils.SendSuccessResponse(w, responseData, a.log)

}

func (a *AuthService) Refresh(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	responseData, err := a.authUsecase.Refresh(ctx)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w, a.log)
		return
	}

	utils.SendSuccessResponse(w, responseData, a.log)
	//TODO: add the invalide tokens in the logout service and if the user changed his password
}
