package student

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/timeRecord"
	"eletronic_point/src/core/messages"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStudent_Setters(t *testing.T) {
	s := &student{}

	err := s.SetName("John Doe")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe", s.name)

	err = s.SetRegistration("2023101010")
	assert.Nil(t, err)
	assert.Equal(t, "2023101010", s.Registration())

	err = s.SetTotalWorkload(100)
	assert.Nil(t, err)
	assert.Equal(t, 100, s.TotalWorkload())

	// Test SetProfilePicture
	profilePic := "http://example.com/pic.jpg"
	err = s.SetProfilePicture(&profilePic)
	assert.Nil(t, err)
	assert.Equal(t, &profilePic, s.ProfilePicture())

	// Test SetProfilePicture nil
	err = s.SetProfilePicture(nil)
	assert.Nil(t, err)
	assert.Nil(t, s.ProfilePicture())

	// Test SetWorkloadCompleted
	err = s.SetWorkloadCompleted(50)
	assert.Nil(t, err)
	assert.Equal(t, 50, s.WorkloadCompleted())

	// Test SetPendingWorkload
	err = s.SetPendingWorkload(50)
	assert.Nil(t, err)
	assert.Equal(t, 50, s.PendingWorkload())

	// Test SetResponsibleTeacherID
	teacherID := uuid.New()
	err = s.SetResponsibleTeacherID(teacherID)
	assert.Nil(t, err)
	assert.Equal(t, teacherID, s.ResponsibleTeacherID())
}

func TestStudent_Setters_Errors(t *testing.T) {
	s := &student{}

	err := s.SetName("")
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.StudentNameErrorMessage)

	err = s.SetRegistration("")
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.StudentRegistrationErrorMessage)

	err = s.SetTotalWorkload(-1)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.StudentTotalWorkloadErrorMessage)

	err = s.SetWorkloadCompleted(-1)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.StudentWorkloadCompletedErrorMessage)

	err = s.SetPendingWorkload(-1)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.StudentPendingWorkloadErrorMessage)

	err = s.SetResponsibleTeacherID(uuid.Nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.StudentResponsibleTeacherIDErrorMessage)
}

