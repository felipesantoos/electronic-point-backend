package response

import (
	"eletronic_point/src/core/domain/simplifiedStudent"

	"github.com/google/uuid"
)

type SimplifiedStudent struct {
	ID             *uuid.UUID  `json:"id"`
	Name           string      `json:"name"`
	ProfilePicture *string     `json:"profile_picture"`
	Institution    Institution `json:"institution"`
	Campus         Campus      `json:"campus"`
	Course         Course      `json:"course"`
	TotalWorkload  int         `json:"total_workload"`
}

type simplifiedStudentBuilder struct{}

func SimplifiedStudentBuilder() *simplifiedStudentBuilder {
	return &simplifiedStudentBuilder{}
}

func (*simplifiedStudentBuilder) BuildFromDomain(data simplifiedStudent.SimplifiedStudent) SimplifiedStudent {
	return SimplifiedStudent{
		ID:             data.ID(),
		Name:           data.Name(),
		ProfilePicture: data.ProfilePicture(),
		Institution:    InstitutionBuilder().BuildFromDomain(data.Institution()),
		Campus:         CampusBuilder().BuildFromDomain(data.Campus()),
		Course:         CourseBuilder().BuildFromDomain(data.Course()),
		TotalWorkload:  data.TotalWorkload(),
	}
}

func (*simplifiedStudentBuilder) BuildFromDomainList(data []simplifiedStudent.SimplifiedStudent) []SimplifiedStudent {
	students := make([]SimplifiedStudent, 0)
	for _, student := range data {
		students = append(students, SimplifiedStudentBuilder().BuildFromDomain(student))
	}
	return students
}
