package request

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/student"
)

type Student struct {
	Name                   string  `form:"name"`
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
	builder := student.NewBuilder().
		WithName(this.Name).
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
