package role

import (
	"eletronic_point/src/core/domain"
	"eletronic_point/src/core/domain/errors"
	"strings"

	"github.com/google/uuid"
)

const (
	ANONYMOUS_ROLE_CODE    = "anonymous"
	ADMIN_ROLE_CODE        = "admin"
	PROFESSIONAL_ROLE_CODE = "professional"
	TEACHER_ROLE_CODE      = "teacher"
	STUDENT_ROLE_CODE      = "student"
)

var possibleRoleCodes = []string{
	PROFESSIONAL_ROLE_CODE,
	ADMIN_ROLE_CODE,
	TEACHER_ROLE_CODE,
	STUDENT_ROLE_CODE,
}

type Role interface {
	domain.Model

	ID() *uuid.UUID
	Name() string
	Code() string
	IsProfessional() bool
	IsAdmin() bool
	IsTeacher() bool
	IsStudent() bool
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

func (r *role) ID() *uuid.UUID {
	return r.id
}

func (r *role) Name() string {
	return r.name
}

func (r *role) Code() string {
	return r.code
}

func (r *role) SetID(id *uuid.UUID) {
	r.id = id
}

func (r *role) SetName(name string) {
	r.name = name
}

func (r *role) SetCode(code string) {
	r.code = code
}

func (r *role) IsAdmin() bool {
	return strings.ToLower(r.code) == ADMIN_ROLE_CODE
}

func (r *role) IsProfessional() bool {
	return strings.ToLower(r.code) == PROFESSIONAL_ROLE_CODE
}

func (r *role) IsTeacher() bool {
	return strings.ToLower(r.code) == TEACHER_ROLE_CODE
}

func (r *role) IsStudent() bool {
	return strings.ToLower(r.code) == STUDENT_ROLE_CODE
}

func PossibleRoleCodes() []string {
	return possibleRoleCodes
}

func (r *role) IsValid() errors.Error {
	var exists bool = false
	for _, role := range possibleRoleCodes {
		if strings.ToLower(role) == strings.ToLower(r.code) {
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
