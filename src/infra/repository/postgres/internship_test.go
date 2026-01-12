package postgres

import (
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/core/domain/simplifiedStudent"
	"eletronic_point/src/core/domain/student"
	"eletronic_point/src/core/services/filters"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestInternshipRepository_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	SetupTestDB(t)
	CleanDB(t)

	repo := NewInternshipRepository()
	locationRepo := NewInternshipLocationRepository()
	studentRepo := NewStudentRepository()
	instRepo := NewInstitutionRepository()
	campusRepo := NewCampusRepository()
	courseRepo := NewCourseRepository()

	// Helper function to create a student
	createStudent := func(t *testing.T) *uuid.UUID {
		// Create a teacher first
		tID := uuid.New()
		tpID := uuid.New()
		tp, _ := person.New(&tpID, "Teacher Test", "teacher@example.com", "1980-01-01", "73595867041", "82988888888", "", "")
		tr, _ := role.New(nil, "Teacher", role.TEACHER_ROLE_CODE)
		tAcc, _ := account.New(&tID, "teacher@example.com", "pass123", tr, tp, nil, nil)
		accRepo := NewAccountRepository()
		_, err := accRepo.Create(tAcc)
		assert.Nil(t, err)

		// Get created teacher to get real person ID
		search := "teacher@example.com"
		teachers, err := accRepo.List(filters.AccountFilters{Search: &search})
		assert.Nil(t, err)
		assert.NotEmpty(t, teachers)
		realTeacherID := teachers[0].Person().ID()

		// Create institution
		instID := uuid.New()
		inst, _ := institution.NewBuilder().WithID(instID).WithName("Test Inst").Build()
		createdInstID, err := instRepo.Create(inst)
		assert.Nil(t, err)

		// Create campus
		campusID := uuid.New()
		c, _ := campus.NewBuilder().WithID(campusID).WithName("Test Campus").WithInstitutionID(*createdInstID).Build()
		createdCampusID, err := campusRepo.Create(c)
		assert.Nil(t, err)

		// Get campus and institution from repo
		campusFromRepo, err := campusRepo.Get(*createdCampusID)
		assert.Nil(t, err)
		instFromRepo, err := instRepo.Get(*createdInstID)
		assert.Nil(t, err)

		// Create course
		courseID := uuid.New()
		courseObj, _ := course.NewBuilder().WithID(courseID).WithName("Test Course").Build()
		createdCourseID, err := courseRepo.Create(courseObj)
		assert.Nil(t, err)

		// Get course from repo
		courseFromRepo, err := courseRepo.Get(*createdCourseID)
		assert.Nil(t, err)

		// Create student
		p, _ := person.NewBuilder().
			WithName("Student Test").
			WithEmail("student@example.com").
			WithBirthDate("1990-01-01").
			WithCPF("11144477735").
			WithPhone("82999999999").
			Build()
		
		student, _ := student.NewBuilder().
			WithPerson(p).
			WithRegistration("12345").
			WithInstitution(instFromRepo).
			WithCampus(campusFromRepo).
			WithCourse(courseFromRepo).
			WithTotalWorkload(100).
			WithResponsibleTeacherID(*realTeacherID).
			Build()

		createdStudentID, err := studentRepo.Create(student)
		assert.Nil(t, err)
		return createdStudentID
	}

	t.Run("Create and Get", func(t *testing.T) {
		CleanDB(t)
		studentID := createStudent(t)

		// Create location
		locationID := uuid.New()
		location, _ := internshipLocation.NewBuilder().
			WithID(locationID).
			WithName("Test Location").
			WithNumber("123").
			WithStreet("Test Street").
			WithNeighborhood("Test Neighborhood").
			WithCity("Test City").
			WithZipCode("12345").
			Build()
		
		createdLocationID, err := locationRepo.Create(location)
		assert.Nil(t, err)

		// Get student to create SimplifiedStudent
		student, err := studentRepo.Get(*studentID, filters.StudentFilters{})
		assert.Nil(t, err)

		// Create SimplifiedStudent
		simplifiedStudent, _ := simplifiedStudent.NewBuilder().
			WithID(*student.ID()).
			WithName(student.Name()).
			WithInstitution(student.Institution()).
			WithCampus(student.Campus()).
			WithCourse(student.Course()).
			WithTotalWorkload(student.TotalWorkload()).
			Build()

		// Get location from repo
		locationFromRepo, err := locationRepo.Get(*createdLocationID)
		assert.Nil(t, err)

		// Create internship
		internshipID := uuid.New()
		startedIn := time.Now()
		intern, _ := internship.NewBuilder().
			WithID(internshipID).
			WithStartedIn(startedIn).
			WithLocation(locationFromRepo).
			WithStudent(simplifiedStudent).
			Build()

		createdID, err := repo.Create(intern)
		assert.Nil(t, err)
		assert.NotNil(t, createdID)

		// Get internship
		found, err := repo.Get(*createdID)
		assert.Nil(t, err)
		assert.Equal(t, startedIn.Format("2006-01-02"), found.StartedIn().Format("2006-01-02"))
		assert.NotNil(t, found.Location())
		assert.NotNil(t, found.Student())
	})

	t.Run("List", func(t *testing.T) {
		CleanDB(t)
		studentID := createStudent(t)

		// Create location
		locationID := uuid.New()
		location, _ := internshipLocation.NewBuilder().
			WithID(locationID).
			WithName("Test Location 2").
			WithNumber("456").
			WithStreet("Test Street 2").
			WithNeighborhood("Test Neighborhood 2").
			WithCity("Test City 2").
			WithZipCode("54321").
			Build()
		
		createdLocationID, err := locationRepo.Create(location)
		assert.Nil(t, err)

		// Get student
		student, err := studentRepo.Get(*studentID, filters.StudentFilters{})
		assert.Nil(t, err)

		simplifiedStudent, _ := simplifiedStudent.NewBuilder().
			WithID(*student.ID()).
			WithName(student.Name()).
			WithInstitution(student.Institution()).
			WithCampus(student.Campus()).
			WithCourse(student.Course()).
			WithTotalWorkload(student.TotalWorkload()).
			Build()

		locationFromRepo, err := locationRepo.Get(*createdLocationID)
		assert.Nil(t, err)

		// Create internship
		internshipID := uuid.New()
		startedIn := time.Now()
		intern, _ := internship.NewBuilder().
			WithID(internshipID).
			WithStartedIn(startedIn).
			WithLocation(locationFromRepo).
			WithStudent(simplifiedStudent).
			Build()

		_, err = repo.Create(intern)
		assert.Nil(t, err)

		// List internships
		internships, err := repo.List(filters.InternshipFilters{StudentID: studentID})
		assert.Nil(t, err)
		assert.NotEmpty(t, internships)
	})

	t.Run("Update", func(t *testing.T) {
		CleanDB(t)
		studentID := createStudent(t)

		// Create location
		locationID := uuid.New()
		location, _ := internshipLocation.NewBuilder().
			WithID(locationID).
			WithName("Test Location 3").
			WithNumber("789").
			WithStreet("Test Street 3").
			WithNeighborhood("Test Neighborhood 3").
			WithCity("Test City 3").
			WithZipCode("98765").
			Build()
		
		createdLocationID, err := locationRepo.Create(location)
		assert.Nil(t, err)

		// Get student
		student, err := studentRepo.Get(*studentID, filters.StudentFilters{})
		assert.Nil(t, err)

		simplifiedStudent, _ := simplifiedStudent.NewBuilder().
			WithID(*student.ID()).
			WithName(student.Name()).
			WithInstitution(student.Institution()).
			WithCampus(student.Campus()).
			WithCourse(student.Course()).
			WithTotalWorkload(student.TotalWorkload()).
			Build()

		locationFromRepo, err := locationRepo.Get(*createdLocationID)
		assert.Nil(t, err)

		// Create internship
		internshipID := uuid.New()
		startedIn := time.Now()
		intern, _ := internship.NewBuilder().
			WithID(internshipID).
			WithStartedIn(startedIn).
			WithLocation(locationFromRepo).
			WithStudent(simplifiedStudent).
			Build()

		createdID, err := repo.Create(intern)
		assert.Nil(t, err)

		// Update internship
		updated, err := repo.Get(*createdID)
		assert.Nil(t, err)

		newStartedIn := time.Now().Add(24 * time.Hour)
		updatedIntern, _ := internship.NewBuilder().
			WithID(updated.ID()).
			WithStartedIn(newStartedIn).
			WithLocation(updated.Location()).
			WithStudent(updated.Student()).
			Build()

		err = repo.Update(updatedIntern)
		assert.Nil(t, err)

		// Verify update
		found, err := repo.Get(*createdID)
		assert.Nil(t, err)
		assert.Equal(t, newStartedIn.Format("2006-01-02"), found.StartedIn().Format("2006-01-02"))
	})

	t.Run("Delete", func(t *testing.T) {
		CleanDB(t)
		studentID := createStudent(t)

		// Create location
		locationID := uuid.New()
		location, _ := internshipLocation.NewBuilder().
			WithID(locationID).
			WithName("Test Location 4").
			WithNumber("000").
			WithStreet("Test Street 4").
			WithNeighborhood("Test Neighborhood 4").
			WithCity("Test City 4").
			WithZipCode("00000").
			Build()
		
		createdLocationID, err := locationRepo.Create(location)
		assert.Nil(t, err)

		// Get student
		student, err := studentRepo.Get(*studentID, filters.StudentFilters{})
		assert.Nil(t, err)

		simplifiedStudent, _ := simplifiedStudent.NewBuilder().
			WithID(*student.ID()).
			WithName(student.Name()).
			WithInstitution(student.Institution()).
			WithCampus(student.Campus()).
			WithCourse(student.Course()).
			WithTotalWorkload(student.TotalWorkload()).
			Build()

		locationFromRepo, err := locationRepo.Get(*createdLocationID)
		assert.Nil(t, err)

		// Create internship
		internshipID := uuid.New()
		startedIn := time.Now()
		intern, _ := internship.NewBuilder().
			WithID(internshipID).
			WithStartedIn(startedIn).
			WithLocation(locationFromRepo).
			WithStudent(simplifiedStudent).
			Build()

		createdID, err := repo.Create(intern)
		assert.Nil(t, err)

		// Delete internship
		err = repo.Delete(*createdID)
		assert.Nil(t, err)

		// Verify deletion
		_, err = repo.Get(*createdID)
		assert.NotNil(t, err)
	})
}
