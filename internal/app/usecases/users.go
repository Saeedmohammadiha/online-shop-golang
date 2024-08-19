package usecases

import (
	"github.com/OnlineShop/internal/app/repositories"
	"github.com/OnlineShop/internal/app/validation"
	"github.com/OnlineShop/internal/pkg/logger"
)

type IUserUsecase interface {
	Create()
}

type UserUsecase struct {
	log        logger.Ilogger
	validator  validation.IUserValidation
	repository repositories.IUserRepository
}

func NewUserUsecase(r repositories.IUserRepository, v validation.IUserValidation, l logger.Ilogger) IUserUsecase {
	return &UserUsecase{
		log:        l,
		validator:  v,
		repository: r,
	}
}

func (u *UserUsecase) Create() {

}
