package dto

type UserResponseDTO struct {
	DefaultDTO
	Name           string `json:"name"`
	Email          string `json:"email"`
	Status         string `json:"status"`
	Image          string `json:"image" nested:"String"`
	SocialProvider string `json:"socialProvider"`
}
