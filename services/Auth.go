package services

import (
	"encoding/json"
	"fmt"
	"net/http"

	dto "github.com/OnlineShop/dto/Auth"
	"github.com/OnlineShop/repository"
	"github.com/OnlineShop/utils"
	"github.com/OnlineShop/validation"
)

type AuthService interface {
	Login(w http.ResponseWriter, r *http.Request)
}
type Auth struct {
	u repository.UserRepo
}

func NewAuthService(r repository.UserRepo) AuthService {
	return &Auth{u: r}
}

func (A *Auth) Login(w http.ResponseWriter, r *http.Request) {

	//get data from the request body and convert to json
	var requestUser dto.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&requestUser)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//validate the inputs
	uv := validation.NewAuthValidator()
	errors := uv.ValidateLogin(requestUser)
	if errors != nil {
		http.Error(w, errors.Error(), http.StatusBadRequest)
		return
	}

	//is there a user
	user, errGetUser := A.u.FindByEmail(requestUser.Email)
	if errGetUser != nil {
		http.Error(w, "can't get the user", http.StatusBadRequest)
		return
	}

	//is the password correct
	ok := utils.NewPasswordHasher().CheckPasswordHash(requestUser.Password, user.Password)
	if !ok {
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}

	//return token and login the user

	accessToken, err := utils.NewAuth().NewAccessToken(int(user.ID))
	if err != nil {
		http.Error(w, "can't genrate accssess token", http.StatusBadRequest)
		return
	}
	refreshToken, err := utils.NewAuth().NewRefreshToken(int(user.ID))
	if err != nil {
		http.Error(w, "can't genrate refresh token", http.StatusBadRequest)
		return
	}

	response := dto.LoginResponse{
		User:          *user,
		AccessToken:   accessToken,
		RefreashToken: refreshToken,
	}

	
	//convert to json
	jsonResponse, errMarshal := json.Marshal(response)
	if errMarshal != nil {
		fmt.Println("fail to marshal user")
	}

	//set headers response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	//send response
	w.Write(jsonResponse)

}
