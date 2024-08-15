package dto

import "github.com/OnlineShop/src/models"

type LoginResponse struct {
	User          models.User
	AccessToken   string
	RefreashToken string
}
