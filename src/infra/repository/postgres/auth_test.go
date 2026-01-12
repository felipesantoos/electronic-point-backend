package postgres

import (
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/credentials"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/role"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestAuthRepository_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	SetupTestDB(t)
	CleanDB(t)

	repo := NewAuthPostgresRepository()
	accountRepo := NewAccountRepository()

	t.Run("Login", func(t *testing.T) {
		CleanDB(t)
		// Create account first
		id := uuid.New()
		pID := uuid.New()
		p, _ := person.New(&pID, "Login Test", "logintest@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
		r, _ := role.New(nil, "Admin", role.ADMIN_ROLE_CODE)
		acc, _ := account.New(&id, "logintest@example.com", "testpass123", r, p, nil, nil)

		createdID, err := accountRepo.Create(acc)
		assert.Nil(t, err)
		assert.NotNil(t, createdID)

		// Set a known password using ResetAccountPassword
		knownPassword := "testpass123"
		err = repo.ResetAccountPassword(createdID, knownPassword)
		assert.Nil(t, err)

		// Test login with correct credentials
		creds := credentials.New("logintest@example.com", knownPassword)
		loggedInAccount, err := repo.Login(creds)
		assert.Nil(t, err)
		assert.NotNil(t, loggedInAccount)
		assert.Equal(t, "logintest@example.com", loggedInAccount.Email())

		// Test login with incorrect password
		wrongCreds := credentials.New("logintest@example.com", "wrongpassword")
		_, err = repo.Login(wrongCreds)
		assert.NotNil(t, err)

		// Test login with non-existent email
		nonExistentCreds := credentials.New("nonexistent@example.com", "password")
		_, err = repo.Login(nonExistentCreds)
		assert.NotNil(t, err)
	})

	t.Run("ResetAccountPassword", func(t *testing.T) {
		CleanDB(t)
		// Create account first
		id := uuid.New()
		pID := uuid.New()
		p, _ := person.New(&pID, "Reset Test", "resettest@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
		r, _ := role.New(nil, "Admin", role.ADMIN_ROLE_CODE)
		acc, _ := account.New(&id, "resettest@example.com", "oldpass123", r, p, nil, nil)

		createdID, err := accountRepo.Create(acc)
		assert.Nil(t, err)
		assert.NotNil(t, createdID)

		// Reset password
		newPassword := "newpass456"
		err = repo.ResetAccountPassword(createdID, newPassword)
		assert.Nil(t, err)

		// Verify password was reset by logging in with new password
		creds := credentials.New("resettest@example.com", newPassword)
		loggedInAccount, err := repo.Login(creds)
		assert.Nil(t, err)
		assert.NotNil(t, loggedInAccount)
		assert.Equal(t, "resettest@example.com", loggedInAccount.Email())

		// Verify old password doesn't work
		oldCreds := credentials.New("resettest@example.com", "oldpass123")
		_, err = repo.Login(oldCreds)
		assert.NotNil(t, err)

		// Test reset with non-existent account ID
		nonExistentID := uuid.New()
		err = repo.ResetAccountPassword(&nonExistentID, "newpass")
		assert.NotNil(t, err)
	})
}
