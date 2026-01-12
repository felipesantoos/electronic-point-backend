package authorization

import (
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/role"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewClaims(t *testing.T) {
	id := uuid.New()
	personID := uuid.New()

	// Simple mock-like entities
	p, _ := person.New(&personID, "John Doe", "john@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
	r, _ := role.New(nil, "Admin", role.ADMIN_ROLE_CODE)
	acc, _ := account.New(&id, "john@example.com", "", r, p, nil, nil)

	claims := newClaims(acc, "access", 3600)

	assert.NotNil(t, claims)
	assert.Equal(t, id.String(), claims.AccountID)
	assert.Equal(t, personID.String(), claims.ProfileID)
	assert.Equal(t, "access", claims.Type)
	assert.Equal(t, int64(3600), claims.Expiry)
}
