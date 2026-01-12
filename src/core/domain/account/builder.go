package account

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/professional"
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/core/domain/student"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/utils/validator"
	"net/mail"
	"strings"

	"github.com/google/uuid"
)

type builder struct {
	fields        []string
	errorMessages []string
	account       *account
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		account:       &account{},
	}
}

func (builder *builder) WithID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		builder.fields = append(builder.fields, messages.AccountID)
		builder.errorMessages = append(builder.errorMessages, messages.AccountIDErrorMessage)
		return builder
	}
	builder.account.id = &id
	return builder
}

func (builder *builder) WithEmail(email string) *builder {
	email = strings.TrimSpace(email)
	if _, err := mail.ParseAddress(email); err != nil {
		builder.fields = append(builder.fields, messages.AccountEmail)
		builder.errorMessages = append(builder.errorMessages, messages.AccountEmailErrorMessage)
		return builder
	}
	builder.account.email = email
	return builder
}

func (builder *builder) WithPassword(password string) *builder {
	password = strings.TrimSpace(password)
	if len(password) < 6 {
		builder.fields = append(builder.fields, messages.AccountPassword)
		builder.errorMessages = append(builder.errorMessages, messages.AccountPasswordErrorMessage)
		return builder
	}
	builder.account.password = password
	return builder
}

func (builder *builder) WithRole(role role.Role) *builder {
	builder.account.role = role
	return builder
}

func (builder *builder) WithPerson(_person person.Person) *builder {
	builder.account.person = _person
	return builder
}

func (builder *builder) WithProfessional(_professional professional.Professional) *builder {
	if _professional != nil && _professional.IsValid() != nil {
		builder.fields = append(builder.fields, messages.Professional)
		builder.errorMessages = append(builder.errorMessages, messages.ProfessionalErrorMessage)
		return builder
	}
	builder.account.professional = _professional
	return builder
}

func (builder *builder) WithStudent(_student student.Student) *builder {
	if _student != nil && _student.IsValid() != nil {
		builder.fields = append(builder.fields, messages.AccountStudent)
		builder.errorMessages = append(builder.errorMessages, messages.AccountStudentErrorMessage)
		return builder
	}
	builder.account._student = _student
	return builder
}

func (builder *builder) Build() (Account, errors.Error) {
	if len(builder.errorMessages) > 0 {
		return nil, errors.NewValidationWithMetadata(builder.errorMessages, map[string]interface{}{
			messages.Fields: builder.fields,
		})
	}
	return builder.account, nil
}
