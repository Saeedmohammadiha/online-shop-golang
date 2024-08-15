package services

import (
	"encoding/json"

	"net/http"

	dto "github.com/OnlineShop/src/dto/Permission"
	"github.com/OnlineShop/src/logger"
	"github.com/OnlineShop/src/models"
	"github.com/OnlineShop/src/repository"
	"github.com/OnlineShop/src/validation"
)

type IPermissionService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
}

type PermissionService struct {
	PermissionRepository repository.IPermissionRepository
	log                  logger.Ilogger
}

func NewPermissionService(p repository.IPermissionRepository, log logger.Ilogger) IPermissionService {
	log.Info("permission service is created")
	return &PermissionService{PermissionRepository: p, log: log}
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

	// TODO: need to sanitize the input
	// TODO: check if already exists
	// TODO: decode the request in another package 
	// TODO: validation and sanatize should be in another package 
	// TODO: encode to json should be a helper function  
	// TODO: sending responses should be a helper function   
	// decode body to json
	var receivedPermission dto.PermissionCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&receivedPermission); err != nil {
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	// validate the permission
	if err := validation.NewPermissionValidation().ValidateCreatePermission(&receivedPermission); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// add the new entry to db
	savedPermission := models.Permission{Title: receivedPermission.Title}
	if _, err := s.PermissionRepository.Create(&savedPermission); err != nil {
		http.Error(w, "Oops! Something went wrong. please try again later", http.StatusInternalServerError)
		return
	}

	//convert to json and response
	responseValue := dto.PermissionCreateResponse{
		ID:    savedPermission.ID,
		Title: savedPermission.Title,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(responseValue); err != nil {
		http.Error(w, "failed to parse json", http.StatusInternalServerError)
	}

}

func (service *PermissionService) FindById(w http.ResponseWriter, r *http.Request) {

}

func (service *PermissionService) Update(w http.ResponseWriter, r *http.Request) {

}

func (service *PermissionService) Delete(w http.ResponseWriter, r *http.Request) {

}
