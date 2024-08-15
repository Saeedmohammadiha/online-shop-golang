package repository

import (
	"github.com/OnlineShop/src/logger"
	"gorm.io/gorm"
)

type RepositoryFactory struct {
	PermissionRepository IPermissionRepository
}

func NewRepositoryFactory(db *gorm.DB, log logger.Ilogger) *RepositoryFactory {
	log.Info("repository factory is created")
	return &RepositoryFactory{
		PermissionRepository: NewPermissionRepository(db, log),
	}

}
