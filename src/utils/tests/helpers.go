package tests

import (
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/role"
	"github.com/google/uuid"
)

func NewMockAccount() account.Account {
	id := uuid.New()
	pID := uuid.New()
	p, _ := person.New(&pID, "John Doe", "john@example.com", "1990-01-01", "03611322055", "82999999999", "", "")
	r, _ := role.New(nil, "Admin", role.ADMIN_ROLE_CODE)
	acc, _ := account.New(&id, "john@example.com", "pass123", r, p, nil, nil)
	return acc
}
