package services

import (
	"net/http"

	dto "github.com/OnlineShop/internal/app/dto/users"
	"github.com/OnlineShop/internal/app/usecases"
	"github.com/OnlineShop/internal/app/utils"
	"github.com/OnlineShop/internal/pkg/logger"
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

func NewUserService(u usecases.IUserUsecase, l logger.Ilogger) IUserService {
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
		utils.SendErrorResponse(ctx, err, w, s.log)
		return
	}

	utils.SendSuccessResponse(w, &permissions, s.log)

}

func (s *UserService) Create(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var requestBody dto.CreateUserRequest
	if err := utils.GetValueFromCtx(ctx, utils.REQUEST_BODY, &requestBody, s.log); err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
	}

	user, err := s.UserUsecase.Create(ctx, &requestBody)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
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

	utils.SendSuccessResponse(w, &responseValue, s.log)

}

func (s *UserService) GetById(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	user, err := s.UserUsecase.GetById(ctx)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
		return
	}

	utils.SendSuccessResponse(w, &user, s.log)

}

func (s *UserService) Update(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	var requestBody dto.CreateUserRequest
	if err := utils.GetValueFromCtx(ctx, utils.REQUEST_BODY, &requestBody, s.log); err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
	}

	updatedUser, err := s.UserUsecase.Update(ctx, &requestBody)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
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

	utils.SendSuccessResponse(w, &responseValue, s.log)

}

func (s *UserService) Delete(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()
	err := s.UserUsecase.Delete(ctx)
	if err != nil {
		utils.SendErrorResponse(ctx, err, w, s.log)
		return
	}

	utils.SendSuccessResponse(w, nil, s.log)

}
