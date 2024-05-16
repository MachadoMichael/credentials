package dto

type UpdatePasswordRequest struct {
	Email       string `json: "email"`
	OldPassword string `json: "oldPassword"`
	NewPassword string `json: "newPassword"`
}
