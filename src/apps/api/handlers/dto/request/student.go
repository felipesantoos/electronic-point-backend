package request

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/student"
)

type CreateStudent struct {
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

func (c *CreateStudent) ToDomain() (student.Student, errors.Error) {
	builder := student.NewBuilder().
		WithName(c.Name).
		WithRegistration(c.Registration).
		WithProfilePicture(c.ProfilePicture).
		WithInstitution(c.Institution).
		WithCourse(c.Course).
		WithInternshipLocationName(c.InternshipLocationName).
		WithInternshipAddress(c.InternshipAddress).
		WithInternshipLocation(c.InternshipLocation).
		WithTotalWorkload(c.TotalWorkload)
	return builder.Build()
}
