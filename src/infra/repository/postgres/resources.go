package postgres

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/core/interfaces/adapters"
	"eletronic_point/src/infra/repository"
	"eletronic_point/src/infra/repository/postgres/query"
)

type resourcesPostgresAdapter struct{}

func NewResourcesPostgresAdapter() adapters.ResourcesAdapter {
	return &resourcesPostgresAdapter{}
}

func (*resourcesPostgresAdapter) ListAccountRoles() ([]role.Role, errors.Error) {
	rows, err := repository.Queryx(query.AccountRole().Select().All())
	if err != nil {
		return nil, err
	}
	var roles = []role.Role{}
	for rows.Next() {
		var serializedRole = map[string]interface{}{}
		rows.MapScan(serializedRole)
		role, err := newRoleFromMapRows(serializedRole)
		if err != nil {
			return nil, errors.NewUnexpected()
		}
		roles = append(roles, role)
	}
	return roles, nil
}
