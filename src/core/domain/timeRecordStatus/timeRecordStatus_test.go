package timeRecordStatus

import (
	"eletronic_point/src/core/messages"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTimeRecordStatus_Setters(t *testing.T) {
	id := uuid.New()
	name := "Pending"
	trs := &timeRecordStatus{}
	
	err := trs.SetID(id)
	assert.Nil(t, err)
	assert.Equal(t, id, trs.id)

	err = trs.SetName(name)
	assert.Nil(t, err)
	assert.Equal(t, name, trs.name)
}

func TestTimeRecordStatus_Setters_Errors(t *testing.T) {
	trs := &timeRecordStatus{}
	
	err := trs.SetID(uuid.Nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.TimeRecordStatusIDErrorMessage)

	err = trs.SetName("")
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.TimeRecordStatusNameErrorMessage)
}
