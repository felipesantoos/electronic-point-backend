package authorization

import (
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/role"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthorization_New(t *testing.T) {
	auth := New()
	assert.NotNil(t, auth)
}

func TestAuthorization_NewFromAccount(t *testing.T) {
	os.Setenv("SERVER_SECRET", "test-secret")
	defer os.Unsetenv("SERVER_SECRET")

	id := uuid.New()
	pID := uuid.New()
	p, _ := person.New(&pID, "John Doe", "john@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
	r, _ := role.New(nil, "Admin", role.ADMIN_ROLE_CODE)
	acc, _ := account.New(&id, "john@example.com", "pass", r, p, nil, nil)

	auth, err := NewFromAccount(acc)
	assert.Nil(t, err)
	assert.NotEmpty(t, auth.Token())
	assert.NotNil(t, auth.ExpirationTime())
}

func TestAuthorization_NewFromToken(t *testing.T) {
	token := "some-token"
	exp := time.Now().Add(time.Hour)
	auth := NewFromToken(token, &exp)

	assert.Equal(t, token, auth.Token())
	assert.Equal(t, &exp, auth.ExpirationTime())
}
