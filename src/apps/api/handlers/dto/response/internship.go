package response

import (
	"eletronic_point/src/core/domain/internship"
	"time"

	"github.com/google/uuid"
)

type Internship struct {
	ID                uuid.UUID          `json:"id"`
	StartedIn         time.Time          `json:"started_in"`
	EndedIn           *time.Time         `json:"ended_in"`
	ScheduleEntryTime *time.Time         `json:"schedule_entry_time"`
	ScheduleExitTime  *time.Time         `json:"schedule_exit_time"`
	Location          InternshipLocation `json:"location"`
	Student           *SimplifiedStudent `json:"student,omitempty"`
}

type internshipBuilder struct{}

func InternshipBuilder() *internshipBuilder {
	return &internshipBuilder{}
}

func (*internshipBuilder) BuildFromDomain(data internship.Internship) Internship {
	if data == nil {
		return Internship{}
	}
	var location InternshipLocation
	if data.Location() != nil {
		location = InternshipLocationBuilder().BuildFromDomain(data.Location())
	}
	var _student *SimplifiedStudent
	if data.Student() != nil {
		if data.Student().ID() != nil {
			aux := SimplifiedStudentBuilder().BuildFromDomain(data.Student())
			_student = &aux
		}
	}
	return Internship{
		ID:                data.ID(),
		StartedIn:         data.StartedIn(),
		EndedIn:           data.EndedIn(),
		ScheduleEntryTime: data.ScheduleEntryTime(),
		ScheduleExitTime:  data.ScheduleExitTime(),
		Location:          location,
		Student:           _student,
	}
}

func (*internshipBuilder) BuildFromDomainList(data []internship.Internship) []Internship {
	internships := make([]Internship, 0)
	for _, _internship := range data {
		internships = append(internships, InternshipBuilder().BuildFromDomain(_internship))
	}
	return internships
}
