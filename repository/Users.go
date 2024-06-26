package repository

import (
	"github.com/OnlineShop/models"
	"gorm.io/gorm"
)

type UserRepo interface {
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(userID int) error
	FindById(userID int) (*models.User, error)
	FindAll() ([]models.User, error)
	FindByEmail(userEmail string) (*models.User, error)
	FindByRoleIdes(roleIds []uint) (*[]models.Role, error)
}

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepo {
	return &UserRepository{Db: db}
}

func (repo *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := repo.Db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Update(user *models.User) (*models.User, error) {
	if err := repo.Db.Model(&user).Updates(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (repo *UserRepository) Delete(userID int) error {
	if err := repo.Db.Where("ID = ?", userID).Delete(userID).Error; err != nil {
		return err
	}
	return nil
}

func (repo *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	if err := repo.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	//	var results []models.User
	//results = append(results, users...)

	return users, nil
}

func (repo *UserRepository) FindById(userID int) (*models.User, error) {
	var user models.User
	if err := repo.Db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil

}
func (repo *UserRepository) FindByEmail(userEmail string) (*models.User, error) {
	var user models.User
	if err := repo.Db.Where("Email = ?", userEmail).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil

}

func (repo *UserRepository) FindByRoleIdes(roleIds []uint) (*[]models.Role, error) {
	var roles []models.Role

	err := repo.Db.Where("id IN ?", roleIds).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return &roles, nil
}
