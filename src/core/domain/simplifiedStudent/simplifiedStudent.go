package simplifiedStudent

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/utils/validator"

	"github.com/google/uuid"
)

var _ SimplifiedStudent = &simplifiedStudent{}

type simplifiedStudent struct {
	id             *uuid.UUID
	name           string
	profilePicture *string
	_institution   institution.Institution
	_campus        campus.Campus
	_course        course.Course
	totalWorkload  int
}

func (s *simplifiedStudent) ID() *uuid.UUID {
	return s.id
}

func (s *simplifiedStudent) Name() string {
	return s.name
}

func (s *simplifiedStudent) ProfilePicture() *string {
	return s.profilePicture
}

func (s *simplifiedStudent) Institution() institution.Institution {
	return s._institution
}

func (s *simplifiedStudent) Campus() campus.Campus {
	return s._campus
}

func (s *simplifiedStudent) Course() course.Course {
	return s._course
}

func (s *simplifiedStudent) TotalWorkload() int {
	return s.totalWorkload
}

func (s *simplifiedStudent) SetID(id *uuid.UUID) errors.Error {
	if !validator.IsUUIDValid(*id) {
		return errors.NewValidationFromString(messages.StudentIDErrorMessage)
	}
	s.id = id
	return nil
}

func (s *simplifiedStudent) SetName(name string) errors.Error {
	if name == "" {
		return errors.NewFromString(messages.StudentNameErrorMessage)
	}
	s.name = name
	return nil
}

func (s *simplifiedStudent) SetProfilePicture(profilePicture *string) errors.Error {
	s.profilePicture = profilePicture
	return nil
}

func (s *simplifiedStudent) SetInstitution(_institution institution.Institution) errors.Error {
	if _institution == nil {
		return errors.NewFromString(messages.StudentInstitutionErrorMessage)
	}
	s._institution = _institution
	return nil
}

func (s *simplifiedStudent) SetCampus(_campus campus.Campus) errors.Error {
	if _campus == nil {
		return errors.NewFromString(messages.StudentCampusErrorMessage)
	}
	s._campus = _campus
	return nil
}

func (s *simplifiedStudent) SetCourse(_course course.Course) errors.Error {
	if _course == nil {
		return errors.NewFromString(messages.StudentCourseErrorMessage)
	}
	s._course = _course
	return nil
}

func (s *simplifiedStudent) SetTotalWorkload(totalWorkload int) errors.Error {
	s.totalWorkload = totalWorkload
	return nil
}
