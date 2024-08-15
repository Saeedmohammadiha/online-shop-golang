package validation

import (
	dto "github.com/OnlineShop/src/dto/Permission"
	validation "github.com/go-ozzo/ozzo-validation"
)

type IPermissionValidation interface {
	ValidateCreatePermission(permission *dto.PermissionCreateRequest) error
}

type PermissionValidation struct{}

func NewPermissionValidation() IPermissionValidation {
	return &PermissionValidation{}
}

func (*PermissionValidation) ValidateCreatePermission(permission *dto.PermissionCreateRequest) error {
	return validation.ValidateStruct(permission,
		validation.Field(&permission.Title, validation.Required.Error("you need to provide a title"), validation.Length(3,50).Error("the title must be ableist 3 character")),
	)
}
