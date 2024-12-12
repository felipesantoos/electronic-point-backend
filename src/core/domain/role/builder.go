package role

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/utils/validator"
	"strings"

	"github.com/google/uuid"
)

type builder struct {
	fields        []string
	errorMessages []string
	role          *role
}

func NewBuilder() *builder {
	return &builder{
		fields:        []string{},
		errorMessages: []string{},
		role:          &role{},
	}
}

func (builder *builder) WithID(id uuid.UUID) *builder {
	if !validator.IsUUIDValid(id) {
		builder.fields = append(builder.fields, messages.RoleID)
		builder.errorMessages = append(builder.errorMessages, messages.RoleIDErrorMessage)
		return builder
	}
	builder.role.id = &id
	return builder
}

func (builder *builder) WithName(name string) *builder {
	name = strings.TrimSpace(name)
	if len(name) == 0 {
		builder.fields = append(builder.fields, messages.RoleName)
		builder.errorMessages = append(builder.errorMessages, messages.RoleNameErrorMessage)
		return builder
	}
	builder.role.name = name
	return builder
}

func (builder *builder) WithCode(code string) *builder {
	code = strings.TrimSpace(code)
	if !Exists(code) {
		builder.fields = append(builder.fields, messages.RoleCode)
		builder.errorMessages = append(builder.errorMessages, messages.RoleCodeErrorMessage)
		return builder
	}
	builder.role.code = code
	return builder
}

func (builder *builder) Build() (Role, errors.Error) {
	if len(builder.errorMessages) > 0 {
		return nil, errors.NewValidationWithMetadata(builder.errorMessages, map[string]interface{}{
			messages.Fields: builder.fields})
	}
	return builder.role, nil
}
