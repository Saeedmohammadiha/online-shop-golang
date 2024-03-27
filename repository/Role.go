package repository

import (
	"github.com/OnlineShop/models"
	"gorm.io/gorm"
)

type RoleRepo interface {
	Create(Role *models.Role) (*models.Role, error)
	Update(Role *models.Role) (*models.Role, error)
	Delete(RoleID int) error
	FindById(RoleID int) (*models.Role, error)
	FindAll() ([]models.Role, error)
}

type RoleRepository struct {
	Db *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepo {
	return &RoleRepository{Db: db}
}

func (repo *RoleRepository) Create(Role *models.Role) (*models.Role, error) {
	
	err := repo.Db.Create(&Role).Error
	if err != nil {
		return nil, err
	}
	return Role, nil
}

func (repo *RoleRepository) Update(Role *models.Role) (*models.Role, error) {
	repo.Db.Model(&Role).Updates(Role)
	return Role, nil
}

func (repo *RoleRepository) Delete(RoleID int) error {
	repo.Db.Where("ID = ?", RoleID).Delete(RoleID)
	return nil
}

func (repo *RoleRepository) FindAll() ([]models.Role, error) {
	var Roles []models.Role
	repo.Db.Find(&Roles)
	var results []models.Role
	results = append(results, Roles...)

	return results, nil
}

func (repo *RoleRepository) FindById(RoleID int) (*models.Role, error) {
	var result models.Role
//	repo.Db.Model(models.Role{ID: 10}).First(&result)
	return &result, nil

}
