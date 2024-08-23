package repositories

import (
	"context"

	apperrors "github.com/OnlineShop/internal/app/app_errors"
	"github.com/OnlineShop/internal/app/models"
	"github.com/OnlineShop/internal/pkg/logger"
	"gorm.io/gorm"
)

type IRolesRepository interface {
	Create(ctx context.Context, Role *models.Role) (*models.Role, *apperrors.AppError)
	IsRoleExists(ctx context.Context, title string) (*models.Role, *apperrors.AppError)
	Update(ctx context.Context, Role *models.Role) (*models.Role, *apperrors.AppError)
	Delete(ctx context.Context, RoleID int) *apperrors.AppError
	GetById(ctx context.Context, RoleID int) (*models.Role, *apperrors.AppError)
	GetAll(ctx context.Context, ) ([]models.Role, *apperrors.AppError)
}

type RolesRepository struct {
	Db  *gorm.DB
	log logger.Ilogger
}

func NewRolesRepository(db *gorm.DB, l logger.Ilogger) IRolesRepository {
	l.Debug("roles Repository is created")
	return &RolesRepository{Db: db, log: l}
}

func (r *RolesRepository) IsRoleExists(ctx context.Context, title string) (*models.Role, *apperrors.AppError) {
	var role models.Role
	err := r.Db.Where("title = ?", title).Take(&role).Error

	// response with error if already exist
	if err != nil {
		r.log.Error("failed to get role", "db error:", err.Error())
		return nil, apperrors.NewDatabaseError("failed to get role", err)
	}
	r.log.Debug("successfully found the role", role)

	return &role, nil
}

func (r *RolesRepository) Create(ctx context.Context, role *models.Role) (*models.Role, *apperrors.AppError) {

	err := r.Db.Create(role).Error
	if err != nil {
		r.log.Error("failed to create role", "db error:", err.Error())
		return nil, apperrors.NewDatabaseError("failed to create the role", err)
	}
	r.log.Debug("new role created", role)
	return role, nil
}

func (r *RolesRepository) Update(ctx context.Context, role *models.Role) (*models.Role, *apperrors.AppError) {
	err := r.Db.Model(role).Updates(role).Error
	if err != nil {
		r.log.Error("failed to update role", "db error:", err.Error())
		return nil, apperrors.NewDatabaseError("failed to update the role", err)
	}
	r.log.Debug("new role updated", role)
	return role, nil
}

func (r *RolesRepository) Delete(ctx context.Context, roleId int) *apperrors.AppError {
	err := r.Db.Where("ID = ?", roleId).Delete(roleId).Error
	if err != nil {
		r.log.Error("failed to delete role", "db error:", err.Error())
		return apperrors.NewDatabaseError("failed to update the role", err)
	}
	r.log.Debug("role has been deleted", roleId)
	return nil
}

func (r *RolesRepository) GetAll(ctx context.Context, ) ([]models.Role, *apperrors.AppError) {
	var roles []models.Role
	if err := r.Db.Find(&roles).Error; err != nil {
		r.log.Error("failed to get all roles", "db error:", err.Error())
		return nil, apperrors.NewDatabaseError("failed to get roles", err)
	}
	r.log.Debug("got the list of roles")

	return roles, nil
}

func (r *RolesRepository) GetById(ctx context.Context, roleID int) (*models.Role, *apperrors.AppError) {
	var role models.Role
	if err := r.Db.First(&role, roleID).Error; err != nil {
		r.log.Error("failed to get role", "db error:", err.Error())
		return nil, apperrors.NewDatabaseError("failed to get role", err)
	}
	r.log.Debug("got the role", role)
	return &role, nil

}
