package services

import (
	"encoding/json"
	"net/http"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	dto "github.com/OnlineShop/internal/app/dto/roles"
	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/app/utils"
	"github.com/OnlineShop/internal/pkg/logger"
	"github.com/gorilla/mux"
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

func NewRolesService(u usecases.IRolesUsecases) IRolesService {
	l := logger.Logger()
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
		utils.SendErrorResponse(ctx, err, w)
		return
	}

	utils.SendSuccessResponse(w, &permissions)

}

func (s *RolesService) GetById(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	params := mux.Vars(r)
	permission, err := s.rolesUsecase.GetById(ctx,params)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w)
		return
	}

	utils.SendSuccessResponse(w, &permission)

}

func (s *RolesService) Create(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var requestBody dto.RoleCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {

		utils.SendErrorResponse(
			ctx,
			apperrors.NewBadRequestError("invalid json", err),
			w,
		)
	}

	savedRole, err := s.rolesUsecase.Create(ctx, &requestBody)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w)
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

	utils.SendSuccessResponse(w, &responseValue)
}

func (s *RolesService) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var requestBody dto.RoleCreateRequest

	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {

		utils.SendErrorResponse(
			ctx,
			apperrors.NewBadRequestError("invalid json", err),
			w,
		)
	}

	

	updatedRole, err := s.rolesUsecase.Update(ctx, &requestBody)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w)
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

	utils.SendSuccessResponse(w, &responseValue)

}

func (s *RolesService) Delete(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	params := mux.Vars(r)
	err := s.rolesUsecase.Delete(ctx,params)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w)
		return
	}

	utils.SendSuccessResponse(w, nil)

}
