package dto

type RoleCreateResponse struct {
	ID          uint     `json:"id"`
	Title       string   `json:"title"`
	Permissions []string `json:"permissions"`
}
