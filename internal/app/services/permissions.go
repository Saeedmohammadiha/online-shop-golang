package services

import (
	"encoding/json"
	"net/http"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	permissionDto "github.com/OnlineShop/internal/app/dto/permissions"
	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/pkg/logger"
	"github.com/gorilla/mux"

	"github.com/OnlineShop/internal/app/utils"
)

type IPermissionService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
}

type PermissionService struct {
	permissionUsecase usecases.IPermissionUsecases
	log               logger.Ilogger
}

func NewPermissionService(u usecases.IPermissionUsecases) IPermissionService {
	l := logger.Logger()
	l.Info("permission service is created")
	return &PermissionService{
		permissionUsecase: u,
		log:               l,
	}
}

func (s *PermissionService) GetAll(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	permissions, err := s.permissionUsecase.GetAll(ctx)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w)
		return
	}

	utils.SendSuccessResponse(w, &permissions)

}

func (s *PermissionService) Create(w http.ResponseWriter, r *http.Request) {

	
	ctx := r.Context()

	var requestBody permissionDto.PermissionCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {

		utils.SendErrorResponse(
			ctx,
			apperrors.NewBadRequestError("invalid json", err),
			w,
		)
	}

	savedPermission, err := s.permissionUsecase.Create(ctx, &requestBody)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w)
		return
	}

	//convert to json and response
	responseValue := permissionDto.PermissionCreateResponse{
		ID:    savedPermission.ID,
		Title: savedPermission.Title,
	}

	utils.SendSuccessResponse(w, &responseValue)

}

func (s *PermissionService) GetById(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	params := mux.Vars(r)

	permission, err := s.permissionUsecase.GetById(ctx, params)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w)
		return
	}

	utils.SendSuccessResponse(w, &permission)

}

func (s *PermissionService) Update(w http.ResponseWriter, r *http.Request) {

	
	ctx := r.Context()

	var requestBody permissionDto.PermissionCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {

		utils.SendErrorResponse(
			ctx,
			apperrors.NewBadRequestError("invalid json", err),
			w,
		)
	}

	savedPermission, err := s.permissionUsecase.Update(ctx, &requestBody)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w)
		return
	}

	//convert to json and response
	responseValue := permissionDto.PermissionCreateResponse{
		ID:    savedPermission.ID,
		Title: savedPermission.Title,
	}

	utils.SendSuccessResponse(w, &responseValue)

}

func (s *PermissionService) Delete(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	params := mux.Vars(r)
	err := s.permissionUsecase.Delete(ctx, params)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w)
		return
	}

	utils.SendSuccessResponse(w, nil)

}
