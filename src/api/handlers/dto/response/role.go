package response

import (
	"dit_backend/src/core/domain/role"

	"github.com/google/uuid"
)

type Role struct {
	ID   *uuid.UUID `json:"id"`
	Name string     `json:"name"`
	Code string     `json:"code"`
}

type roleBuilder struct{}

func AccountRoleBuilder() *roleBuilder {
	return &roleBuilder{}
}

func (*roleBuilder) FromDomain(data role.Role) Role {
	return Role{
		ID:   data.ID(),
		Name: data.Name(),
		Code: data.Code(),
	}
}

func (instance *Role) ToDomain() role.Role {
	return role.New(nil, instance.Name, instance.Code)
}
