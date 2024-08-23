package dto

type CreateUserResponse struct {
	Id          uint    `json:"id"`
	Name        string `json:"name"`
	LastName    string `json:"lastName"`
	PhoneNumber string `json:"phoneNumber"`
	Email       string `json:"email"`
}
