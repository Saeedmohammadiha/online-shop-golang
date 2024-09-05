package usecases

import (
	"context"
	"errors"
	"strconv"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	dto "github.com/OnlineShop/internal/app/dto/roles"
	"github.com/OnlineShop/internal/app/models"
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/validation"
	"github.com/OnlineShop/internal/pkg/logger"
	"gorm.io/gorm"
)

type IRolesUsecases interface {
	Create(ctx context.Context, data *dto.RoleCreateRequest) (*models.Role, *apperrors.AppError)
	Update(ctx context.Context, data *dto.RoleCreateRequest) (*models.Role, *apperrors.AppError)
	Delete(ctx context.Context, params map[string]string) *apperrors.AppError
	GetAll(ctx context.Context) (*[]models.Role, *apperrors.AppError)
	GetById(ctx context.Context, params map[string]string) (*models.Role, *apperrors.AppError)
}

type RolesUsecases struct {
	log                   logger.Ilogger
	validator             validation.IRolesValidation
	repository            repositories.IRolesRepository
	permissionsRepository repositories.IPermissionRepository
}

func NewRolesUsecases(r repositories.IRolesRepository, pr repositories.IPermissionRepository, v validation.IRolesValidation) IRolesUsecases {
	l := logger.Logger()
	l.Debug("new role usecase is created")
	return &RolesUsecases{
		log:                   l,
		validator:             v,
		repository:            r,
		permissionsRepository: pr,
	}
}

func (u *RolesUsecases) Delete(ctx context.Context, params map[string]string) *apperrors.AppError {

	
	

	roleId, error := strconv.Atoi(params["id"])
	if error != nil {
		return apperrors.NewBadRequestError("you need to pass the id", error)
	}

	err := u.repository.Delete(ctx, roleId)
	if err != nil {
		return err
	}

	return nil
}

func (u *RolesUsecases) GetAll(ctx context.Context) (*[]models.Role, *apperrors.AppError) {
	roles, err := u.repository.GetAll(ctx)

	if err != nil {
		return nil, err
	}
	return &roles, nil
}

func (u *RolesUsecases) GetById(ctx context.Context, params map[string]string) (*models.Role, *apperrors.AppError) {
	
	

	roleId, error := strconv.Atoi(params["id"])
	if error != nil {
		return nil, apperrors.NewBadRequestError("you need to pass the id", error)
	}

	role, err := u.repository.GetById(ctx, roleId)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (u *RolesUsecases) Update(ctx context.Context, data *dto.RoleCreateRequest) (*models.Role, *apperrors.AppError) {
	if err := u.validator.ValidateCreateRole(ctx, data); err != nil {
		return nil, err
	}

	permissions, err := u.permissionsRepository.GetByIds(ctx, data.PermissionIds)
	if err != nil {
		return nil, err
	}

	updatedRole := models.Role{Title: data.Title, Permissions: *permissions}
	if _, err := u.repository.Create(ctx, &updatedRole); err != nil {
		return nil, err
	}
	return &updatedRole, nil
}

func (u *RolesUsecases) Create(ctx context.Context, data *dto.RoleCreateRequest) (*models.Role, *apperrors.AppError) {
	if err := u.validator.ValidateCreateRole(ctx, data); err != nil {
		return nil, err
	}

	if _, err := u.repository.IsRoleExists(ctx, data.Title); err != nil {
		// return nil, err

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}

	}

	permissions, err := u.permissionsRepository.GetByIds(ctx, data.PermissionIds)
	if err != nil {
		return nil, err
	}

	newRole := models.Role{Title: data.Title, Permissions: *permissions}
	if _, err := u.repository.Create(ctx, &newRole); err != nil {
		return nil, err
	}

	return &newRole, nil
}
