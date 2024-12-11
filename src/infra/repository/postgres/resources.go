package postgres

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/infra/repository"
	"eletronic_point/src/infra/repository/postgres/query"
)

type resourcesPostgresPort struct{}

func NewResourcesRepository() secondary.ResourcesPort {
	return &resourcesPostgresPort{}
}

func (*resourcesPostgresPort) ListAccountRoles() ([]role.Role, errors.Error) {
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
