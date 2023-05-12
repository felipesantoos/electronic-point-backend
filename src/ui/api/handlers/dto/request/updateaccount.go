package request

import (
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/person"
)

type UpdateAccountProfile struct {
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
	Phone     string `json:"phone"`
}

type updateAccountProfileBuilder struct{}

func UpdateAccount() *updateAccountProfileBuilder {
	return &updateAccountProfileBuilder{}
}

func (instance *UpdateAccountProfile) ToDomain() (person.Person, errors.Error) {
	data, err := person.New(nil, instance.Name, instance.BirthDate, "", "", instance.Phone, "", "")
	messages := err.ValidationMessagesByMetadataFields([]string{"name", "birth_date"})
	if len(messages) > 0 {
		return nil, errors.NewValidation(messages)
	}
	return data, nil
}
