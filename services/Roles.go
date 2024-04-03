package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	dto "github.com/OnlineShop/dto/Role"
	"github.com/OnlineShop/models"
	"github.com/OnlineShop/repository"
)

type RoleService interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type RS struct {
	RoleRepo repository.RoleRepo
}

func NewRoleService(RoleRepo repository.RoleRepo) RoleService {
	return &RS{RoleRepo: RoleRepo}
}

func (Role *RS) Create(w http.ResponseWriter, r *http.Request) {
	//get data from the request body and convert to json
	var newRole dto.RoleCreateRequet
	err := json.NewDecoder(r.Body).Decode(&newRole)
	if err != nil {
		//fmt.Println("fail to decode the body")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// per := models.Permission{
	// 	Title: "test",
	// }
	roo := models.Role{
		Title:       newRole.Title,
		//Permissions: []models.Permission{per},
	}
	//create the user in db
	role, er := Role.RoleRepo.Create(&roo)
	if er != nil {
		fmt.Println("faield to create user", er)
	}
	fmt.Println(role)
	//convert to json
	jsonResponse, errMarshal := json.Marshal(role)
	if errMarshal != nil {
		fmt.Println("fail to marshal user")
	}

	//set headers response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//send response
	w.Write(jsonResponse)
}
