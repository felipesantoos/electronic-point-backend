package role

import (
	"backend_template/src/core/domain"
	"backend_template/src/core/domain/errors"
	"strings"

	"github.com/google/uuid"
)

const (
	ANONYMOUS_ROLE_CODE    = "anonymous"
	ADMIN_ROLE_CODE        = "admin"
	PROFESSIONAL_ROLE_CODE = "professional"
)

var possibleRoleCodes = []string{PROFESSIONAL_ROLE_CODE, ADMIN_ROLE_CODE}

type Role interface {
	domain.Model

	ID() *uuid.UUID
	Name() string
	Code() string
	IsProfessional() bool
	IsAdmin() bool
}

type role struct {
	id   *uuid.UUID
	name string
	code string
}

func New(id *uuid.UUID, name, code string) (data Role, err errors.Error) {
	data = &role{id, name, code}
	err = data.IsValid()
	return
}

func (instance *role) ID() *uuid.UUID {
	return instance.id
}

func (instance *role) Name() string {
	return instance.name
}

func (instance *role) Code() string {
	return instance.code
}

func (instance *role) SetID(id *uuid.UUID) {
	instance.id = id
}

func (instance *role) SetName(name string) {
	instance.name = name
}

func (instance *role) SetCode(code string) {
	instance.code = code
}

func (instance *role) IsAdmin() bool {
	return strings.ToLower(instance.code) == ADMIN_ROLE_CODE
}

func (instance *role) IsProfessional() bool {
	return strings.ToLower(instance.code) == PROFESSIONAL_ROLE_CODE
}

func PossibleRoleCodes() []string {
	return possibleRoleCodes
}

func (instance *role) IsValid() errors.Error {
	var exists bool = false
	for _, role := range possibleRoleCodes {
		if strings.ToLower(role) == strings.ToLower(instance.code) {
			exists = true
			break
		}
	}
	if !exists {
		return errors.NewValidationFromString("you must enter a valid role code. Valid Options: " + strings.Join(possibleRoleCodes, ", "))
	}
	return nil
}

func Exists(role string) bool {
	for _, item := range possibleRoleCodes {
		if role == item {
			return true
		}
	}
	return false
}
