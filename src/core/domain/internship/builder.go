package internship

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/utils/validator"
	"time"

	"github.com/google/uuid"
)

type Builder struct {
	fields        []string
	errorMessages []string
	internship    *internship
}

func NewBuilder() *Builder {
	return &Builder{
		fields:        []string{},
		errorMessages: []string{},
		internship:    &internship{},
	}
}

func (b *Builder) WithID(id uuid.UUID) *Builder {
	if !validator.IsUUIDValid(id) {
		b.fields = append(b.fields, messages.InternshipID)
		b.errorMessages = append(b.errorMessages, messages.InternshipIDErrorMessage)
		return b
	}
	b.internship.id = id
	return b
}

func (b *Builder) WithStartedIn(startedIn time.Time) *Builder {
	if startedIn.IsZero() {
		b.fields = append(b.fields, messages.InternshipStartedIn)
		b.errorMessages = append(b.errorMessages, messages.InternshipStartedInErrorMessage)
		return b
	}
	b.internship.startedIn = startedIn
	return b
}

func (b *Builder) WithEndedIn(endedIn *time.Time) *Builder {
	b.internship.endedIn = endedIn
	return b
}

func (b *Builder) WithLocation(location internshipLocation.InternshipLocation) *Builder {
	if location == nil {
		b.fields = append(b.fields, messages.InternshipLocation)
		b.errorMessages = append(b.errorMessages, messages.InternshipLocationErrorMessage)
		return b
	}
	b.internship.location = location
	return b
}

func (b *Builder) Build() (Internship, errors.Error) {
	if len(b.errorMessages) > 0 {
		return nil, errors.NewValidationWithMetadata(b.errorMessages, map[string]interface{}{
			messages.Fields: b.fields,
		})
	}
	return b.internship, nil
}
