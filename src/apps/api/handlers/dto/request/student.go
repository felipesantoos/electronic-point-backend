package request

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/student"
)

type Student struct {
	Name                   string  `json:"name"`
	Registration           string  `json:"registration"`
	ProfilePicture         *string `json:"profile_picture"`
	Institution            string  `json:"institution"`
	Course                 string  `json:"course"`
	InternshipLocationName string  `json:"internship_location_name"`
	InternshipAddress      string  `json:"internship_address"`
	InternshipLocation     string  `json:"internship_location"`
	TotalWorkload          int     `json:"total_workload"`
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
