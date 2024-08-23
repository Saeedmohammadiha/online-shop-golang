package services

import (
	"net/http"

	dto "github.com/OnlineShop/internal/app/dto/roles"
	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/app/utils"
	"github.com/OnlineShop/internal/pkg/logger"
)

type IRolesService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
}

type RolesService struct {
	rolesUsecase usecases.IRolesUsecases
	log          logger.Ilogger
}

func NewRolesService(u usecases.IRolesUsecases, l logger.Ilogger) IRolesService {
	l.Info("roles service is created")
	return &RolesService{
		rolesUsecase: u,
		log:          l,
	}
}

func (s *RolesService) GetAll(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	permissions, err := s.rolesUsecase.GetAll(ctx)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
		return
	}

	utils.SendSuccessResponse(w, &permissions, s.log)

}

func (s *RolesService) GetById(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	permission, err := s.rolesUsecase.GetById(ctx)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
		return
	}

	utils.SendSuccessResponse(w, &permission, s.log)

}

func (s *RolesService) Create(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var requestBody dto.RoleCreateRequest
	if err := utils.GetValueFromCtx(ctx, utils.REQUEST_BODY, &requestBody, s.log); err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
	}

	savedRole, err := s.rolesUsecase.Create(ctx, &requestBody)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
		return
	}

	var permissionsTitle []string
	for _, v := range savedRole.Permissions {
		permissionsTitle = append(permissionsTitle, v.Title)
	}

	//convert to json and response
	responseValue := dto.RoleCreateResponse{
		ID:          savedRole.ID,
		Title:       savedRole.Title,
		Permissions: permissionsTitle,
	}

	utils.SendSuccessResponse(w, &responseValue, s.log)
}

func (s *RolesService) Update(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var requestBody dto.RoleCreateRequest
	if err := utils.GetValueFromCtx(ctx, utils.REQUEST_BODY, &requestBody, s.log); err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
	}

	updatedRole, err := s.rolesUsecase.Update(ctx, &requestBody)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
		return
	}

	var permissionsTitle []string
	for _, v := range updatedRole.Permissions {
		permissionsTitle = append(permissionsTitle, v.Title)
	}
	//convert to json and response
	responseValue := dto.RoleCreateResponse{
		ID:          updatedRole.ID,
		Title:       updatedRole.Title,
		Permissions: permissionsTitle,
	}

	utils.SendSuccessResponse(w, &responseValue, s.log)

}

func (s *RolesService) Delete(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	err := s.rolesUsecase.Delete(ctx)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
		return
	}

	utils.SendSuccessResponse(w, nil, s.log)

}
