package postgres

import (
	"backend_template/src/core/domain/errors"
	"backend_template/src/core/domain/role"
	"backend_template/src/core/interfaces/adapters"
	"backend_template/src/infra/repository"
	"backend_template/src/infra/repository/postgres/query"
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
