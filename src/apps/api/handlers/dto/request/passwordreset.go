package request

type CreatePasswordReset struct {
	Email string `json:"email" validate:"validemail"`
}

type UpdatePasswordByPasswordReset struct {
	NewPassword string `json:"new_password"`
}

type createPasswordResetBuilder struct{}

func CreatePasswordResetBuilder() *createPasswordResetBuilder {
	return &createPasswordResetBuilder{}
}
