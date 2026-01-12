package internshipLocation

import (
	"eletronic_point/src/core/messages"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInternshipLocation_Setters(t *testing.T) {
	id := uuid.New()
	name := "Location A"
	i := &internshipLocation{}
	
	err := i.SetID(id)
	assert.Nil(t, err)
	assert.Equal(t, id, i.ID())

	err = i.SetName(name)
	assert.Nil(t, err)
	assert.Equal(t, name, i.Name())
}

func TestInternshipLocation_Setters_Errors(t *testing.T) {
	i := &internshipLocation{}
	
	err := i.SetID(uuid.Nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.InternshipLocationIDErrorMessage)

	err = i.SetName("")
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.InternshipLocationNameErrorMessage)
}
