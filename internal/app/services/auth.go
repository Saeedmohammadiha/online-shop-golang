package services

import (
	"context"
	"encoding/json"
	"net/http"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	dto "github.com/OnlineShop/internal/app/dto/auth"
	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/app/utils"

	"github.com/OnlineShop/internal/pkg/logger"
)

type IAuthService interface {
	Login(w http.ResponseWriter, r *http.Request)
	Logout(w http.ResponseWriter, r *http.Request)
	Refresh(w http.ResponseWriter, r *http.Request)
}
type AuthService struct {
	authUsecase usecases.IAuthUsecases
	log         logger.Ilogger
}

func NewAuthService(u usecases.IAuthUsecases) IAuthService {
	l := logger.Logger()
	l.Info("new Auth service is created")
	return &AuthService{
		authUsecase: u,
		log:         l,
	}
}

func (a *AuthService) Logout(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	response, err := a.authUsecase.Logout(ctx)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w)
		return
	}

	utils.SendSuccessResponse(w, response)

}

func (a *AuthService) Login(w http.ResponseWriter, r *http.Request) {

	ctx := context.WithValue(r.Context(), utils.REQUEST_BODY, r.Body)
	var requestBody dto.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {

		utils.SendErrorResponse(
			ctx,
			apperrors.NewBadRequestError("invalid json", err),
			w,
		)
	}

	responseData, err := a.authUsecase.Login(ctx, &requestBody)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w)
		return
	}

	utils.SendSuccessResponse(w, responseData)

}

func (a *AuthService) Refresh(w http.ResponseWriter, r *http.Request) {

	ctx := context.WithValue(r.Context(), utils.REQUEST_BODY, r.Body)
	var requestBody dto.RefreshRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {

		utils.SendErrorResponse(
			ctx,
			apperrors.NewBadRequestError("invalid json", err),
			w,
		)
	}
	responseData, err := a.authUsecase.Refresh(ctx, &requestBody)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w)
		return
	}

	utils.SendSuccessResponse(w, responseData)
	//TODO: add the invalide tokens in the logout service and if the user changed his password
}
