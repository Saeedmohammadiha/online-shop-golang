package dto

import "github.com/OnlineShop/internal/app/models"


type LoginResponse struct {
	User          models.User
	AccessToken   string
	RefreshToken string
}
