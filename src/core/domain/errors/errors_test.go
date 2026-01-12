package errors

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError_NewFromString(t *testing.T) {
	msg := "error message"
	err := NewFromString(msg)
	assert.Equal(t, msg, err.String())
	assert.Equal(t, []string{msg}, err.Messages())
}

func TestError_Types(t *testing.T) {
	assert.True(t, NewInternal(errors.New("test")).CausedInternally())
	assert.True(t, NewValidation([]string{"test"}).CausedByValidation())
	assert.True(t, NewClient("test").CausedByClient())
	assert.True(t, NewForbidden("test").CausedByForbiddenAccess())
	assert.True(t, NewConflict("test").CausedByConflict())
	assert.True(t, NewUnauthorized("test").CausedByUnauthorization())
}

func TestError_Metadata(t *testing.T) {
	metadata := map[string]interface{}{"key": "value"}
	err := NewValidationWithMetadata([]string{"msg"}, metadata)
	assert.Equal(t, metadata, err.Metadata())
}

func TestError_ValidationMessagesByMetadataFields(t *testing.T) {
	err := NewValidationWithMetadata(
		[]string{"msg1", "msg2"},
		map[string]interface{}{"fields": []string{"field1", "field2"}},
	)
	
	msgs := err.ValidationMessagesByMetadataFields([]string{"field1"})
	assert.Equal(t, []string{"msg1"}, msgs)
}
