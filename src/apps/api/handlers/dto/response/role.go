package response

import (
	"eletronic_point/src/core/domain/role"

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

func (*roleBuilder) BuildFromDomain(data role.Role) Role {
	return Role{
		ID:   data.ID(),
		Name: data.Name(),
		Code: data.Code(),
	}
}

func (r *roleBuilder) BuildFromDomainList(data []role.Role) []Role {
	result := make([]Role, 0)
	for _, item := range data {
		result = append(result, r.BuildFromDomain(item))
	}
	return result
}
