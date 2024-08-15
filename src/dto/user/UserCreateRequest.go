package dto

type CreateUserRequest struct {
	Name        string `json:"name"`
	LastName    string `json:"lastName,omitempty"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
	RoleIDs     []uint `json:"roleIds,omitempty"`
	Password    string `json:"password"`
}
