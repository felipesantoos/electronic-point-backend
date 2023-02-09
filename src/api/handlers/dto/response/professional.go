package response

import (
	"dit_backend/src/core/domain/professional"

	"github.com/google/uuid"
)

type Professional struct {
	ID       *uuid.UUID `json:"id" novalidate:"true"`
	PersonID *uuid.UUID `json:"person_id" novalidate:"true"`
}

type professionalBuilder struct{}

func ProfessionalBuilder() *professionalBuilder {
	return &professionalBuilder{}
}

func (*professionalBuilder) FromDomain(data professional.Professional) *Professional {
	return &Professional{
		ID:       data.ID(),
		PersonID: data.PersonID(),
	}
}
