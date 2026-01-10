package simplifiedStudent

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/institution"

	"github.com/google/uuid"
)

type SimplifiedStudent interface {
	ID() *uuid.UUID
	Name() string
	ProfilePicture() *string
	Institution() institution.Institution
	Campus() campus.Campus
	Course() course.Course
	TotalWorkload() int

	SetID(*uuid.UUID) errors.Error
	SetName(string) errors.Error
	SetProfilePicture(*string) errors.Error
	SetInstitution(institution.Institution) errors.Error
	SetCampus(campus.Campus) errors.Error
	SetCourse(course.Course) errors.Error
	SetTotalWorkload(int) errors.Error
}
