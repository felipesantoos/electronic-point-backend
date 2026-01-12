package campus

import (
	"eletronic_point/src/core/messages"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCampus_Setters(t *testing.T) {
	id := uuid.New()
	instID := uuid.New()
	name := "Campus A"
	
	c := &campus{}
	
	err := c.SetID(id)
	assert.Nil(t, err)
	assert.Equal(t, id, c.ID())

	err = c.SetName(name)
	assert.Nil(t, err)
	assert.Equal(t, name, c.Name())

	err = c.SetInstitutionID(instID)
	assert.Nil(t, err)
	assert.Equal(t, instID, c.InstitutionID())
}

func TestCampus_Setters_Errors(t *testing.T) {
	c := &campus{}
	
	err := c.SetID(uuid.Nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.CampusIDErrorMessage)

	err = c.SetName("")
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.CampusNameErrorMessage)

	err = c.SetInstitutionID(uuid.Nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.CampusInstitutionIDErrorMessage)
}
