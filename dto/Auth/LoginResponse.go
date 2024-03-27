package dto

import "github.com/OnlineShop/models"

type LoginResponse struct {
	User          models.User
	AccessToken   string
	RefreashToken string
}
