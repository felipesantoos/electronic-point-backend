package request

import updatepassword "backend_template/src/core/domain/updatePassword"

type UpdatePassword struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

type updatePasswordBuilder struct{}

func UpdatePasswordBuilder() *updatePasswordBuilder {
	return &updatePasswordBuilder{}
}

func (*updatePasswordBuilder) FromBody(data map[string]interface{}) *UpdatePassword {
	return &UpdatePassword{
		CurrentPassword: data["current_password"].(string),
		NewPassword:     data["new_password"].(string),
	}
}

func (u *UpdatePassword) ToDomain() updatepassword.UpdatePassword {
	return updatepassword.New(u.CurrentPassword, u.NewPassword)
}
