package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	dto "github.com/OnlineShop/dto/Auth"
	"github.com/OnlineShop/repository"
	"github.com/OnlineShop/utils"
	"github.com/OnlineShop/validation"
	"github.com/golang-jwt/jwt/v5"
)

type AuthService interface {
	Login(w http.ResponseWriter, r *http.Request)
}
type Auth struct {
	u repository.UserRepo
	p repository.PermissionRepo
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

	accessToken, err := utils.NewAuth().NewAccessToken(int(user.ID), user.RoleID)
	if err != nil {
		http.Error(w, "can't genrate accssess token", http.StatusBadRequest)
		return
	}
	refreshToken, err := utils.NewAuth().NewRefreshToken(int(user.ID), user.RoleID)
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

func (a *Auth) AuthMiddleware(next http.Handler, resourceId int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")

		claims := utils.JWTClaim{}
		if len(authHeader) != 2 {
			fmt.Println("Malformed token")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Malformed Token"))
		} else {
			jwtToken := authHeader[1]
			_, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
				return []byte(utils.GetEnv("TOKEN_SECRET")), nil
			})
			if err != nil {
				http.Error(w, "Unexpected signing method", http.StatusUnauthorized)
				return
			}
		}
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r)
	})

}

func (a *Auth) GaurdMiddleware(next http.Handler, resourceId int) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		//check if the user has the access
		permission, err := a.p.FindByRoleAndResource(claims.RoleID, resourceId)
		if err != nil {
			fmt.Println("Error querying permission:", err)

		}

		switch r.Method {
		case "GET":
			if !permission.Read {
				http.Error(w, "you don't have access", http.StatusUnauthorized)
				return
			}
		case "POST":
			if !permission.Create {
				http.Error(w, "you don't have access", http.StatusUnauthorized)
				return
			}
		case "PUT":
			if !permission.Update {
				http.Error(w, "you don't have access", http.StatusUnauthorized)
				return
			}
		case "DELETE":
			if !permission.Delete {
				http.Error(w, "you don't have access", http.StatusUnauthorized)
				return
			}

		}
		next.ServeHTTP(w, r)
	})
}
