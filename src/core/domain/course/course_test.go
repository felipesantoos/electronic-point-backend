package course

import (
	"eletronic_point/src/core/messages"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestCourse_Setters(t *testing.T) {
	id := uuid.New()
	name := "Course A"
	
	c := &course{}
	
	err := c.SetID(id)
	assert.Nil(t, err)
	assert.Equal(t, id, c.ID())

	err = c.SetName(name)
	assert.Nil(t, err)
	assert.Equal(t, name, c.Name())
}

func TestCourse_Setters_Errors(t *testing.T) {
	c := &course{}
	
	err := c.SetID(uuid.Nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.CourseIDErrorMessage)

	err = c.SetName("")
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.CourseNameErrorMessage)
}
