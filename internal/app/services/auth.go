package services

import (
	"net/http"

	dto "github.com/OnlineShop/internal/app/dto/auth"
	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/app/utils"

	"github.com/OnlineShop/internal/pkg/logger"
)

type IAuthService interface {
	Login(w http.ResponseWriter, r *http.Request)
}
type AuthService struct {
	authUsecase usecases.IAuthUsecases
	log         logger.Ilogger
}

func NewAuthService(u usecases.IAuthUsecases, l logger.Ilogger) IAuthService {
	l.Info("new Auth service is created")
	return &AuthService{authUsecase: u, log: l}
}

func (a *AuthService) Login(w http.ResponseWriter, r *http.Request) {

	//get data from the context
	ctx := r.Context()
	var requestBody dto.LoginRequest
	if err := utils.GetValueFromCtx(ctx, utils.REQUEST_BODY, &requestBody, a.log); err != nil {
		utils.SendErrorResponse(ctx, err, w, a.log)
	}

	//get the response or err
	responseData, err := a.authUsecase.Login(&requestBody)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w, a.log)
		return
	}

	utils.SendSuccessResponse(w, responseData, a.log)
	return

}

// func (a *Auth) AuthMiddleware(next http.Handler, resourceId int) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")

// 		claims := utils.JWTClaim{}
// 		if len(authHeader) != 2 {
// 			fmt.Println("Malformed token")
// 			w.WriteHeader(http.StatusUnauthorized)
// 			w.Write([]byte("Malformed Token"))
// 		} else {
// 			jwtToken := authHeader[1]
// 			_, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
// 				return []byte(utils.GetEnv("TOKEN_SECRET")), nil
// 			})
// 			if err != nil {
// 				http.Error(w, "Unexpected signing method", http.StatusUnauthorized)
// 				return
// 			}
// 		}
// 		//ctx := context.WithValue(r.Context(), "claims", claims)
// 		next.ServeHTTP(w, r)
// 	})

// }

// func (a *Auth) GaurdMiddleware(next http.Handler, resourceId int) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

// 		//check if the user has the access
// 		permission, err := a.p.FindByRoleAndResource(1, resourceId)
// 		if err != nil {
// 			fmt.Println("Error querying permission:", err)

// 		}

// 		switch r.Method {
// 		case "GET":
// 			if !permission.Read {
// 				http.Error(w, "you don't have access", http.StatusUnauthorized)
// 				return
// 			}
// 		case "POST":
// 			if !permission.Create {
// 				http.Error(w, "you don't have access", http.StatusUnauthorized)
// 				return
// 			}
// 		case "PUT":
// 			if !permission.Update {
// 				http.Error(w, "you don't have access", http.StatusUnauthorized)
// 				return
// 			}
// 		case "DELETE":
// 			if !permission.Delete {
// 				http.Error(w, "you don't have access", http.StatusUnauthorized)
// 				return
// 			}

// 		}
// 		next.ServeHTTP(w, r)
// 	})
// }
