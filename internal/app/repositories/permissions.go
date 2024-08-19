package repositories

import (
	"fmt"

	"github.com/OnlineShop/internal/app/models"
	"github.com/OnlineShop/internal/app/utils"
	"github.com/OnlineShop/internal/pkg/logger"
	"gorm.io/gorm"
)

type IPermissionRepository interface {
	Create(permission *models.Permission) (*models.Permission, error)
	Update(permission *models.Permission) (*models.Permission, error)
	Delete(permissionID int) error
	GetById(permissionID int) (*models.Permission, error)
	GetAll() ([]models.Permission, error)
	IsPermissionExists(title string) (*models.Permission, error)
	// FindByRoleAndResource(roleId int, resourceId int) (*models.Permission, error)
	//	FindByRoleIdes(roleIds []uint) (*[]models.Role, error)
}

type PermissionRepository struct {
	Db  *gorm.DB
	log logger.Ilogger
}

func NewPermissionRepository(db *gorm.DB, log logger.Ilogger) IPermissionRepository {
	log.Info("permission Repository is created")
	return &PermissionRepository{Db: db, log: log}
}

func (r *PermissionRepository) IsPermissionExists(title string) (*models.Permission, error) {
	var permission models.Permission
	err := r.Db.Where("title = ?", title).Take(&permission).Error

	// response with error if already exist
	if err != nil {
		r.log.Error("failed to get permission", "db error:", err.Error())
		return nil, fmt.Errorf("%s%w", utils.ErrDatabaseTag, err)
	}
	r.log.Error("successfully found the permission", "db error:", err.Error())

	return &permission, nil
}

func (r *PermissionRepository) Create(permission *models.Permission) (*models.Permission, error) {
	//receive a pointer and pass the pointer to gorm create function
	if err := r.Db.Create(permission).Error; err != nil {
		r.log.Error("failed to create permission", "db error:", err.Error())
		return nil, fmt.Errorf("%s%w", utils.ErrDatabaseTag, err)
	}
	r.log.Info("new Permission created", permission)
	return permission, nil
}

func (r *PermissionRepository) Update(permission *models.Permission) (*models.Permission, error) {
	if err := r.Db.Model(permission).Updates(permission).Error; err != nil {
		r.log.Error("failed to update permission", "db error:", err.Error())
		return nil, fmt.Errorf("%s%w", utils.ErrDatabaseTag, err)
	}
	r.log.Info("Permission updated", permission)
	return permission, nil
}

func (r *PermissionRepository) Delete(permissionID int) error {
	if err := r.Db.Where("ID = ?", permissionID).Delete(permissionID).Error; err != nil {
		r.log.Error("failed to delete permission", "db error:", err.Error())
		return fmt.Errorf("%s%w", utils.ErrDatabaseTag, err)
	}
	r.log.Info("Permission updated", permissionID)
	return nil
}

func (r *PermissionRepository) GetAll() ([]models.Permission, error) {
	var permissions []models.Permission
	if err := r.Db.Find(&permissions).Error; err != nil {
		r.log.Error("failed to get all permissions", "db error:", err.Error())
		return nil, fmt.Errorf("%s%w", utils.ErrDatabaseTag, err)
	}
	r.log.Info("got the list of Permissions")

	return permissions, nil
}

func (r *PermissionRepository) GetById(permissionID int) (*models.Permission, error) {
	var permission models.Permission
	if err := r.Db.First(&permission, permissionID).Error; err != nil {
		r.log.Error("failed to get permission", "db error:", err.Error())
		return nil, fmt.Errorf("%s%w", utils.ErrDatabaseTag, err)
	}
	r.log.Info("got the Permission", permission)
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
