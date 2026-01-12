package postgres

import (
	"eletronic_point/src/infra/repository"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func SetupTestDB(t *testing.T) {
	err := godotenv.Load("../../../../.env.test")
	if err != nil {
		t.Log("Warning: .env.test not found, assuming environment variables are already set")
	}

	// Check connection
	_, dbErr := repository.ExecQuery("SELECT 1")
	if dbErr != nil {
		t.Fatalf("Failed to connect to test database: %v", dbErr)
	}
}

func CleanDB(t *testing.T) {
	tables := []string{
		"time_record_status_movement",
		"time_record",
		"time_record_status",
		"internship",
		"internship_location",
		"student_linked_to_teacher",
		"student",
		"course",
		"campus",
		"institution",
		"account_role",
		"professional",
		"account",
		"person",
	}

	for _, table := range tables {
		_, err := repository.ExecQuery(fmt.Sprintf("TRUNCATE TABLE %s CASCADE", table))
		if err != nil {
			t.Logf("Warning: Failed to truncate table %s: %v", table, err)
		}
	}

	// Seed essential data
	SeedRoles(t)
	SeedStatuses(t)
}

func SeedRoles(t *testing.T) {
	roles := []struct {
		id   string
		name string
		code string
	}{
		{"e4f29304-62f7-4eb2-af4b-1156d6648196", "Administrador", "ADMIN"},
		{"fcf5df6c-340c-4e91-a20a-bc96160881a7", "Profissional", "PROFESSIONAL"},
		{"dc740240-e055-442c-b1eb-ac41e8dc7a3e", "Professor", "TEACHER"},
		{"3169e51e-9396-47eb-b4a5-47bdae043a30", "Estudante", "STUDENT"},
	}

	for _, r := range roles {
		_, err := repository.ExecQuery("INSERT INTO account_role (id, name, code) VALUES ($1, $2, $3) ON CONFLICT DO NOTHING", r.id, r.name, r.code)
		if err != nil {
			t.Fatalf("Failed to seed role %s: %v", r.code, err)
		}
	}
}

func SeedStatuses(t *testing.T) {
	statuses := []struct {
		id   string
		name string
	}{
		{"52613242-6b50-490a-9b4c-90cc3f263e9a", "pending"},
		{"faa4a69d-fe41-4ffe-b8d0-f752085f016a", "approved"},
		{"7f58a284-c8a5-4f89-a18e-320e8ea8960f", "disapproved"},
	}

	for _, s := range statuses {
		_, err := repository.ExecQuery("INSERT INTO time_record_status (id, name) VALUES ($1, $2) ON CONFLICT DO NOTHING", s.id, s.name)
		if err != nil {
			t.Fatalf("Failed to seed status %s: %v", s.name, err)
		}
	}
}
