package usecases

import (
	"context"
	"errors"
	"strconv"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	dto "github.com/OnlineShop/internal/app/dto/users"
	"github.com/OnlineShop/internal/app/models"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/utils"
	"github.com/OnlineShop/internal/app/validation"
	"github.com/OnlineShop/internal/pkg/logger"
	"gorm.io/gorm"
)

type IUserUsecase interface {
	Create(ctx context.Context, data *dto.CreateUserRequest) (*models.User, *apperrors.AppError)
	Update(ctx context.Context, data *dto.CreateUserRequest) (*models.User, *apperrors.AppError)
	Delete(ctx context.Context) *apperrors.AppError
	GetAll(ctx context.Context) (*[]models.User, *apperrors.AppError)
	GetById(ctx context.Context) (*models.User, *apperrors.AppError)
}

type UserUsecase struct {
	log        logger.Ilogger
	validator  validation.IUserValidation
	repository repositories.IUserRepository
}

func NewUserUsecase(r repositories.IUserRepository, v validation.IUserValidation) IUserUsecase {
	l := logger.Logger()
	l.Debug("new user usecase is created")
	return &UserUsecase{
		log:        l,
		validator:  v,
		repository: r,
	}
}

func (u *UserUsecase) Create(ctx context.Context, data *dto.CreateUserRequest) (*models.User, *apperrors.AppError) {
	if err := u.validator.ValidateCreateUser(ctx, data); err != nil {
		return nil, err
	}

	if _, err := u.repository.IsUserExists(ctx, data.Email); err != nil {

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

	}

	user := models.User{
		Name:        data.Name,
		LastName:    data.LastName,
		Email:       data.Email,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
	}
	if _, err := u.repository.Create(ctx, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (u *UserUsecase) Update(ctx context.Context, data *dto.CreateUserRequest) (*models.User, *apperrors.AppError) {
	// TODO:needs to add validation
	updatedUser := models.User{
		Name:        data.Name,
		LastName:    data.LastName,
		Email:       data.Email,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
	}
	if _, err := u.repository.Update(ctx, &updatedUser); err != nil {
		return nil, err
	}
	return &updatedUser, nil

}

func (u *UserUsecase) Delete(ctx context.Context) *apperrors.AppError {
	var params string
	err := utils.GetValueFromCtx(ctx, utils.REQUEST_PARAMS, &params)
	if err != nil {
		return apperrors.NewBadRequestError("there is no param", err)
	}

	userId, error := strconv.Atoi(params)
	if error != nil {
		return apperrors.NewBadRequestError("you need to pass the id", err)
	}

	err = u.repository.Delete(ctx, userId)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUsecase) GetById(ctx context.Context) (*models.User, *apperrors.AppError) {
	var params string
	err := utils.GetValueFromCtx(ctx, utils.REQUEST_PARAMS, &params)
	if err != nil {
		return nil, apperrors.NewBadRequestError("there is no param", err)
	}

	userId, error := strconv.Atoi(params)
	if error != nil {
		return nil, apperrors.NewBadRequestError("you need to pass the id", err)
	}

	user, err := u.repository.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}
	return user, nil

}

func (u *UserUsecase) GetAll(ctx context.Context) (*[]models.User, *apperrors.AppError) {
	users, err := u.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}
	return users, nil
}
