package usecases

import (
	"context"
	"errors"
	"strconv"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	dto "github.com/OnlineShop/internal/app/dto/roles"
	"github.com/OnlineShop/internal/app/models"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/utils"
	"github.com/OnlineShop/internal/app/validation"
	"github.com/OnlineShop/internal/pkg/logger"
	"gorm.io/gorm"
)

type IRolesUsecases interface {
	Create(ctx context.Context, data *dto.RoleCreateRequest) (*models.Role, *apperrors.AppError)
	Update(ctx context.Context, data *dto.RoleCreateRequest) (*models.Role, *apperrors.AppError)
	Delete(ctx context.Context) *apperrors.AppError
	GetAll(ctx context.Context) (*[]models.Role, *apperrors.AppError)
	GetById(ctx context.Context) (*models.Role, *apperrors.AppError)
}

type RolesUsecases struct {
	log                   logger.Ilogger
	validator             validation.IRolesValidation
	repository            repositories.IRolesRepository
	permissionsRepository repositories.IPermissionRepository
}

func NewRolesUsecases(r repositories.IRolesRepository, pr repositories.IPermissionRepository, v validation.IRolesValidation, l logger.Ilogger) IRolesUsecases {
	return &RolesUsecases{
		log:                   l,
		validator:             v,
		repository:            r,
		permissionsRepository: pr,
	}
}

func (u *RolesUsecases) Delete(ctx context.Context) *apperrors.AppError {

	var params string
	err := utils.GetValueFromCtx(ctx, utils.REQUEST_PARAMS, &params, u.log)
	if err != nil {
		return apperrors.NewBadRequestError("there is no param", err)
	}

	roleId, error := strconv.Atoi(params)
	if error != nil {
		return apperrors.NewBadRequestError("you need to pass the id", err)
	}

	err = u.repository.Delete(roleId)
	if err != nil {
		return err
	}

	return nil
}

func (u *RolesUsecases) GetAll(ctx context.Context) (*[]models.Role, *apperrors.AppError) {
	roles, err := u.repository.GetAll()

	if err != nil {
		return nil, err
	}
	return &roles, nil
}

func (u *RolesUsecases) GetById(ctx context.Context) (*models.Role, *apperrors.AppError) {
	var params string
	err := utils.GetValueFromCtx(ctx, utils.REQUEST_PARAMS, &params, u.log)
	if err != nil {
		return nil, apperrors.NewBadRequestError("there is no param", err)
	}

	roleId, error := strconv.Atoi(params)
	if error != nil {
		return nil, apperrors.NewBadRequestError("you need to pass the id", err)
	}

	role, err := u.repository.GetById(roleId)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (u *RolesUsecases) Update(ctx context.Context, data *dto.RoleCreateRequest) (*models.Role, *apperrors.AppError) {
	if err := u.validator.ValidateCreateRole(data); err != nil {
		return nil, err
	}

	permissions, err := u.permissionsRepository.GetByIds(data.PermissionIds)
	if err != nil {
		return nil, err
	}

	updatedRole := models.Role{Title: data.Title, Permissions: *permissions}
	if _, err := u.repository.Create(&updatedRole); err != nil {
		return nil, err
	}
	return &updatedRole, nil
}

func (u *RolesUsecases) Create(ctx context.Context, data *dto.RoleCreateRequest) (*models.Role, *apperrors.AppError) {
	if err := u.validator.ValidateCreateRole(data); err != nil {
		return nil, err
	}

	if _, err := u.repository.IsRoleExists(data.Title); err != nil {
		// return nil, err

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

	}

	permissions, err := u.permissionsRepository.GetByIds(data.PermissionIds)
	if err != nil {
		return nil, err
	}

	newRole := models.Role{Title: data.Title, Permissions: *permissions}
	if _, err := u.repository.Create(&newRole); err != nil {
		return nil, err
	}

	return &newRole, nil
}
