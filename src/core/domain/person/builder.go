package person

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/utils/validator"
	"net/mail"
	"strings"

	"regexp"

	"github.com/google/uuid"
	"github.com/paemuri/brdoc"
)

type builder struct {
	fields        []string
	errorMessages []string
	person        *person
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		person:        &person{},
	}
}

func (builder *builder) WithID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		builder.fields = append(builder.fields, messages.PersonID)
		builder.errorMessages = append(builder.errorMessages, messages.PersonIDErrorMessage)
		return builder
	}

	builder.person.id = &id
	return builder
}

func (builder *builder) WithName(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 || len(strings.Split(name, " ")) == 1 {
		builder.fields = append(builder.fields, messages.PersonName)
		builder.errorMessages = append(builder.errorMessages, messages.PersonNameErrorMessage)
		return builder
	}

	builder.person.name = name
	return builder
}

func (builder *builder) WithEmail(email string) *builder {
	email = strings.TrimSpace(email)
	if addr, err := mail.ParseAddress(email); err != nil || addr == nil {
		builder.fields = append(builder.fields, messages.PersonEmail)
		builder.errorMessages = append(builder.errorMessages, messages.PersonEmailErrorMessage)
		return builder
	}

	builder.person.email = email
	return builder
}

func (builder *builder) WithBirthDate(birthDate string) *builder {
	birthDate = strings.TrimSpace(birthDate)
	if ok, _ := regexp.MatchString(birthDatePattern, birthDate); !ok {
		builder.fields = append(builder.fields, messages.PersonBirthDate)
		builder.errorMessages = append(builder.errorMessages, messages.PersonBirthDateErrorMessage)
		return builder
	}

	builder.person.birthDate = birthDate
	return builder
}

func (builder *builder) WithCPF(cpf string) *builder {
	cpf = strings.TrimSpace(cpf)
	if len(cpf) != 11 || !brdoc.IsCPF(cpf) {
		builder.fields = append(builder.fields, messages.PersonCPF)
		builder.errorMessages = append(builder.errorMessages, messages.PersonCPFErrorMessage)
		return builder
	}

	builder.person.cpf = cpf
	return builder
}

func (builder *builder) WithPhone(phone string) *builder {
	phone = strings.TrimSpace(phone)
	if len(phone) == 0 {
		builder.fields = append(builder.fields, messages.PersonPhone)
		builder.errorMessages = append(builder.errorMessages, messages.PersonPhoneErrorMessage)
		return builder
	}

	builder.person.phone = phone
	return builder
}

func (builder *builder) WithCreatedAt(createdAt string) *builder {
	createdAt = strings.TrimSpace(createdAt)
	if len(createdAt) == 0 {
		builder.fields = append(builder.fields, messages.PersonCreatedAt)
		builder.errorMessages = append(builder.errorMessages, messages.PersonCreatedAtErrorMessage)
		return builder
	}

	builder.person.createdAt = createdAt
	return builder
}

func (builder *builder) WithUpdatedAt(updatedAt string) *builder {
	updatedAt = strings.TrimSpace(updatedAt)
	if len(updatedAt) == 0 {
		builder.fields = append(builder.fields, messages.PersonUpdatedAt)
		builder.errorMessages = append(builder.errorMessages, messages.PersonUpdatedAtErrorMessage)
		return builder
	}

	builder.person.updatedAt = updatedAt
	return builder
}

func (builder *builder) Build() (Person, errors.Error) {
	if len(builder.errorMessages) > 0 {
		return nil, errors.NewValidationWithMetadata(builder.errorMessages, map[string]interface{}{
			messages.Fields: builder.fields,
		})
	}
	return builder.person, nil
}
