package role

import (
	"fmt"

	"github.com/google/uuid"
)

const (
	professionalRole = "professional"
)

var possibleRoleCodes = []string{professionalRole}

type Role interface {
	ID() *uuid.UUID
	Name() string
	Code() string

	SetID(*uuid.UUID)
	SetName(string)
	SetCode(string)
}

type role struct {
	id   *uuid.UUID
	name string
	code string
}

func New(id *uuid.UUID, name, code string) Role {
	return &role{id, name, code}
}

func NewEmpty() Role {
	return &role{}
}

func NewFromMap(data map[string]interface{}) (Role, error) {
	id, err := uuid.Parse(string(data["id"].([]uint8)))
	if err != nil {
		return nil, err
	}
	return &role{&id, fmt.Sprint(data["name"]), fmt.Sprint(data["code"])}, nil
}

func NewFromDerivedMap(data map[string]interface{}) (Role, error) {
	id, err := uuid.Parse(string(data["role_id"].([]uint8)))
	if err != nil {
		return nil, err
	}
	return &role{&id, fmt.Sprint(data["role_name"]), fmt.Sprint(data["role_code"])}, nil
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

func PossibleRoleCodes() []string {
	return possibleRoleCodes
}

func ProfessionalRoleCode() string {
	return professionalRole
}
