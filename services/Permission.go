package services

import (
	"encoding/json"

	"net/http"

	dto "github.com/OnlineShop/dto/Permission"
	"github.com/OnlineShop/models"
	"github.com/OnlineShop/repository"
	"github.com/OnlineShop/validation"
)

type IPermissionService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
}

type PermissionService struct {
	PermissionRepo repository.PermissionRepo
}

func NewPermissionService(permissionRepo repository.PermissionRepo) IPermissionService {
	return &PermissionService{PermissionRepo: permissionRepo}
}

func (service *PermissionService) FindAll(w http.ResponseWriter, r *http.Request) {

	//get users

	permissions, err := service.PermissionRepo.FindAll()
	if err != nil {
		http.Error(w, "faild to ger users", http.StatusBadRequest)
		return
	}

	//convert to json
	jsonResponse, errMarshal := json.Marshal(permissions)
	if errMarshal != nil {
		http.Error(w, "faild to parse json to serve", http.StatusBadRequest)
		return
	}

	//set headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//send response
	w.Write(jsonResponse)

}

func (service *PermissionService) Create(w http.ResponseWriter, r *http.Request) {

	// TODO: need to sanitize the input
	// TODO: check if already exists
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
	if _, err := service.PermissionRepo.Create(&savedPermission); err != nil {
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
