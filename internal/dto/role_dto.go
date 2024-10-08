package dto

type RoleRequestDTO struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

type RoleResponseDTO struct {
	DefaultDTO
	Name        string `json:"name"`
	Description string `json:"description" nested:"String"`
}

type UpdateRoleRequestDTO struct {
	Description *string `json:"description,omitempty"`
}

type AddRolePermissionRequestDTO struct {
	PermissionId string `json:"permissionId" binding:"required"`
}
