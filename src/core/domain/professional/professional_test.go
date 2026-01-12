package professional

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestProfessional_Setters(t *testing.T) {
	id := uuid.New()
	pID := uuid.New()
	p := &professional{}
	
	p.SetPersonID(&pID)
	p.id = &id

	assert.Equal(t, id, *p.ID())
	assert.Equal(t, pID, *p.PersonID())
}

func TestProfessional_IsValid(t *testing.T) {
	t.Run("Valid Professional", func(t *testing.T) {
		id := uuid.New()
		personID := uuid.New()
		p := &professional{
			id:       &id,
			personID: &personID,
		}
		assert.Nil(t, p.IsValid())
	})

	t.Run("Nil ID and PersonID", func(t *testing.T) {
		p := &professional{}
		assert.Nil(t, p.IsValid())
	})

	t.Run("Only ID", func(t *testing.T) {
		id := uuid.New()
		p := &professional{
			id: &id,
		}
		assert.Nil(t, p.IsValid())
	})

	t.Run("Only PersonID", func(t *testing.T) {
		personID := uuid.New()
		p := &professional{
			personID: &personID,
		}
		assert.Nil(t, p.IsValid())
	})
}

func TestNew(t *testing.T) {
	t.Run("Successful Creation", func(t *testing.T) {
		id := uuid.New()
		personID := uuid.New()
		prof, err := New(&id, &personID)

		assert.Nil(t, err)
		assert.NotNil(t, prof)
		assert.Equal(t, id, *prof.ID())
		assert.Equal(t, personID, *prof.PersonID())
	})

	t.Run("Creation with Nil IDs", func(t *testing.T) {
		prof, err := New(nil, nil)

		assert.Nil(t, err)
		assert.NotNil(t, prof)
		assert.Nil(t, prof.ID())
		assert.Nil(t, prof.PersonID())
	})
}
