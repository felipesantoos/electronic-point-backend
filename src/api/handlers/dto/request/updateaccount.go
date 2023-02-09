package request

import (
	"dit_backend/src/core/domain/account"
	"dit_backend/src/core/domain/errors"
	"dit_backend/src/core/domain/person"
	"fmt"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

type UpdateAccountProfile struct {
	ProfileID string `json:"profile_id"`
	Name      string `json:"name"`
	BirthDate string `json:"birth_date"`
	Phone     string `json:"phone"`

	profileID *uuid.UUID
}

type updateAccountProfileBuilder struct{}

func UpdateAccount() *updateAccountProfileBuilder {
	return &updateAccountProfileBuilder{}
}

func (*updateAccountProfileBuilder) FromBody(data map[string]interface{}) (*UpdateAccountProfile, errors.Error) {
	dto := &UpdateAccountProfile{}
	if id, parseErr := uuid.Parse(fmt.Sprint(data["profile_id"])); parseErr != nil {
		return nil, errors.New(parseErr)
	} else {
		dto.profileID = &id
	}
	if len(strings.Split(fmt.Sprint(data["name"]), " ")) < 2 {
		return nil, errors.NewFromString("your name must have a minimum of two words!")
	}
	dto.Name = fmt.Sprint(data["name"])
	birthDate := fmt.Sprint(data["birth_date"])
	if ok, _ := regexp.Match(birthDatePattern, []byte(birthDate)); !ok {
		return nil, errors.NewFromString("you must provide a date according to the following syntax: yyyy-MM-dd")
	}
	dto.BirthDate = birthDate
	dto.Phone = fmt.Sprint(data["phone"])
	return dto, nil
}

func (instance *UpdateAccountProfile) ToDomain() account.Account {
	return account.New(
		nil,
		"",
		"",
		nil,
		person.New(instance.profileID, instance.Name, "", "", "", instance.Phone, "", ""),
		nil,
	)
}
