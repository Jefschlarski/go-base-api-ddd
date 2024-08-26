package dtos

// Password represents a password struct
type UpdatePassword struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}
