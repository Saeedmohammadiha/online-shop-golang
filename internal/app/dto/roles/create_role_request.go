package dto

type RoleCreateRequest struct {
	Title         string  `json:"title"`
	PermissionIds []int   `json:"permissionIds"`
}
