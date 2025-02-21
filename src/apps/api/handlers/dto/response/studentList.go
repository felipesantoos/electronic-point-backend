package response

import (
	"eletronic_point/src/core/domain/student"
)

type StudentList struct {
	Person
	Registration      string      `json:"registration"`
	ProfilePicture    *string     `json:"profile_picture"`
	Institution       Institution `json:"institution"`
	Campus            Campus      `json:"campus"`
	Course            Course      `json:"course"`
	TotalWorkload     int         `json:"total_workload"`
	WorkloadCompleted int         `json:"workload_completed"`
	PendingWorkload   int         `json:"pending_workload"`
	CurrentInternship *Internship `json:"current_internship"`
}

type studentListBuilder struct{}

func StudentListBuilder() *studentListBuilder {
	return &studentListBuilder{}
}

func (*studentListBuilder) BuildFromDomain(data student.Student) StudentList {
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
	return StudentList{
		Person:            _person,
		Registration:      data.Registration(),
		ProfilePicture:    data.ProfilePicture(),
		Institution:       InstitutionBuilder().BuildFromDomain(data.Institution()),
		Campus:            CampusBuilder().BuildFromDomain(data.Campus()),
		Course:            CourseBuilder().BuildFromDomain(data.Course()),
		TotalWorkload:     data.TotalWorkload(),
		WorkloadCompleted: data.WorkloadCompleted(),
		PendingWorkload:   data.PendingWorkload(),
		CurrentInternship: currentInternship,
	}
}

func (*studentListBuilder) BuildFromDomainList(data []student.Student) []StudentList {
	students := make([]StudentList, 0)
	for _, student := range data {
		students = append(students, StudentListBuilder().BuildFromDomain(student))
	}
	return students
}
