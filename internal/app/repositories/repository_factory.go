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

func NewRepositoryFactory(db *gorm.DB, log logger.Ilogger) *RepositoryFactory {
	log.Info("repository factory is created")
	return &RepositoryFactory{
		PermissionRepository: NewPermissionRepository(db, log),
		UserRepository:       NewUserRepository(db, log),
		RolesRepository:      NewRolesRepository(db, log),
	}

}
