package repository

import (
	"github.com/OnlineShop/models"
	"gorm.io/gorm"
)

type PermissionRepo interface {
	Create(permission *models.Permission) (*models.Permission, error)
	Update(permission *models.Permission) (*models.Permission, error)
	Delete(permissionID int) error
	FindById(permissionID int) (*models.Permission, error)
	FindAll() ([]models.Permission, error)
	FindByRoleAndResource(roleId int, resourceId int) (*models.Permission, error)
	//	FindByRoleIdes(roleIds []uint) (*[]models.Role, error)
}

type PermissionRepository struct {
	Db *gorm.DB
}

func NewPermissionRepository(db *gorm.DB) PermissionRepo {
	return &PermissionRepository{Db: db}
}

func (repo *PermissionRepository) Create(permission *models.Permission) (*models.Permission, error) {
	if err := repo.Db.Create(&permission).Error; err != nil {
		return nil, err
	}
	return permission, nil
}

func (repo *PermissionRepository) Update(permission *models.Permission) (*models.Permission, error) {
	if err := repo.Db.Model(&permission).Updates(permission).Error; err != nil {
		return nil, err
	}
	return permission, nil
}

func (repo *PermissionRepository) Delete(permissionID int) error {
	if err := repo.Db.Where("ID = ?", permissionID).Delete(permissionID).Error; err != nil {
		return err
	}
	return nil
}

func (repo *PermissionRepository) FindAll() ([]models.Permission, error) {
	var permissions []models.Permission
	if err := repo.Db.Find(&permissions).Error; err != nil {
		return nil, err
	}
	//	var results []models.Permission
	//results = append(results, permissions...)

	return permissions, nil
}

func (repo *PermissionRepository) FindById(permissionID int) (*models.Permission, error) {
	var permission models.Permission
	if err := repo.Db.First(&permission, permissionID).Error; err != nil {
		return nil, err
	}
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

func (p *PermissionRepository) FindByRoleAndResource(roleId int, resourceId int) (*models.Permission, error) {
	var permission models.Permission
	err := p.Db.Where("role_id = ? AND resource_id = ?", roleId, resourceId).First(&permission).Error
	if err != nil {
		return nil, err

	}
	return &permission, nil
}
