package dto

type UserCreateRequest struct {
	Name        string `json:"name,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Email       string `json:"email" binding:"required"`
	DiscountID  uint   `json:"discountId,omitempty"`
	Password    string `json:"password" binding:"required"`
	Roles       []int  `json:"roles,omitempty"`
	Addresses   []int  `json:"adresses,omitempty"`
}
