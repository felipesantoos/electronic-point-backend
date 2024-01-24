package response

import (
	"backend_template/src/core/domain/professional"

	"github.com/google/uuid"
)

type Professional struct {
	ID *uuid.UUID `json:"id"`
}

type professionalBuilder struct{}

func ProfessionalBuilder() *professionalBuilder {
	return &professionalBuilder{}
}

func (*professionalBuilder) BuildFromDomain(data professional.Professional) *Professional {
	return &Professional{
		ID: data.ID(),
	}
}
