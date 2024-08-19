package repositories

import (
	"fmt"

	"github.com/OnlineShop/internal/app/models"
	"github.com/OnlineShop/internal/app/utils"
	"github.com/OnlineShop/internal/pkg/logger"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(user *models.User) (*models.User, error)
	Update(user *models.User) (*models.User, error)
	Delete(userID int) error
	GetById(userID int) (*models.User, error)
	GetAll() (*[]models.User, error)
	IsUserExists(email string) (*models.User, error)

	//	FindByRoleIdes(roleIds []uint) (*[]models.Role, error)
}

type UserRepository struct {
	Db  *gorm.DB
	log logger.Ilogger
}

func NewUserRepository(db *gorm.DB, l logger.Ilogger) IUserRepository {
	l.Info("new user Repository is created")
	return &UserRepository{Db: db, log: l}
}

func (r *UserRepository) IsUserExists(email string) (*models.User, error) {

	var user models.User
	err := r.Db.Where("email = ?", email).Take(&user).Error

	if err != nil {
		r.log.Error("failed to get user",
			"db error:", err.Error(),
			"data", email,
		)
		return nil, fmt.Errorf("%s%w", utils.ErrDatabaseTag, err)
	}
	r.log.Debug("successfully found the user",
		"data", user,
	)
	return &user, nil

}

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := r.Db.Create(user).Error; err != nil {
		r.log.Error("failed to create user",
			"db error:", err.Error(),
			"data", user,
		)
		return nil, fmt.Errorf("%s%w", utils.ErrDatabaseTag, err)
	}
	r.log.Debug("successfully created the user",
		"data", user,
	)
	return user, nil
}

func (r *UserRepository) Update(user *models.User) (*models.User, error) {
	if err := r.Db.Model(user).Updates(user).Error; err != nil {
		r.log.Error("failed to update user",
			"db error:", err.Error(),
			"data", user,
		)
		return nil, fmt.Errorf("%s%w", utils.ErrDatabaseTag, err)
	}
	r.log.Debug("successfully updated the user",
		"data", user,
	)
	return user, nil
}

func (r *UserRepository) Delete(userID int) error {
	if err := r.Db.Where("ID = ?", userID).Delete(userID).Error; err != nil {
		r.log.Error("failed to delete user",
			"db error:", err.Error(),
			"data", userID,
		)
		return fmt.Errorf("%s%w", utils.ErrDatabaseTag, err)
	}
	r.log.Debug("successfully deleted the user",
		"data", userID,
	)
	return nil
}

func (r *UserRepository) GetAll() (*[]models.User, error) {
	var users []models.User
	if err := r.Db.Find(&users).Error; err != nil {
		r.log.Error("failed to get users",
			"db error:", err.Error(),
			"data", nil,
		)
		return nil, fmt.Errorf("%s%w", utils.ErrDatabaseTag, err)
	}
	r.log.Debug("successfully got all the users",
		"data", users,
	)
	return &users, nil
}

func (r *UserRepository) GetById(userID int) (*models.User, error) {
	var user models.User
	if err := r.Db.First(&user, userID).Error; err != nil {
		r.log.Error("failed to find user",
			"db error:", err.Error(),
			"data", userID,
		)
		return nil, fmt.Errorf("%s%w", utils.ErrDatabaseTag, err)
	}
	r.log.Debug("successfully found the user",
		"data", user,
	)
	return &user, nil

}

// func (r *UserRepository) FindByRoleIdes(roleIds []uint) (*[]models.Role, error) {
// 	var roles []models.Role

// 	err := r.Db.Where("id IN ?", roleIds).Find(&roles).Error
// 	if err != nil {
// 		r.log.Error("failed to get user",
// 			"db error:", err.Error(),
// 			"data", email,
// 		)
// 		return nil, fmt.Errorf("%s%w", utils.ErrDatabaseTag, err)
// 	}

// 	return &roles, nil
// }
