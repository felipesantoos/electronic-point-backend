package postgres

import (
	"dit_backend/src/core/domain/role"
	"dit_backend/src/core/interfaces/adapters"
	"dit_backend/src/infra"
	"dit_backend/src/infra/repository"
	"dit_backend/src/infra/repository/postgres/query"
)

type resourcesPostgresAdapter struct{}

func NewResourcesPostgresAdapter() adapters.ResourcesAdapter {
	return &resourcesPostgresAdapter{}
}

func (*resourcesPostgresAdapter) ListAccountRoles() ([]role.Role, infra.Error) {
	rows, err := repository.Queryx(query.AccountRole().Select().All())
	if err != nil {
		return nil, err
	}
	var roles = []role.Role{}
	for rows.Next() {
		var serializedRole = map[string]interface{}{}
		rows.MapScan(serializedRole)
		role, err := role.NewFromMap(serializedRole)
		if err != nil {
			return nil, infra.NewUnexpectedSourceErr()
		}
		roles = append(roles, role)
	}
	return roles, nil
}
