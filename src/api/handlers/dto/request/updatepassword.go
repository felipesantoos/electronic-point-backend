package request

import updatepassword "dit_backend/src/core/domain/updatePassword"

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

func (instance *UpdatePassword) ToDomain() updatepassword.UpdatePassword {
	return updatepassword.New(instance.CurrentPassword, instance.NewPassword)
}
