package validation

import (
	"context"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	dto "github.com/OnlineShop/internal/app/dto/roles"
	"github.com/OnlineShop/internal/pkg/logger"
	validation "github.com/go-ozzo/ozzo-validation"
)

type IRolesValidation interface {
	ValidateCreateRole(ctx context.Context, role *dto.RoleCreateRequest) *apperrors.AppError
}

type RolesValidation struct {
	log logger.Ilogger
}

func NewRolesValidation(l logger.Ilogger) IRolesValidation {
	return &RolesValidation{log: l}
}

func (v *RolesValidation) ValidateCreateRole(ctx context.Context, role *dto.RoleCreateRequest) *apperrors.AppError {
	err := validation.ValidateStruct(role,
		validation.Field(&role.Title, validation.Required.Error("you need to provide a title"), validation.Length(3, 50).Error("the title must be ableist 3 character")),
		validation.Field(&role.PermissionIds, validation.Required.Error("you need to provide at least one permission")),
	)
	if err != nil {
		v.log.Error("validation error accuared", err)
		return apperrors.NewValidationError("the inputs are not valid", err)
	}
	v.log.Info("validation was successful")
	return nil
}
