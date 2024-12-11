package response

import (
	"eletronic_point/src/core/domain/student"

	"github.com/google/uuid"
)

type Student struct {
	ID                     uuid.UUID `json:"id"`
	Name                   string    `json:"name"`
	Registration           string    `json:"registration"`
	ProfilePicture         *string   `json:"profile_picture"`
	Institution            string    `json:"institution"`
	Course                 string    `json:"course"`
	InternshipLocationName string    `json:"internship_location_name"`
	InternshipAddress      string    `json:"internship_address"`
	InternshipLocation     string    `json:"internship_location"`
	TotalWorkload          int       `json:"total_workload"`
	WorkloadCompleted      int       `json:"workload_completed"`
	PendingWorkload        int       `json:"pending_workload"`
	FrequencyHistory       string    `json:"frequency_history"`
}

type studentBuilder struct{}

func StudentBuilder() *studentBuilder {
	return &studentBuilder{}
}

func (*studentBuilder) BuildFromDomain(data student.Student) Student {
	return Student{
		ID:                     data.ID(),
		Name:                   data.Name(),
		Registration:           data.Registration(),
		ProfilePicture:         data.ProfilePicture(),
		Institution:            data.Institution(),
		Course:                 data.Course(),
		InternshipLocationName: data.InternshipLocationName(),
		InternshipAddress:      data.InternshipAddress(),
		InternshipLocation:     data.InternshipLocation(),
		TotalWorkload:          data.TotalWorkload(),
		WorkloadCompleted:      data.WorkloadCompleted(),
		PendingWorkload:        data.PendingWorkload(),
		FrequencyHistory:       data.FrequencyHistory(),
	}
}

func (*studentBuilder) BuildFromDomainList(data []student.Student) []Student {
	var students []Student
	for _, student := range data {
		students = append(students, StudentBuilder().BuildFromDomain(student))
	}
	return students
}
