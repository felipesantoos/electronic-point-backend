package timeRecord

import (
	"eletronic_point/src/core/messages"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_WithID(t *testing.T) {
	tests := []struct {
		name          string
		id            uuid.UUID
		expectedError bool
	}{
		{"Valid ID", uuid.New(), false},
		{"Nil ID", uuid.Nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithID(tt.id)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.TimeRecordIDErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_WithDate(t *testing.T) {
	tests := []struct {
		name          string
		date          time.Time
		expectedError bool
	}{
		{"Valid Date", time.Now(), false},
		{"Zero Date", time.Time{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithDate(tt.date)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.TimeRecordDateErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_WithEntryTime(t *testing.T) {
	tests := []struct {
		name          string
		entryTime     time.Time
		expectedError bool
	}{
		{"Valid Entry Time", time.Now(), false},
		{"Zero Entry Time", time.Time{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithEntryTime(tt.entryTime)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.TimeRecordEntryTimeErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_WithLocation(t *testing.T) {
	tests := []struct {
		name          string
		location      string
		expectedError bool
	}{
		{"Valid Location", "On-site", false},
		{"Empty Location", "", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBuilder().WithLocation(tt.location)
			_, err := b.Build()

			if tt.expectedError {
				assert.NotNil(t, err)
				assert.Contains(t, err.Messages(), messages.TimeRecordLocationErrorMessage)
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestBuilder_Build(t *testing.T) {
	t.Run("Build with Multiple Errors", func(t *testing.T) {
		_, err := NewBuilder().
			WithID(uuid.Nil).
			WithDate(time.Time{}).
			WithEntryTime(time.Time{}).
			WithLocation("").
			Build()

		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.TimeRecordIDErrorMessage)
		assert.Contains(t, err.Messages(), messages.TimeRecordDateErrorMessage)
		assert.Contains(t, err.Messages(), messages.TimeRecordEntryTimeErrorMessage)
		assert.Contains(t, err.Messages(), messages.TimeRecordLocationErrorMessage)
	})
}
