package response

import (
	"eletronic_point/src/core/domain/internshipLocation"

	"github.com/google/uuid"
)

type InternshipLocation struct {
	ID           uuid.UUID `json:"id"`
	Name         string    `json:"name"`
	Number       string    `json:"number"`
	Street       string    `json:"street"`
	Neighborhood string    `json:"neighborhood"`
	City         string    `json:"city"`
	ZipCode      string    `json:"zip_code"`
	Lat          float64   `json:"lat"`
	Long         float64   `json:"long"`
}

type internshipLocationBuilder struct{}

func InternshipLocationBuilder() *internshipLocationBuilder {
	return &internshipLocationBuilder{}
}

func (*internshipLocationBuilder) BuildFromDomain(data internshipLocation.InternshipLocation) InternshipLocation {
	return InternshipLocation{
		ID:           data.ID(),
		Name:         data.Name(),
		Number:       data.Number(),
		Street:       data.Street(),
		Neighborhood: data.Neighborhood(),
		City:         data.City(),
		ZipCode:      data.ZipCode(),
		Lat:          data.Lat(),
		Long:         data.Long(),
	}
}

func (*internshipLocationBuilder) BuildFromDomainList(data []internshipLocation.InternshipLocation) []InternshipLocation {
	internshipLocations := make([]InternshipLocation, 0)
	for _, location := range data {
		internshipLocations = append(internshipLocations, InternshipLocationBuilder().BuildFromDomain(location))
	}
	return internshipLocations
}
