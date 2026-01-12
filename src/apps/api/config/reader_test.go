package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReader(t *testing.T) {
	// Set environment variables
	os.Setenv("SERVER_HOST", "localhost")
	os.Setenv("SERVER_PORT", "8080")
	os.Setenv("DATABASE_USER", "user")
	os.Setenv("DATABASE_PASSWORD", "pass")
	os.Setenv("DATABASE_NAME", "db")
	os.Setenv("DATABASE_HOST", "localhost")
	os.Setenv("DATABASE_PORT", "5432")
	os.Setenv("DATABASE_SSL_MODE", "disable")
	os.Setenv("DATABASE_MIGRATIONS_PATH", "./migrations")
	os.Setenv("AUTH_MODEL_PATH", "./model.conf")
	os.Setenv("AUTH_POLICY_PATH", "./policy.csv")
	os.Setenv("SERVER_ALLOWED_HOSTS", "localhost,example.com")

	defer func() {
		os.Unsetenv("SERVER_HOST")
		os.Unsetenv("SERVER_PORT")
		os.Unsetenv("DATABASE_USER")
		os.Unsetenv("DATABASE_PASSWORD")
		os.Unsetenv("DATABASE_NAME")
		os.Unsetenv("DATABASE_HOST")
		os.Unsetenv("DATABASE_PORT")
		os.Unsetenv("DATABASE_SSL_MODE")
		os.Unsetenv("DATABASE_MIGRATIONS_PATH")
		os.Unsetenv("AUTH_MODEL_PATH")
		os.Unsetenv("AUTH_POLICY_PATH")
		os.Unsetenv("SERVER_ALLOWED_HOSTS")
	}()

	r := reader{}
	cfg, err := r.config()

	assert.Nil(t, err)
	assert.NotNil(t, cfg)
	assert.Equal(t, "localhost", cfg.HttpServer.HostAddress)
	assert.Equal(t, 8080, cfg.HttpServer.Port)
	assert.Equal(t, "user", cfg.Postgres.User)
	assert.Equal(t, "localhost,example.com", os.Getenv("SERVER_ALLOWED_HOSTS"))
	assert.Equal(t, []string{"localhost", "example.com"}, cfg.CORSPolicy.ServerAllowedHosts)
}

func TestReader_Errors(t *testing.T) {
	// Clear environment variables
	os.Clearenv()

	r := reader{}
	cfg, err := r.config()

	assert.NotNil(t, err)
	assert.Nil(t, cfg)
}
