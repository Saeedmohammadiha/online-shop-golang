package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/OnlineShop/models"
	"github.com/OnlineShop/repository"
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

func (u *UserService) Create(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		fmt.Println("fail to decode the body")
	}

	a, er := u.UserRepo.Create(&newUser)
	if er != nil {
		fmt.Println("hi")
	}
	jsonResponse, errMarshal := json.Marshal(a)
	if errMarshal != nil {
		fmt.Println("fail to marshal user")
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(jsonResponse)
}
