package repository

import "gorm.io/gorm"

type RepositoryFactory struct {
	PermissionRepository IPermissionRepository
}

func NewRepositoryFactory(db *gorm.DB) *RepositoryFactory {
	return &RepositoryFactory{
		PermissionRepository: NewPermissionRepo(db),
	}

}
