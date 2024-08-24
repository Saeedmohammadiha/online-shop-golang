package repositories

import (
	"github.com/OnlineShop/internal/pkg/logger"
	"gorm.io/gorm"
)

type RepositoryFactory struct {
	PermissionRepository IPermissionRepository
	UserRepository       IUserRepository
	RolesRepository      IRolesRepository
}

func NewRepositoryFactory(db *gorm.DB) *RepositoryFactory {
	logger.Logger().Info("repository factory is created")
	return &RepositoryFactory{
		PermissionRepository: NewPermissionRepository(db),
		UserRepository:       NewUserRepository(db),
		RolesRepository:      NewRolesRepository(db),
	}

}
