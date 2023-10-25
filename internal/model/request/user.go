package request

type UpdateUserRequest struct {
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email"`
}

type UpdateUserPasswordRequest struct {
	CurrentPassword string `json:"current_password" validate:"required"`
	NewPassword     string `json:"new_password" validate:"required,gte=4,lte=16"`
}

type UpdateUserPictureRequest struct {
	Picture string `json:"picture" validate:"required,image"`
}
