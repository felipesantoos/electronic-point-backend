package response

import (
	"eletronic_point/src/core/domain/campus"

	"github.com/google/uuid"
)

type Campus struct {
	ID            uuid.UUID  `json:"id"`
	Name          string     `json:"name"`
	InstitutionID *uuid.UUID `json:"institution_id,omitempty"`
}

type campusBuilder struct{}

func CampusBuilder() *campusBuilder {
	return &campusBuilder{}
}

func (*campusBuilder) BuildFromDomain(data campus.Campus) Campus {
	var institutionID *uuid.UUID
	if data.InstitutionID() != uuid.Nil {
		aux := data.InstitutionID()
		institutionID = &aux
	}
	return Campus{
		ID:            data.ID(),
		Name:          data.Name(),
		InstitutionID: institutionID,
	}
}

func (*campusBuilder) BuildFromDomainList(data []campus.Campus) []Campus {
	campuses := make([]Campus, 0)
	for _, _campus := range data {
		campuses = append(campuses, CampusBuilder().BuildFromDomain(_campus))
	}
	return campuses
}
