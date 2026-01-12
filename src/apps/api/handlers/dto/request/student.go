package request

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/student"

	"github.com/google/uuid"
)

type Student struct {
	Name                 string     `form:"name"`
	BirthDate            string     `form:"birth_date"`
	CPF                  string     `form:"cpf"`
	Email                string     `form:"email"`
	Phone                string     `form:"phone"`
	Registration         string     `form:"registration"`
	ProfilePicture       *string    `form:"profile_picture"`
	CampusID             uuid.UUID  `form:"campus_id"`
	CourseID             uuid.UUID  `form:"course_id"`
	TotalWorkload        int        `form:"total_workload"`
	ResponsibleTeacherID *uuid.UUID `form:"responsible_teacher_id"`
}

func (this *Student) ToDomain() (student.Student, errors.Error) {
	_person, validationError := person.NewBuilder().
		WithName(this.Name).
		WithBirthDate(this.BirthDate).
		WithCPF(this.CPF).
		WithEmail(this.Email).
		WithPhone(this.Phone).Build()
	if validationError != nil {
		return nil, validationError
	}
	_campus, validationError := campus.NewBuilder().WithID(this.CampusID).Build()
	if validationError != nil {
		return nil, validationError
	}
	_course, validationError := course.NewBuilder().WithID(this.CourseID).Build()
	if validationError != nil {
		return nil, validationError
	}
	builder := student.NewBuilder().
		WithPerson(_person).
		WithRegistration(this.Registration).
		WithProfilePicture(this.ProfilePicture).
		WithCampus(_campus).
		WithCourse(_course).
		WithTotalWorkload(this.TotalWorkload)
	return builder.Build()
}
