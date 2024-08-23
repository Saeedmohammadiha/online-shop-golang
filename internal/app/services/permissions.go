package services

import (
	"net/http"

	permissionDto "github.com/OnlineShop/internal/app/dto/permissions"
	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/pkg/logger"

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

func NewPermissionService(u usecases.IPermissionUsecases, l logger.Ilogger) IPermissionService {
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
		utils.SendErrorResponse(ctx, err, w, s.log)
		return
	}

	utils.SendSuccessResponse(w, &permissions, s.log)

}

func (s *PermissionService) Create(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var requestBody permissionDto.PermissionCreateRequest
	if err := utils.GetValueFromCtx(ctx, utils.REQUEST_BODY, &requestBody, s.log); err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
	}

	savedPermission, err := s.permissionUsecase.Create(ctx, &requestBody)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
		return
	}

	//convert to json and response
	responseValue := permissionDto.PermissionCreateResponse{
		ID:    savedPermission.ID,
		Title: savedPermission.Title,
	}

	utils.SendSuccessResponse(w, &responseValue, s.log)

}

func (s *PermissionService) GetById(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	permission, err := s.permissionUsecase.GetById(ctx)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
		return
	}

	utils.SendSuccessResponse(w, &permission, s.log)

}

func (s *PermissionService) Update(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var requestBody permissionDto.PermissionCreateRequest
	if err := utils.GetValueFromCtx(ctx, utils.REQUEST_BODY, &requestBody, s.log); err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
	}

	savedPermission, err := s.permissionUsecase.Update(ctx, &requestBody)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
		return
	}

	//convert to json and response
	responseValue := permissionDto.PermissionCreateResponse{
		ID:    savedPermission.ID,
		Title: savedPermission.Title,
	}

	utils.SendSuccessResponse(w, &responseValue, s.log)

}

func (s *PermissionService) Delete(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	err := s.permissionUsecase.Delete(ctx)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
		return
	}

	utils.SendSuccessResponse(w, nil, s.log)
	
}
