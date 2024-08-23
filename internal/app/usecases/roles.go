package usecases

import (
	"errors"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	dto "github.com/OnlineShop/internal/app/dto/roles"
	"github.com/OnlineShop/internal/app/models"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/validation"
	"github.com/OnlineShop/internal/pkg/logger"
	"gorm.io/gorm"
)

type IRolesUsecases interface {
	Create(data *dto.RoleCreateRequest) (*models.Role, *apperrors.AppError)
}

type RolesUsecases struct {
	log        logger.Ilogger
	validator  validation.IRolesValidation
	repository repositories.IRolesRepository
}

func NewRolesUsecases(r repositories.IRolesRepository, v validation.IRolesValidation, l logger.Ilogger) IRolesUsecases {
	return &RolesUsecases{
		log:        l,
		validator:  v,
		repository: r,
	}
}

func (u *RolesUsecases) Create(data *dto.RoleCreateRequest) (*models.Role, *apperrors.AppError) {
	if err := u.validator.ValidateCreateRole(data); err != nil {
		return nil, err
	}

	if _, err := u.repository.IsRoleExists(data.Title); err != nil {
		// return nil, err

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

	}

	savedRole := models.Role{Title: data.Title}
	if _, err := u.repository.Create(&savedRole); err != nil {
		return nil, err
	}

	return &savedRole, nil
}
