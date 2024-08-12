package services

import (
	"encoding/json"
	"fmt"
	//	"fmt"
	"net/http"
	////	"strconv"

	//	"github.com/OnlineShop/dto/User"
	//	"github.com/OnlineShop/models"
	"github.com/OnlineShop/models"
	"github.com/OnlineShop/repository"
	// "github.com/OnlineShop/utils"
	// "github.com/OnlineShop/validation"
	// "github.com/gorilla/mux"
	// "gorm.io/gorm"
)

type IPermissionService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Updata(w http.ResponseWriter, r *http.Request)
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

	recievedValue := &models.Permission{
		Title: "updateUser",
	}

	permission, err := service.PermissionRepo.Create(recievedValue)

fmt.Println("udpasd;fglkjhdslapfdsasdfasfdasfdas")
	if err != nil {
		http.Error(w, "errrororororor", http.StatusBadRequest)
		return
	}

	//convert to json
	jsonResponse, errMarshal := json.Marshal(permission)
	if errMarshal != nil {
		http.Error(w, "faild to parse json to serve", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(permission)

	w.Write(jsonResponse)

}

func (service *PermissionService) FindById(w http.ResponseWriter, r *http.Request) {

}

func (service *PermissionService) Updata(w http.ResponseWriter, r *http.Request) {

}

func (service *PermissionService) Delete(w http.ResponseWriter, r *http.Request) {

}
