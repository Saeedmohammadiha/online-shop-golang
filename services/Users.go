package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/OnlineShop/dto/user"
	"github.com/OnlineShop/models"
	"github.com/OnlineShop/repository"
	"github.com/OnlineShop/validation"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
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

	//get data from the request body and convert to json
	var requestUser dto.UserCreateRequest
	err := json.NewDecoder(r.Body).Decode(&requestUser)
	if err != nil {
		//fmt.Println("fail to decode the body")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//validate the inputs
	uv := validation.NewUserValidator()
	errors := uv.ValidateCreateUser(requestUser)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusBadRequest)
		return
	}

	//map the inputs to the user
	var newUser = models.User{
		Name:         requestUser.Name,
		LastName:     requestUser.LastName,
		PhoneNumber:  requestUser.PhoneNumber,
		Email:        requestUser.Email,
		PasswordHash: requestUser.Password, //needs to hash first
		DiscountID:   &requestUser.DiscountID,
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
	userId := mux.Vars(r)["id"]
	uId, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	//get the user from db
	var user models.User
	err = u.UserRepo.Db.First(&user, uId).Error
	if err != nil {
		http.Error(w, "can't get the user", http.StatusBadRequest)
		return
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
	userId := mux.Vars(r)["id"]
	uId, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	//does user exist
	if err = u.UserRepo.Db.Where("id = ?", uId).First(&models.User{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "there is no such user", http.StatusBadRequest)
			return
		}
	}

	//get data from the request body and convert to json
	var requestUser dto.UserUpdateRequest
	err = json.NewDecoder(r.Body).Decode(&requestUser)
	if err != nil {
		//fmt.Println("fail to decode the body")
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//validate the inputs
	uv := validation.NewUserValidator()
	errors := uv.ValidateUpdateUser(requestUser)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusBadRequest)
		return
	}

	//map the inputs to the user
	var updatedUser = models.User{
		ID:           uint(uId),
		Name:         requestUser.Name,
		LastName:     requestUser.LastName,
		PhoneNumber:  requestUser.PhoneNumber,
		Email:        requestUser.Email,
		PasswordHash: requestUser.Password, //needs to hash first
		DiscountID:   &requestUser.DiscountID,
	}

	// set the new sata
	dataSaved := u.UserRepo.Db.Save(&updatedUser)
	if dataSaved.Error != nil {
		fmt.Println("can't update the user", dataSaved.Error)
	}

	jsonResponse, _ := json.Marshal(updatedUser)

	//set response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//send the response
	w.Write(jsonResponse)
}

func (u *UserService) Delete(w http.ResponseWriter, r *http.Request) {

	//get the id from uri
	userId := mux.Vars(r)["id"]
	uId, err := strconv.Atoi(userId)
	if err != nil {
		http.Error(w, "invalid user id", http.StatusBadRequest)
		return
	}

	//does user exist
	if err = u.UserRepo.Db.Where("id = ?", uId).First(&models.User{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "there is no such user", http.StatusBadRequest)
			return
		}
	}
	//delete the user
	err = u.UserRepo.Db.Delete(&models.User{}, uId).Error
	if err != nil {
		http.Error(w, "cant delete the user", http.StatusBadRequest)
		return
	}

	//set response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//send the response
	w.Write([]byte{})
}
