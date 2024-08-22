package validation

import (

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	dto "github.com/OnlineShop/internal/app/dto/permissions"
	"github.com/OnlineShop/internal/pkg/logger"
	validation "github.com/go-ozzo/ozzo-validation"
)

type IPermissionValidation interface {
	ValidateCreatePermission(permission *dto.PermissionCreateRequest) *apperrors.AppError
}

type PermissionValidation struct {
	log logger.Ilogger
}

func NewPermissionValidation(l logger.Ilogger) IPermissionValidation {
	return &PermissionValidation{log: l}
}

func (v *PermissionValidation) ValidateCreatePermission(permission *dto.PermissionCreateRequest) *apperrors.AppError {
	err := validation.ValidateStruct(permission,
		validation.Field(&permission.Title, validation.Required.Error("you need to provide a title"), validation.Length(3, 50).Error("the title must be ableist 3 character")),
	)
	if err != nil {
		v.log.Error("validation error accuared", err)
		return apperrors.NewValidationError("the inputs are not valid", err)
	}
	v.log.Info("validation was successful")
	return nil
}
