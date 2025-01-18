package response

import (
	"eletronic_point/src/core/domain/institution"

	"github.com/google/uuid"
)

type Institution struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type institutionBuilder struct{}

func InstitutionBuilder() *institutionBuilder {
	return &institutionBuilder{}
}

func (*institutionBuilder) BuildFromDomain(data institution.Institution) Institution {
	return Institution{
		ID:   data.ID(),
		Name: data.Name(),
	}
}

func (*institutionBuilder) BuildFromDomainList(data []institution.Institution) []Institution {
	institutions := make([]Institution, 0)
	for _, _institution := range data {
		institutions = append(institutions, InstitutionBuilder().BuildFromDomain(_institution))
	}
	return institutions
}
