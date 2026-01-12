package institution

import (
	"eletronic_point/src/core/messages"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInstitution_Setters(t *testing.T) {
	id := uuid.New()
	name := "Institution A"
	
	i := &institution{}
	
	err := i.SetID(id)
	assert.Nil(t, err)
	assert.Equal(t, id, i.ID())

	err = i.SetName(name)
	assert.Nil(t, err)
	assert.Equal(t, name, i.Name())
}

func TestInstitution_Setters_Errors(t *testing.T) {
	i := &institution{}
	
	err := i.SetID(uuid.Nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.InstitutionIDErrorMessage)

	err = i.SetName("")
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.InstitutionNameErrorMessage)
}
