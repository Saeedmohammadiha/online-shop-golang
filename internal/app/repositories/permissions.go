package repositories

import (
	"context"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	"github.com/OnlineShop/internal/app/models"
	"github.com/OnlineShop/internal/pkg/logger"
	"gorm.io/gorm"
)

type IPermissionRepository interface {
	Create(ctx context.Context, permission *models.Permission) (*models.Permission, *apperrors.AppError)
	Update(ctx context.Context, permission *models.Permission) (*models.Permission, *apperrors.AppError)
	Delete(ctx context.Context, permissionID int) *apperrors.AppError
	GetById(ctx context.Context, permissionID int) (*models.Permission, *apperrors.AppError)
	GetByIds(ctx context.Context, permissionID []int) (*[]models.Permission, *apperrors.AppError)
	GetAll(ctx context.Context, ) ([]models.Permission, *apperrors.AppError)
	IsPermissionExists(ctx context.Context, title string) (*models.Permission, *apperrors.AppError)
	// FindByRoleAndResource(roleId int, resourceId int) (*models.Permission, *apperrors.AppError)
	//	FindByRoleIdes(roleIds []uint) (*[]models.Role, *apperrors.AppError)
}

type PermissionRepository struct {
	Db  *gorm.DB
	log logger.Ilogger
}

func NewPermissionRepository(db *gorm.DB) IPermissionRepository {
	l := logger.Logger()
	l.Debug("permission Repository is created")
	return &PermissionRepository{Db: db, log: l}
}

func (r *PermissionRepository) GetByIds(ctx context.Context, permissionIds []int) (*[]models.Permission, *apperrors.AppError) {
	var permissions []models.Permission
	err := r.Db.Where("id IN ?", permissionIds).Find(&permissions).Error

	if err != nil {
		r.log.Error("failed to get permissions", "db error:", err.Error())
		return nil, apperrors.NewDatabaseError("can't get the permissions", err)
	}

	r.log.Debug("successfully found the permissions", permissions)
	return &permissions, nil
}

func (r *PermissionRepository) IsPermissionExists(ctx context.Context, title string) (*models.Permission, *apperrors.AppError) {
	var permission models.Permission
	err := r.Db.Where("title = ?", title).Take(&permission).Error

	// response with error if already exist
	if err != nil {
		r.log.Error("failed to get permission", "db error:", err.Error())
		return nil, apperrors.NewDatabaseError("failed to get permission", err)
	}
	r.log.Debug("successfully found the permission", permission)

	return &permission, nil
}

func (r *PermissionRepository) Create(ctx context.Context, permission *models.Permission) (*models.Permission, *apperrors.AppError) {
	//receive a pointer and pass the pointer to gorm create function
	if err := r.Db.Create(permission).Error; err != nil {
		r.log.Error("failed to create permission", "db error:", err.Error())
		return nil, apperrors.NewDatabaseError("failed to create permission", err)
	}
	r.log.Debug("new Permission created", permission)
	return permission, nil
}

func (r *PermissionRepository) Update(ctx context.Context, permission *models.Permission) (*models.Permission, *apperrors.AppError) {
	if err := r.Db.Model(permission).Updates(permission).Error; err != nil {
		r.log.Error("failed to update permission", "db error:", err.Error())
		return nil, apperrors.NewDatabaseError("failed to update permission", err)
	}
	r.log.Debug("Permission updated", permission)
	return permission, nil
}

func (r *PermissionRepository) Delete(ctx context.Context, permissionID int) *apperrors.AppError {
	if err := r.Db.Where("ID = ?", permissionID).Delete(permissionID).Error; err != nil {
		r.log.Error("failed to delete permission", "db error:", err.Error())
		return apperrors.NewDatabaseError("falied to delete permission", err)
	}
	r.log.Debug("Permission updated", permissionID)
	return nil
}

func (r *PermissionRepository) GetAll(ctx context.Context, ) ([]models.Permission, *apperrors.AppError) {
	var permissions []models.Permission
	if err := r.Db.Find(&permissions).Error; err != nil {
		r.log.Error("failed to get all permissions", "db error:", err.Error())
		return nil, apperrors.NewDatabaseError("failed to get permissions", err)
	}
	r.log.Debug("got the list of Permissions")

	return permissions, nil
}

func (r *PermissionRepository) GetById(ctx context.Context, permissionID int) (*models.Permission, *apperrors.AppError) {
	var permission models.Permission
	if err := r.Db.First(&permission, permissionID).Error; err != nil {
		r.log.Error("failed to get permission", "db error:", err.Error())
		return nil, apperrors.NewDatabaseError("failed to get permission", err)
	}
	r.log.Debug("got the Permission", permission)
	return &permission, nil

}

// func (repo *PermissionRepository) FindByRoleIdes(roleIds []uint) (*[]models.Role, error) {
// 	var roles []models.Role

// 	err := repo.Db.Where("id IN ?", roleIds).Find(&roles).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &roles, nil
// }

// func (p *PermissionRepository) FindByRoleAndResource(roleId int, resourceId int) (*models.Permission, error) {
// 	var permission models.Permission
// 	err := p.Db.Where("role_id = ? AND resource_id = ?", roleId, resourceId).First(&permission).Error
// 	if err != nil {
// 		return nil, err

// 	}
// 	return &permission, nil
// }
