package response

import (
	"eletronic_point/src/core/domain/internship"
	"time"
)

type Internship struct {
	StartedIn time.Time          `json:"started_in"`
	EndedIn   *time.Time         `json:"ended_in"`
	Location  InternshipLocation `json:"location"`
}

type internshipBuilder struct{}

func InternshipBuilder() *internshipBuilder {
	return &internshipBuilder{}
}

func (*internshipBuilder) BuildFromDomain(data internship.Internship) Internship {
	var location InternshipLocation
	if data.Location() != nil {
		location = InternshipLocationBuilder().BuildFromDomain(data.Location())
	}
	return Internship{
		StartedIn: data.StartedIn(),
		EndedIn:   data.EndedIn(),
		Location:  location,
	}
}

func (*internshipBuilder) BuildFromDomainList(data []internship.Internship) []Internship {
	internships := make([]Internship, 0)
	for _, _internship := range data {
		internships = append(internships, InternshipBuilder().BuildFromDomain(_internship))
	}
	return internships
}
