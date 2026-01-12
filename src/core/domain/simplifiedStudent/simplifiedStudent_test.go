package simplifiedStudent

import (
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/institution"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestSimplifiedStudent_Setters(t *testing.T) {
	id := uuid.New()
	s := &simplifiedStudent{}

	err := s.SetID(&id)
	assert.Nil(t, err)
	assert.Equal(t, id, *s.ID())

	err = s.SetName("John Doe")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe", s.Name())

	url := "http://example.com/pic.jpg"
	err = s.SetProfilePicture(&url)
	assert.Nil(t, err)
	assert.Equal(t, &url, s.ProfilePicture())

	err = s.SetTotalWorkload(100)
	assert.Nil(t, err)
	assert.Equal(t, 100, s.TotalWorkload())

	// Test SetProfilePicture nil
	err = s.SetProfilePicture(nil)
	assert.Nil(t, err)
	assert.Nil(t, s.ProfilePicture())
}

func TestSimplifiedStudent_Setters_Errors(t *testing.T) {
	s := &simplifiedStudent{}

	err := s.SetID(&uuid.Nil)
	assert.NotNil(t, err)

	err = s.SetName("")
	assert.NotNil(t, err)

	err = s.SetInstitution(nil)
	assert.NotNil(t, err)

	err = s.SetCampus(nil)
	assert.NotNil(t, err)

	err = s.SetCourse(nil)
	assert.NotNil(t, err)
}

func TestSimplifiedStudent_ComplexSetters(t *testing.T) {
	s := &simplifiedStudent{}

	// Test SetInstitution with valid object
	instID := uuid.New()
	inst, _ := institution.NewBuilder().WithID(instID).WithName("Institution A").Build()
	err := s.SetInstitution(inst)
	assert.Nil(t, err)
	assert.Equal(t, inst, s.Institution())

	// Test SetCampus with valid object
	campusID := uuid.New()
	c, _ := campus.NewBuilder().WithID(campusID).WithName("Campus A").WithInstitutionID(instID).Build()
	err = s.SetCampus(c)
	assert.Nil(t, err)
	assert.Equal(t, c, s.Campus())

	// Test SetCourse with valid object
	courseID := uuid.New()
	course, _ := course.NewBuilder().WithID(courseID).WithName("Course A").Build()
	err = s.SetCourse(course)
	assert.Nil(t, err)
	assert.Equal(t, course, s.Course())
}

func TestSimplifiedStudent_Getters(t *testing.T) {
	id := uuid.New()
	profilePic := "http://example.com/pic.jpg"
	instID := uuid.New()
	inst, _ := institution.NewBuilder().WithID(instID).WithName("Institution A").Build()
	campusID := uuid.New()
	c, _ := campus.NewBuilder().WithID(campusID).WithName("Campus A").WithInstitutionID(instID).Build()
	courseID := uuid.New()
	courseObj, _ := course.NewBuilder().WithID(courseID).WithName("Course A").Build()

	s := &simplifiedStudent{
		id:             &id,
		name:           "John Doe",
		profilePicture: &profilePic,
		_institution:   inst,
		_campus:        c,
		_course:        courseObj,
		totalWorkload:  100,
	}

	assert.Equal(t, &id, s.ID())
	assert.Equal(t, "John Doe", s.Name())
	assert.Equal(t, &profilePic, s.ProfilePicture())
	assert.Equal(t, inst, s.Institution())
	assert.Equal(t, c, s.Campus())
	assert.Equal(t, courseObj, s.Course())
	assert.Equal(t, 100, s.TotalWorkload())
}
