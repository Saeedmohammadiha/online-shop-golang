package dto

type UserUpdateRequest struct {
	Name        string `json:"name,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	PhoneNumber string `json:"phoneNumber,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	Roles       []int  `json:"roles,omitempty"`
	Addresses   []int  `json:"adresses,omitempty"`
}
