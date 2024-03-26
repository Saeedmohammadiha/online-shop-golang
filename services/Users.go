package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/OnlineShop/dto/user"
	"github.com/OnlineShop/models"
	"github.com/OnlineShop/repository"
	utils "github.com/OnlineShop/utils/hashPassword"
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

	users, err := u.UserRepo.FindAll()
	if err != nil {
		http.Error(w, "faild to ger users", http.StatusBadRequest)
		return
	}

	//convert to json
	jsonResponse, errMarshal := json.Marshal(users)
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

func (u *UserService) Create(w http.ResponseWriter, r *http.Request) {

	//get data from the request body and convert to json
	var requestUser dto.CreateUserRequest
	err := json.NewDecoder(r.Body).Decode(&requestUser)
	if err != nil {
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

	var roles []models.Role
	if len(requestUser.RoleIDs) > 0 {
		// Fetch roles from the database using role IDs
		err := u.UserRepo.Db.Where("id IN ?", requestUser.RoleIDs).Find(&roles).Error
		if err != nil {
			http.Error(w, "there is no role with this id", http.StatusBadRequest)
			return
		}

		// Check if all role IDs are valid
		if len(roles) != len(requestUser.RoleIDs) {
			http.Error(w, "invalid role ids provided", http.StatusBadRequest)
			return
		}
	}

	//hash the password
	hashedPass, err := utils.NewPasswordHasher().HashPassword(requestUser.Password)
	if err != nil {
		http.Error(w, "encripting password faild", http.StatusBadRequest)
		return
	}

	//map the inputs to the user
	var newUser = models.User{
		Name:        requestUser.Name,
		LastName:    requestUser.LastName,
		PhoneNumber: requestUser.PhoneNumber,
		Email:       requestUser.Email,
		Password:    hashedPass,
	}
	if len(roles) > 0 {
		newUser.Roles = roles
	}

	//create the user in db
	user, er := u.UserRepo.Create(&newUser)
	if er != nil {
		http.Error(w, er.Error(), http.StatusBadRequest)
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

	user, errGetUser := u.UserRepo.FindById(uId)
	if errGetUser != nil {
		http.Error(w, "can't get the user", http.StatusBadRequest)
		return
	}

	//convert the user to json
	jsonResponse, errMarshal := json.Marshal(&user)
	if errMarshal != nil {
		http.Error(w, "faild to parse json to serve", http.StatusBadRequest)
		return
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
	_, errFInd := u.UserRepo.FindById(uId)
	if errFInd != nil {
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

	//hash the password
	var hashedPass = requestUser.Password
	if requestUser.Password != "" {
		hashedPass, err = utils.NewPasswordHasher().HashPassword(requestUser.Password)
		if err != nil {
			http.Error(w, "encripting password faild", http.StatusBadRequest)
			return
		}
	}

	//map the inputs to the user
	var updatedUser = models.User{
		Name:        requestUser.Name,
		LastName:    requestUser.LastName,
		PhoneNumber: requestUser.PhoneNumber,
		Email:       requestUser.Email,
		Password:    hashedPass,
	}

	// set the new sata
	_, errUpdate := u.UserRepo.Update(&updatedUser)
	if errUpdate != nil {
		http.Error(w, "can't update the uesr", http.StatusBadRequest)
		return
	}

	//set response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//send the response
	json.NewEncoder(w).Encode(updatedUser)
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
	user, err := u.UserRepo.FindById(uId)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			http.Error(w, "there is no such user", http.StatusBadRequest)
			return
		}
	}
	//delete the user
	err = u.UserRepo.Delete(int(user.ID))
	if err != nil {
		http.Error(w, "cant delete the user", http.StatusBadRequest)
		return
	}

	//set response header
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//send the response
	json.NewEncoder(w).Encode(models.User{})

}
