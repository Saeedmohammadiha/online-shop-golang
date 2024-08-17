package services

import (
	"encoding/json"

	"net/http"

	errDto "github.com/OnlineShop/internal/app/dto/error"
	permissionDto "github.com/OnlineShop/internal/app/dto/permissions"
	"github.com/OnlineShop/internal/app/middlewares"
	"github.com/OnlineShop/internal/pkg/logger"

	"github.com/OnlineShop/internal/app/models"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/utils"
	"github.com/OnlineShop/internal/app/validation"
)

type IPermissionService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
}

type PermissionService struct {
	PermissionRepository repositories.IPermissionRepository
	log                  logger.Ilogger
	validator            validation.IPermissionValidation
}

func NewPermissionService(p repositories.IPermissionRepository, v validation.IPermissionValidation, log logger.Ilogger) IPermissionService {
	log.Info("permission service is created")
	return &PermissionService{PermissionRepository: p, log: log, validator: v}
}

func (s *PermissionService) FindAll(w http.ResponseWriter, r *http.Request) {

	//get users

	permissions, err := s.PermissionRepository.FindAll()
	if err != nil {
		// TODO: prepare a model for error response to envelop the response
		http.Error(w, "failed to get permissions", http.StatusBadRequest)
		s.log.Error("responded to user with failed to get permissions",
			"error from service:", err,
			"statusCode:", http.StatusBadRequest,
		)
		return
	}

	//convert to json
	jsonResponse, errMarshal := json.Marshal(permissions)
	if errMarshal != nil {
		http.Error(w, "failed to parse json to serve", http.StatusInternalServerError)
		s.log.Error("failed to parse permissions list to json and responded to user",
			"jsonError:", errMarshal,
			"statusCode:", http.StatusInternalServerError)
		return
	}

	s.log.Info("parsed permissions to json")

	//set headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//send response
	w.Write(jsonResponse)
	s.log.Info("responded to user with json file", jsonResponse)

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

	// validate the permission
	if err := s.validator.ValidateCreatePermission(&requestBody); err != nil {
		responseError := errDto.Error{
			Data: errDto.ErrorData{
				Error:   err,
				Message: err.Error(),
				Status:  http.StatusBadRequest,
			},
		}
		utils.SendErrorResponse(ctx, w, &responseError, s.log)
		return
	}

	// add the new entry to db
	savedPermission := models.Permission{Title: requestBody.Title}
	if _, err := s.PermissionRepository.Create(&savedPermission); err != nil {
		responseError := errDto.Error{
			Data: errDto.ErrorData{
				Error:   err,
				Message: "Oops! Something went wrong. please try again later",
				Status:  http.StatusInternalServerError,
			},
		}
		utils.SendErrorResponse(ctx, w, &responseError, s.log)
		return
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
