package response

import (
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/role"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAccountResponse_BuildFromDomain(t *testing.T) {
	id := uuid.New()
	pID := uuid.New()
	p, _ := person.New(&pID, "John Doe", "john@example.com", "1990-01-01", "03611322055", "82999999999", "", "")
	r, _ := role.New(nil, "Admin", role.ADMIN_ROLE_CODE)
	acc, _ := account.New(&id, "john@example.com", "", r, p, nil, nil)

	resp := AccountBuilder().BuildFromDomain(acc)
	assert.Equal(t, id.String(), resp.ID.String())
	assert.Equal(t, "john@example.com", resp.Email)
}
