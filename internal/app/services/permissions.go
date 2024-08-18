package services

import (
	"errors"
	"net/http"

	dto "github.com/OnlineShop/internal/app/dto/error"
	permissionDto "github.com/OnlineShop/internal/app/dto/permissions"
	apperrors "github.com/OnlineShop/internal/app/errors"
	"github.com/OnlineShop/internal/app/middlewares"
	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/pkg/logger"

	"github.com/OnlineShop/internal/app/utils"
)

type IPermissionService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
}

type PermissionService struct {
	permissionUsecase usecases.IPermissionUsecases
	log               logger.Ilogger
}

func NewPermissionService(u usecases.IPermissionUsecases, log logger.Ilogger) IPermissionService {
	log.Info("permission service is created")
	return &PermissionService{permissionUsecase: u, log: log}
}

func (s *PermissionService) FindAll(w http.ResponseWriter, r *http.Request) {

	//get users

	// permissions, err := s.PermissionRepository.FindAll()
	// if err != nil {
	// 	// TODO: prepare a model for error response to envelop the response
	// 	http.Error(w, "failed to get permissions", http.StatusBadRequest)
	// 	s.log.Error("responded to user with failed to get permissions",
	// 		"error from service:", err,
	// 		"statusCode:", http.StatusBadRequest,
	// 	)
	// 	return
	// }

	//convert to json
	// jsonResponse, errMarshal := json.Marshal(permissions)
	// if errMarshal != nil {
	// 	http.Error(w, "failed to parse json to serve", http.StatusInternalServerError)
	// 	s.log.Error("failed to parse permissions list to json and responded to user",
	// 		"jsonError:", errMarshal,
	// 		"statusCode:", http.StatusInternalServerError)
	// 	return
	// }

	s.log.Info("parsed permissions to json")

	//set headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//send response
	w.Write([]byte{})
	s.log.Info("responded to user with json file", "jsonResponse")

}

func (s *PermissionService) Create(w http.ResponseWriter, r *http.Request) {

	// TODO: check if already exists

	//TODO: add a middleware for authenticated routes and implement it in the router to use it

	// TODO: handle the keys type
	// TODO: extract the business logic to another package (use case )
	// TODO: make an error package to handle all errors in the app

	ctx := r.Context()
	rawBody := ctx.Value(middlewares.KEYCON)
	var requestBody permissionDto.PermissionCreateRequest
	utils.ConvertCtxValueToStruct(&rawBody, &requestBody, s.log)

	savedPermission, err := s.permissionUsecase.Create(&requestBody)
	if err != nil {
		responseError := dto.Error{
			Data: dto.ErrorData{
				Error:   err,
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			},
		}
		if errors.Is(err, apperrors.ErrValidation) {
			utils.SendErrorResponse(ctx, w, &responseError, s.log)
			return
		}
		if errors.Is(err, apperrors.ErrDatabase) {
			responseError.Data.Status = http.StatusInternalServerError
			responseError.Data.Message = "internal error"
			utils.SendErrorResponse(ctx, w, &responseError, s.log)
			return
		}

	}

	//convert to json and response
	responseValue := permissionDto.PermissionCreateResponse{
		ID:    savedPermission.ID,
		Title: savedPermission.Title,
	}

	utils.SendSuccessResponse(w, &responseValue, s.log)

}

func (service *PermissionService) FindById(w http.ResponseWriter, r *http.Request) {

}

func (service *PermissionService) Update(w http.ResponseWriter, r *http.Request) {

}

func (service *PermissionService) Delete(w http.ResponseWriter, r *http.Request) {

}