func TestStudent_IsValid(t *testing.T) {
	t.Run("Valid Student with Valid Person", func(t *testing.T) {
		p, _ := person.New(nil, "John Doe", "john@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
		s := &student{
			Person: p,
		}
		assert.Nil(t, s.IsValid())
	})

	t.Run("Invalid Student - Nil Person", func(t *testing.T) {
		s := &student{
			Person: nil,
		}
		err := s.IsValid()
		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.PersonErrorMessage)
	})

	t.Run("Invalid Student - Invalid Person", func(t *testing.T) {
		p, _ := person.New(nil, "John", "invalid-email", "", "", "", "", "") // Invalid person
		s := &student{
			Person: p,
		}
		err := s.IsValid()
		assert.NotNil(t, err)
		// Should delegate to person validation
		assert.Contains(t, err.Messages(), "you need to provide a name with two words or more.")
	})

	t.Run("Valid Student with Registration", func(t *testing.T) {
		p, _ := person.New(nil, "John Doe", "john@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
		s := &student{
			Person:       p,
			registration: "2023101010",
		}
		assert.Nil(t, s.IsValid())
	})
}

func TestStudent_IntegrationWithPerson(t *testing.T) {
	t.Run("Student delegates validation to Person", func(t *testing.T) {
		p, _ := person.New(nil, "John Doe", "john@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
		s := &student{
			Person: p,
		}

		// Person is valid, so student should be valid
		assert.Nil(t, s.IsValid())

		// If we make person invalid, student should also be invalid
		invalidP, _ := person.New(nil, "John", "invalid", "", "", "", "", "")
		s.Person = invalidP
		err := s.IsValid()
		assert.NotNil(t, err)
	})

	t.Run("Student maintains Person relationship", func(t *testing.T) {
		p, _ := person.New(nil, "John Doe", "john@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
		s := &student{
			Person: p,
		}

		// Access Person methods through Student
		assert.Equal(t, "John Doe", s.Person.Name())
		assert.Equal(t, "john@example.com", s.Person.Email())
		assert.Equal(t, "11144477735", s.Person.CPF())
	})
}

func TestStudent_ComplexSetters(t *testing.T) {
	s := &student{}

	// Test SetInstitution
	instID := uuid.New()
	inst, _ := institution.NewBuilder().WithID(instID).WithName("Institution A").Build()
	err := s.SetInstitution(inst)
	assert.Nil(t, err)
	assert.Equal(t, inst, s.Institution())

	// Test SetCampus
	campusID := uuid.New()
	c, _ := campus.NewBuilder().WithID(campusID).WithName("Campus A").WithInstitutionID(instID).Build()
	err = s.SetCampus(c)
	assert.Nil(t, err)
	assert.Equal(t, c, s.Campus())

	// Test SetCourse
	courseID := uuid.New()
	course, _ := course.NewBuilder().WithID(courseID).WithName("Course A").Build()
	err = s.SetCourse(course)
	assert.Nil(t, err)
	assert.Equal(t, course, s.Course())

	// Test SetCurrentInternships
	internships := []internship.Internship{}
	err = s.SetCurrentInternships(internships)
	assert.Nil(t, err)
	assert.Equal(t, internships, s.CurrentInternships())

	// Test SetInternshipHistory
	history := []internship.Internship{}
	err = s.SetInternshipHistory(history)
	assert.Nil(t, err)
	assert.Equal(t, history, s.InternshipHistory())

	// Test SetFrequencyHistory
	frequencyHistory := []timeRecord.TimeRecord{}
	err = s.SetFrequencyHistory(frequencyHistory)
	assert.Nil(t, err)
	assert.Equal(t, frequencyHistory, s.FrequencyHistory())
}

func TestStudent_ComplexSetters_Errors(t *testing.T) {
	s := &student{}

	err := s.SetInstitution(nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.StudentInstitutionErrorMessage)

	err = s.SetCampus(nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.StudentCampusErrorMessage)

	err = s.SetCourse(nil)
	assert.NotNil(t, err)
	assert.Contains(t, err.Messages(), messages.StudentCourseErrorMessage)
}

func TestStudent_Getters(t *testing.T) {
	p, _ := person.New(nil, "John Doe", "john@example.com", "1990-01-01", "11144477735", "82999999999", "", "")
	instID := uuid.New()
	inst, _ := institution.NewBuilder().WithID(instID).WithName("Institution A").Build()
	campusID := uuid.New()
	c, _ := campus.NewBuilder().WithID(campusID).WithName("Campus A").WithInstitutionID(instID).Build()
	courseID := uuid.New()
	courseObj, _ := course.NewBuilder().WithID(courseID).WithName("Course A").Build()
	profilePic := "http://example.com/pic.jpg"
	teacherID := uuid.New()

	s := &student{
		Person:             p,
		name:               "John Doe",
		registration:       "2023101010",
		profilePicture:     &profilePic,
		_institution:       inst,
		_campus:            c,
		_course:            courseObj,
		totalWorkload:      100,
		workloadCompleted:  50,
		pendingWorkload:    50,
		responsibleTeacherID: teacherID,
		currentInternships: []internship.Internship{},
		internshipHistory:  []internship.Internship{},
		frequencyHistory:   []timeRecord.TimeRecord{},
	}

	assert.Equal(t, "2023101010", s.Registration())
	assert.Equal(t, &profilePic, s.ProfilePicture())
	assert.Equal(t, inst, s.Institution())
	assert.Equal(t, c, s.Campus())
	assert.Equal(t, courseObj, s.Course())
	assert.Equal(t, 100, s.TotalWorkload())
	assert.Equal(t, 50, s.WorkloadCompleted())
	assert.Equal(t, 50, s.PendingWorkload())
	assert.Equal(t, teacherID, s.ResponsibleTeacherID())
	assert.Equal(t, []internship.Internship{}, s.CurrentInternships())
	assert.Equal(t, []internship.Internship{}, s.InternshipHistory())
	assert.Equal(t, []timeRecord.TimeRecord{}, s.FrequencyHistory())
}
