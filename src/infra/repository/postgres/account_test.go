package postgres

import (
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/core/services/filters"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAccountRepository_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	SetupTestDB(t)
	CleanDB(t)

	repo := NewAccountRepository()

	t.Run("Create and Find", func(t *testing.T) {
		CleanDB(t)
		id := uuid.New()
		pID := uuid.New()
		p, _ := person.New(&pID, "John Doe", "john@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
		r, _ := role.New(nil, "Admin", role.ADMIN_ROLE_CODE)
		acc, _ := account.New(&id, "john@example.com", "pass123", r, p, nil, nil)

		createdID, err := repo.Create(acc)
		assert.Nil(t, err)
		assert.NotNil(t, createdID)

		found, err := repo.FindByID(createdID)
		assert.Nil(t, err)
		assert.Equal(t, acc.Email(), found.Email())
	})

	t.Run("List", func(t *testing.T) {
		CleanDB(t)
		// Create an account first
		id := uuid.New()
		pID := uuid.New()
		p, _ := person.New(&pID, "John Doe", "john@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
		r, _ := role.New(nil, "Admin", role.ADMIN_ROLE_CODE)
		acc, _ := account.New(&id, "john@example.com", "pass123", r, p, nil, nil)
		_, _ = repo.Create(acc)

		accounts, err := repo.List(filters.AccountFilters{})
		assert.Nil(t, err)
		assert.NotEmpty(t, accounts)
	})

	t.Run("Update", func(t *testing.T) {
		CleanDB(t)
		// Create account first
		id := uuid.New()
		pID := uuid.New()
		p, _ := person.New(&pID, "Jane Doe", "jane@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
		r, _ := role.New(nil, "Admin", role.ADMIN_ROLE_CODE)
		acc, _ := account.New(&id, "jane@example.com", "pass123", r, p, nil, nil)

		createdID, err := repo.Create(acc)
		assert.Nil(t, err)
		assert.NotNil(t, createdID)

		// Update account
		found, err := repo.FindByID(createdID)
		assert.Nil(t, err)

		found.SetEmail("jane.updated@example.com")
		updatedPerson, _ := person.NewBuilder().
			WithID(*found.Person().ID()).
			WithName("Jane Updated").
			WithEmail("jane.updated@example.com").
			WithBirthDate("1990-01-01").
			WithCPF("11144477735").
			WithPhone("82988888888").
			Build()
		found.SetPerson(updatedPerson)

		err = repo.Update(found)
		assert.Nil(t, err)

		// Verify update
		updated, err := repo.FindByID(createdID)
		assert.Nil(t, err)
		assert.Equal(t, "jane.updated@example.com", updated.Email())
		assert.Equal(t, "Jane Updated", updated.Person().Name())
	})

	t.Run("Delete", func(t *testing.T) {
		CleanDB(t)
		// Create account first
		id := uuid.New()
		pID := uuid.New()
		p, _ := person.New(&pID, "Delete Test", "delete@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
		r, _ := role.New(nil, "Admin", role.ADMIN_ROLE_CODE)
		acc, _ := account.New(&id, "delete@example.com", "pass123", r, p, nil, nil)

		createdID, err := repo.Create(acc)
		assert.Nil(t, err)
		assert.NotNil(t, createdID)

		// Delete account
		err = repo.Delete(*createdID)
		assert.Nil(t, err)

		// Verify deletion
		_, err = repo.FindByID(createdID)
		assert.NotNil(t, err)
	})

	t.Run("UpdateAccountProfile", func(t *testing.T) {
		CleanDB(t)
		// Create account first
		id := uuid.New()
		pID := uuid.New()
		p, _ := person.New(&pID, "Profile Test", "profile@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
		r, _ := role.New(nil, "Admin", role.ADMIN_ROLE_CODE)
		acc, _ := account.New(&id, "profile@example.com", "pass123", r, p, nil, nil)

		createdID, err := repo.Create(acc)
		assert.Nil(t, err)
		assert.NotNil(t, createdID)

		// Get account to update profile
		found, err := repo.FindByID(createdID)
		assert.Nil(t, err)

		// Update profile
		updatedPerson, _ := person.NewBuilder().
			WithID(*found.Person().ID()).
			WithName("Profile Updated").
			WithEmail(found.Person().Email()).
			WithBirthDate("1991-02-02").
			WithCPF(found.Person().CPF()).
			WithPhone("82977777777").
			Build()

		err = repo.UpdateAccountProfile(updatedPerson)
		assert.Nil(t, err)

		// Verify update
		updated, err := repo.FindByID(createdID)
		assert.Nil(t, err)
		assert.Equal(t, "Profile Updated", updated.Person().Name())
		assert.Equal(t, "82977777777", updated.Person().Phone())
	})

	// Note: UpdateAccountPassword test is skipped because Create generates a random password
	// that is sent via email. A full integration test would require mocking email or
	// storing the generated password. This is left as a TODO for future implementation.
}
