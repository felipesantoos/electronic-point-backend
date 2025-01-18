package response

import (
	"eletronic_point/src/core/domain/student"
)

type Student struct {
	Person
	Registration      string       `json:"registration"`
	ProfilePicture    *string      `json:"profile_picture"`
	Institution       Institution  `json:"institution"`
	Campus            Campus       `json:"campus"`
	Course            string       `json:"course"`
	TotalWorkload     int          `json:"total_workload"`
	WorkloadCompleted int          `json:"workload_completed"`
	PendingWorkload   int          `json:"pending_workload"`
	CurrentInternship *Internship  `json:"current_internship"`
	InternshipHistory []Internship `json:"internship_history"`
	FrequencyHistory  []TimeRecord `json:"frequency_history"`
}

type studentBuilder struct{}

func StudentBuilder() *studentBuilder {
	return &studentBuilder{}
}

func (*studentBuilder) BuildFromDomain(data student.Student) Student {
	_person := Person{
		ID:        data.ID(),
		Name:      data.Name(),
		BirthDate: data.BirthDate(),
		Email:     data.Email(),
		CPF:       data.CPF(),
		Phone:     data.Phone(),
	}
	var currentInternship *Internship
	if data.CurrentInternship() != nil {
		aux := InternshipBuilder().BuildFromDomain(data.CurrentInternship())
		currentInternship = &aux
	}
	return Student{
		Person:            _person,
		Registration:      data.Registration(),
		ProfilePicture:    data.ProfilePicture(),
		Institution:       InstitutionBuilder().BuildFromDomain(data.Institution()),
		Campus:            CampusBuilder().BuildFromDomain(data.Campus()),
		Course:            data.Course(),
		TotalWorkload:     data.TotalWorkload(),
		WorkloadCompleted: data.WorkloadCompleted(),
		PendingWorkload:   data.PendingWorkload(),
		CurrentInternship: currentInternship,
		InternshipHistory: InternshipBuilder().BuildFromDomainList(data.InternshipHistory()),
		FrequencyHistory:  TimeRecordBuilder().BuildFromDomainList(data.FrequencyHistory()),
	}
}

func (*studentBuilder) BuildFromDomainList(data []student.Student) []Student {
	students := make([]Student, 0)
	for _, student := range data {
		students = append(students, StudentBuilder().BuildFromDomain(student))
	}
	return students
}
