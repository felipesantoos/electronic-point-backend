package student

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/timeRecord"
	"eletronic_point/src/core/messages"

	"github.com/google/uuid"
)

var _ Student = &student{}

type student struct {
	person.Person
	studentID              uuid.UUID
	name                   string
	registration           string
	profilePicture         *string
	institution            string
	course                 string
	internshipLocationName string
	internshipAddress      string
	internshipLocation     string
	totalWorkload          int
	workloadCompleted      int
	pendingWorkload        int
	frequencyHistory       []timeRecord.TimeRecord
}

func (s *student) StudentID() uuid.UUID {
	return s.studentID
}

func (s *student) Registration() string {
	return s.registration
}

func (s *student) ProfilePicture() *string {
	return s.profilePicture
}

func (s *student) Institution() string {
	return s.institution
}

func (s *student) Course() string {
	return s.course
}

func (s *student) InternshipLocationName() string {
	return s.internshipLocationName
}

func (s *student) InternshipAddress() string {
	return s.internshipAddress
}

func (s *student) InternshipLocation() string {
	return s.internshipLocation
}

func (s *student) TotalWorkload() int {
	return s.totalWorkload
}

func (s *student) WorkloadCompleted() int {
	return s.workloadCompleted
}

func (s *student) PendingWorkload() int {
	return s.pendingWorkload
}

func (s *student) FrequencyHistory() []timeRecord.TimeRecord {
	return s.frequencyHistory
}

func (s *student) SetStudentID(studentID uuid.UUID) errors.Error {
	if studentID == uuid.Nil {
		return errors.NewFromString(messages.StudentIDErrorMessage)
	}
	s.studentID = studentID
	return nil
}

func (s *student) SetName(name string) errors.Error {
	if name == "" {
		return errors.NewFromString(messages.StudentNameErrorMessage)
	}
	s.name = name
	return nil
}

func (s *student) SetRegistration(registration string) errors.Error {
	if registration == "" {
		return errors.NewFromString(messages.StudentRegistrationErrorMessage)
	}
	s.registration = registration
	return nil
}

func (s *student) SetProfilePicture(profilePicture *string) errors.Error {
	s.profilePicture = profilePicture
	return nil
}

func (s *student) SetInstitution(institution string) errors.Error {
	if institution == "" {
		return errors.NewFromString(messages.StudentInstitutionErrorMessage)
	}
	s.institution = institution
	return nil
}

func (s *student) SetCourse(course string) errors.Error {
	if course == "" {
		return errors.NewFromString(messages.StudentCourseErrorMessage)
	}
	s.course = course
	return nil
}

func (s *student) SetInternshipLocationName(locationName string) errors.Error {
	if locationName == "" {
		return errors.NewFromString(messages.StudentInternshipLocationNameErrorMessage)
	}
	s.internshipLocationName = locationName
	return nil
}

func (s *student) SetInternshipAddress(address string) errors.Error {
	if address == "" {
		return errors.NewFromString(messages.StudentInternshipAddressErrorMessage)
	}
	s.internshipAddress = address
	return nil
}

func (s *student) SetInternshipLocation(location string) errors.Error {
	if location == "" {
		return errors.NewFromString(messages.StudentInternshipLocationErrorMessage)
	}
	s.internshipLocation = location
	return nil
}

func (s *student) SetTotalWorkload(totalWorkload int) errors.Error {
	if totalWorkload < 0 {
		return errors.NewFromString(messages.StudentTotalWorkloadErrorMessage)
	}
	s.totalWorkload = totalWorkload
	return nil
}

func (s *student) SetWorkloadCompleted(workloadCompleted int) errors.Error {
	if workloadCompleted < 0 {
		return errors.NewFromString(messages.StudentWorkloadCompletedErrorMessage)
	}
	s.workloadCompleted = workloadCompleted
	return nil
}

func (s *student) SetPendingWorkload(pendingWorkload int) errors.Error {
	if pendingWorkload < 0 {
		return errors.NewFromString(messages.StudentPendingWorkloadErrorMessage)
	}
	s.pendingWorkload = pendingWorkload
	return nil
}

func (s *student) SetFrequencyHistory(frequencyHistory []timeRecord.TimeRecord) errors.Error {
	s.frequencyHistory = frequencyHistory
	return nil
}
