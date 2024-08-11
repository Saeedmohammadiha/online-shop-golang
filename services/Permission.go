package services

import (
	"encoding/json"
//	"fmt"
	"net/http"
////	"strconv"

//	"github.com/OnlineShop/dto/User"
//	"github.com/OnlineShop/models"
	"github.com/OnlineShop/repository"
	// "github.com/OnlineShop/utils"
	// "github.com/OnlineShop/validation"
	// "github.com/gorilla/mux"
	// "gorm.io/gorm"
)

type PermissionServiceType interface {
	Create(w http.ResponseWriter, r *http.Request)
	Updata(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
}

type PermissionService struct {
	PermissionRepo repository.PermissionRepo
}

func NewPermissionService(permissionRepo repository.PermissionRepo) PermissionServiceType {
	return &PermissionService{PermissionRepo: permissionRepo}
}

func (p *PermissionService) FindAll(w http.ResponseWriter, r *http.Request) {

	//get users

	permissions, err := p.PermissionRepo.FindAll()
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

func (p *PermissionService) Create(w http.ResponseWriter, r *http.Request) {

}

func (p *PermissionService) FindById(w http.ResponseWriter, r *http.Request) {

	
}

func (p *PermissionService) Updata(w http.ResponseWriter, r *http.Request) {

}

func (p *PermissionService) Delete(w http.ResponseWriter, r *http.Request) {


}


