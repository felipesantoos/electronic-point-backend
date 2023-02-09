package request

import (
	"dit_backend/src/core/domain/errors"
	"fmt"
	"net/mail"
)

type CreatePasswordReset struct {
	Email string `json:"email"`
}

type UpdatePasswordByPasswordReset struct {
	NewPassword string `json:"new_password"`
}

type createPasswordResetBuilder struct{}

func CreatePasswordResetBuilder() *createPasswordResetBuilder {
	return &createPasswordResetBuilder{}
}

func (*createPasswordResetBuilder) FromBody(data map[string]interface{}) (*CreatePasswordReset, errors.Error) {
	dto := &CreatePasswordReset{}
	email := fmt.Sprint(data["email"])
	if addr, _ := mail.ParseAddress(email); addr == nil {
		return nil, errors.NewFromString("you must provide a valid email!")
	} else {
		dto.Email = email
	}
	return dto, nil
}

type updatePasswordByPasswordResetBuilder struct{}

func UpdatePasswordByPasswordResetBuilder() *updatePasswordByPasswordResetBuilder {
	return &updatePasswordByPasswordResetBuilder{}
}

func (*updatePasswordByPasswordResetBuilder) FromBody(data map[string]interface{}) *UpdatePasswordByPasswordReset {
	return &UpdatePasswordByPasswordReset{
		NewPassword: fmt.Sprint(data["new_password"]),
	}
}
