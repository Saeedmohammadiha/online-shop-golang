package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/OnlineShop/models"
	"github.com/OnlineShop/repository"
	"github.com/OnlineShop/validation"
	"github.com/gorilla/mux"
)

type UserServiceType interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
}

type UserService struct {
	UserRepo repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: *userRepo}
}

func (u *UserService) FindAll(w http.ResponseWriter, r *http.Request) {

	//get users
	var users []models.User
	result := u.UserRepo.Db.Find(&users)
	if result.Error != nil {
		fmt.Println("failed to get users ", result.Error)
	}

	//convert to json
	jsonResponse, errMarshal := json.Marshal(users)
	if errMarshal != nil {
		fmt.Println("failed to parse json", errMarshal)
	}

	//set headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//send response
	w.Write(jsonResponse)
}

func (u *UserService) Create(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	//validate the inputs
	uv := validation.NewUserValidator(&newUser)
	errors := uv.ValidateCreateUser(r)
if errors != nil {
	println(errors)
	a, _ :=json.Marshal(errors)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//send response
	w.Write(a)

}
	//get data from the request body and convert to json
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		//fmt.Println("fail to decode the body")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//create the user in db
	user, er := u.UserRepo.Create(&newUser)
	if er != nil {
		fmt.Println("faield to create user", er)
	}

	//convert to json
	jsonResponse, errMarshal := json.Marshal(user)
	if errMarshal != nil {
		fmt.Println("fail to marshal user")
	}

	//set headers response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//send response
	w.Write(jsonResponse)
}

func (u *UserService) FindById(w http.ResponseWriter, r *http.Request) {

	//get the id from uri
	id := mux.Vars(r)["id"]

	//get the user from db
	var user models.User
	result := u.UserRepo.Db.Find(&user, id)

	if result.Error != nil {
		fmt.Println("cant get the user", result.Error)
	}

	//convert the user to json
	jsonResponse, errMarshal := json.Marshal(user)
	if errMarshal != nil {
		fmt.Println("failed to parse json")
	}

	//set response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//send the response
	w.Write(jsonResponse)
}

func (u *UserService) Updata(w http.ResponseWriter, r *http.Request) {

	//get the id from uri
	id := mux.Vars(r)["id"]

	//get the user from db
	var foundUser models.User
	result := u.UserRepo.Db.Find(&foundUser, id)
	if result.Error != nil {
		fmt.Println("cant get the user", result.Error)
	}

	//get data from the request body and convert to json
	var recivedUser models.User
	err := json.NewDecoder(r.Body).Decode(&recivedUser)
	if err != nil {
		fmt.Println("fail to decode the body")
	}

	// set the new sata
	dataSaved := u.UserRepo.Db.Save(&recivedUser)
	if dataSaved.Error != nil {
		fmt.Println("cant set the new data for the user", result.Error)
	}

	//convert the user to json
	jsonResponse, errMarshal := json.Marshal(&recivedUser)
	if errMarshal != nil {
		fmt.Println("failed to parse json")
	}

	//set response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//send the response
	w.Write(jsonResponse)
}

func (u *UserService) Delete(w http.ResponseWriter, r *http.Request) {

	//get the id from uri
	id := mux.Vars(r)["id"]

	//get the user from db
	var foundUser models.User
	result := u.UserRepo.Db.Find(&foundUser, id)
	if result.Error != nil {
		fmt.Println("cant get the user", result.Error)
	}

	// set the new sata
	dataSaved := u.UserRepo.Db.Delete(&foundUser)
	if dataSaved.Error != nil {
		fmt.Println("cant delete the user", result.Error)
	}

	//set response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//send the response
	w.Write([]byte{})
}
