package request

import (
	"eletronic_point/src/core/domain/role"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccount_ToDomain(t *testing.T) {
	dto := CreateAccount{
		Name:      "John Doe",
		BirthDate: "2000-01-01",
		Email:     "john.doe@email.com",
		CPF:       "11144477735",
		Phone:     "82999999999",
		RoleCode:  role.ADMIN_ROLE_CODE,
	}

	acc, err := dto.ToDomain()

	assert.Nil(t, err)
	assert.NotNil(t, acc)
	assert.Equal(t, dto.Email, acc.Email())
	assert.Equal(t, dto.Name, acc.Person().Name())
}
