package request

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/student"
)

type Student struct {
	Name                   string  `form:"name"`
	BirthDate              string  `form:"birth_date"`
	CPF                    string  `form:"cpf"`
	Email                  string  `form:"email"`
	Phone                  string  `form:"phone"`
	Registration           string  `form:"registration"`
	ProfilePicture         *string `form:"profile_picture"`
	Institution            string  `form:"institution"`
	Course                 string  `form:"course"`
	InternshipLocationName string  `form:"internship_location_name"`
	InternshipAddress      string  `form:"internship_address"`
	InternshipLocation     string  `form:"internship_location"`
	TotalWorkload          int     `form:"total_workload"`
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
	builder := student.NewBuilder().
		WithPerson(_person).
		WithRegistration(this.Registration).
		WithProfilePicture(this.ProfilePicture).
		WithInstitution(this.Institution).
		WithCourse(this.Course).
		WithInternshipLocationName(this.InternshipLocationName).
		WithInternshipAddress(this.InternshipAddress).
		WithInternshipLocation(this.InternshipLocation).
		WithTotalWorkload(this.TotalWorkload)
	return builder.Build()
}
