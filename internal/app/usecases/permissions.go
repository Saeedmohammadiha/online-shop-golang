package usecases

import (
	"context"
	"errors"
	"net/url"
	"strconv"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	dto "github.com/OnlineShop/internal/app/dto/permissions"
	"github.com/OnlineShop/internal/app/models"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/utils"
	"github.com/OnlineShop/internal/app/validation"
	"github.com/OnlineShop/internal/pkg/logger"
	"gorm.io/gorm"
)

type IPermissionUsecases interface {
	Create(ctx context.Context, data *dto.PermissionCreateRequest) (*models.Permission, *apperrors.AppError)
	Update(ctx context.Context, data *dto.PermissionCreateRequest) (*models.Permission, *apperrors.AppError)
	Delete(ctx context.Context, params url.Values) *apperrors.AppError
	GetAll(ctx context.Context) (*[]models.Permission, *apperrors.AppError)
	GetById(ctx context.Context, params url.Values) (*models.Permission, *apperrors.AppError)
}

type PermissionUsecase struct {
	log        logger.Ilogger
	validator  validation.IPermissionValidation
	repository repositories.IPermissionRepository
}

func NewPermissionUsecase(r repositories.IPermissionRepository, v validation.IPermissionValidation) IPermissionUsecases {
	l := logger.Logger()
	l.Debug("new permission usecase is created")
	return &PermissionUsecase{
		log:        l,
		validator:  v,
		repository: r,
	}
}

func (u *PermissionUsecase) GetById(ctx context.Context, params url.Values) (*models.Permission, *apperrors.AppError) {

	permissionId := int(params.Get("id"))
	// if permissionId == nil {
	// 	return nil, apperrors.NewBadRequestError("you need to pass the id", err)
	// }

	permission, err := u.repository.GetById(ctx, permissionId)
	if err != nil {
		return nil, err
	}
	return permission, nil

}

func (u *PermissionUsecase) GetAll(ctx context.Context) (*[]models.Permission, *apperrors.AppError) {
	permissions, err := u.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}
	return &permissions, nil
}

func (u *PermissionUsecase) Delete(ctx context.Context) *apperrors.AppError {
	var params string
	err := utils.GetValueFromCtx(ctx, utils.REQUEST_PARAMS, &params)
	if err != nil {
		return apperrors.NewBadRequestError("there is no param", err)
	}

	permissionId, error := strconv.Atoi(params)
	if error != nil {
		return apperrors.NewBadRequestError("you need to pass the id", err)
	}

	err = u.repository.Delete(ctx, permissionId)
	if err != nil {
		return err
	}

	return nil
}

func (u *PermissionUsecase) Update(ctx context.Context, data *dto.PermissionCreateRequest) (*models.Permission, *apperrors.AppError) {
	// TODO:needs to add validation
	updatedPermission := models.Permission{Title: data.Title}
	if _, err := u.repository.Update(ctx, &updatedPermission); err != nil {
		return nil, err
	}
	return &updatedPermission, nil

}

func (u *PermissionUsecase) Create(ctx context.Context, data *dto.PermissionCreateRequest) (*models.Permission, *apperrors.AppError) {

	if err := u.validator.ValidateCreatePermission(ctx, data); err != nil {
		return nil, err
	}

	if _, err := u.repository.IsPermissionExists(ctx, data.Title); err != nil {
		// return nil, err

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

	}

	savedPermission := models.Permission{Title: data.Title}
	if _, err := u.repository.Create(ctx, &savedPermission); err != nil {
		return nil, err
	}

	return &savedPermission, nil
}
