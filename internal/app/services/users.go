package services

import (
	"encoding/json"
	"net/http"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	dto "github.com/OnlineShop/internal/app/dto/users"
	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/app/utils"
	"github.com/OnlineShop/internal/pkg/logger"
	"github.com/gorilla/mux"
)

type IUserService interface {
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetById(w http.ResponseWriter, r *http.Request)
}

type UserService struct {
	UserUsecase usecases.IUserUsecase
	log         logger.Ilogger
}

func NewUserService(u usecases.IUserUsecase) IUserService {
	l := logger.Logger()
	l.Info("permission service is created")
	return &UserService{
		log:         l,
		UserUsecase: u,
	}
}

func (s *UserService) GetAll(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	permissions, err := s.UserUsecase.GetAll(ctx)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w)
		return
	}

	utils.SendSuccessResponse(w, &permissions)

}

func (s *UserService) Create(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var requestBody dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {

		utils.SendErrorResponse(
			ctx,
			apperrors.NewBadRequestError("invalid json", err),
			w,
		)
	}

	user, err := s.UserUsecase.Create(ctx, &requestBody)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w)
		return
	}

	//convert to json and response
	responseValue := dto.CreateUserResponse{
		Id:          user.ID,
		Name:        user.Name,
		LastName:    user.LastName,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	utils.SendSuccessResponse(w, &responseValue)

}

func (s *UserService) GetById(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	params := mux.Vars(r)
	user, err := s.UserUsecase.GetById(ctx,params)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w)
		return
	}

	utils.SendSuccessResponse(w, &user)

}

func (s *UserService) Update(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var requestBody dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {

		utils.SendErrorResponse(
			ctx,
			apperrors.NewBadRequestError("invalid json", err),
			w,
		)
	}

	updatedUser, err := s.UserUsecase.Update(ctx, &requestBody)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w)
		return
	}

	//convert to json and response
	responseValue := dto.CreateUserResponse{
		Id:          updatedUser.ID,
		Name:        updatedUser.Name,
		LastName:    updatedUser.LastName,
		Email:       updatedUser.Email,
		PhoneNumber: updatedUser.PhoneNumber,
	}

	utils.SendSuccessResponse(w, &responseValue)

}

func (s *UserService) Delete(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	params := mux.Vars(r)
	err := s.UserUsecase.Delete(ctx, params)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w)
		return
	}

	utils.SendSuccessResponse(w, nil)

}
