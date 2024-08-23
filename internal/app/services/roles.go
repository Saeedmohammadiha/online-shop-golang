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

func (s *RolesService) Create(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var requestBody dto.RoleCreateRequest
	if err := utils.GetValueFromCtx(ctx, utils.REQUEST_BODY, &requestBody, s.log); err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
	}

	savedRole, err := s.rolesUsecase.Create(&requestBody)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
		return
	}

	//convert to json and response
	responseValue := dto.RoleCreateResponse{
		ID:    savedRole.ID,
		Title: savedRole.Title,
	}

	utils.SendSuccessResponse(w, &responseValue, s.log)
}
