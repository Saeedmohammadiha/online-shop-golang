package usecases

import (
	dto "github.com/OnlineShop/internal/app/dto/permissions"
	"github.com/OnlineShop/internal/app/models"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/validation"
	"github.com/OnlineShop/internal/pkg/logger"
)

type IPermissionUsecases interface {
	Create(data *dto.PermissionCreateRequest) (*models.Permission, error)
}

type PermissionUsecase struct {
	log        logger.Ilogger
	validator  validation.IPermissionValidation
	repository repositories.IPermissionRepository
}

func NewPermissionUsecase(r repositories.IPermissionRepository, v validation.IPermissionValidation, l logger.Ilogger) IPermissionUsecases {
	return &PermissionUsecase{
		log:        l,
		validator:  v,
		repository: r,
	}
}

func (u *PermissionUsecase) Create(data *dto.PermissionCreateRequest) (*models.Permission, error) {
	if err := u.validator.ValidateCreatePermission(data); err != nil {
		return nil, err
	}


	savedPermission := models.Permission{Title: data.Title}
	if _, err := u.repository.Create(&savedPermission); err != nil {
		return nil, err
	}

	return &savedPermission, nil
}
