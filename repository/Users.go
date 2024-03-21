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
	FindAll() ([]*models.User, error)
}

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{Db: db}
}

func (repo *UserRepository) Create(user *models.User) (*models.User, error) {
	repo.Db.Create(&user)
	return user, nil
}

func (repo *UserRepository) Update(user *models.User) (*models.User, error) {
	repo.Db.Model(&user).Updates(user)
	return user, nil
}

func (repo *UserRepository) Delete(userID int) error {
	repo.Db.Where("ID = ?", userID).Delete(userID)
	return nil
}

func (repo *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	repo.Db.Find(&users)
	var results []models.User
	results = append(results, users...)

	return results, nil
}

func (repo *UserRepository) FindById(userID int) (*models.User, error) {
	var result models.User
	repo.Db.Model(models.User{ID: 10}).First(&result)
	return &result, nil

}
