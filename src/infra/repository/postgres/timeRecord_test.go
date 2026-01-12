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
	studentDomain "eletronic_point/src/core/domain/student"
	timeRecordDomain "eletronic_point/src/core/domain/timeRecord"
	"eletronic_point/src/core/services/filters"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestTimeRecordRepository_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	SetupTestDB(t)

	repo := NewTimeRecordRepository()
	locationRepo := NewInternshipLocationRepository()
	internshipRepo := NewInternshipRepository()
	studentRepo := NewStudentRepository()
	instRepo := NewInstitutionRepository()
	campusRepo := NewCampusRepository()
	courseRepo := NewCourseRepository()
	statusRepo := NewTimeRecordStatusRepository()
	accRepo := NewAccountRepository()

	// Helper function to create a student
	createStudent := func(t *testing.T) (*uuid.UUID, *uuid.UUID) {
		// Create a teacher first
		tID := uuid.New()
		tpID := uuid.New()
		tp, _ := person.New(&tpID, "Teacher Test", "teacher@example.com", "1980-01-01", "73595867041", "82988888888", "", "")
		tr, _ := role.New(nil, "Teacher", role.TEACHER_ROLE_CODE)
		tAcc, _ := account.New(&tID, "teacher@example.com", "pass123", tr, tp, nil, nil)
		_, err := accRepo.Create(tAcc)
		assert.Nil(t, err)

		// Get created teacher to get real person ID
		search := "teacher@example.com"
		teachers, err := accRepo.List(filters.AccountFilters{Search: &search})
		assert.Nil(t, err)
		assert.NotEmpty(t, teachers)
		realTeacherID := *teachers[0].Person().ID()

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

		student, _ := studentDomain.NewBuilder().
			WithPerson(p).
			WithRegistration("12345").
			WithInstitution(instFromRepo).
			WithCampus(campusFromRepo).
			WithCourse(courseFromRepo).
			WithTotalWorkload(100).
			WithResponsibleTeacherID(realTeacherID).
			Build()

		createdStudentID, err := studentRepo.Create(student)
		assert.Nil(t, err)
		return createdStudentID, &realTeacherID
	}

	// Helper function to create an internship
	createInternship := func(t *testing.T, studentID *uuid.UUID) *uuid.UUID {
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

		createdID, err := internshipRepo.Create(intern)
		assert.Nil(t, err)
		return createdID
	}

	t.Run("Create and Get", func(t *testing.T) {
		CleanDB(t)
		studentID, _ := createStudent(t)
		internshipID := createInternship(t, studentID)

		// Create time record
		date := time.Now()
		entryTime := time.Date(date.Year(), date.Month(), date.Day(), 8, 0, 0, 0, time.UTC)
		tr, _ := timeRecordDomain.NewBuilder().
			WithDate(date).
			WithEntryTime(entryTime).
			WithLocation("Test Location").
			WithStudentID(*studentID).
			WithInternshipID(*internshipID).
			Build()

		createdID, err := repo.Create(tr)
		assert.Nil(t, err)
		assert.NotNil(t, createdID)

		// Get time record
		found, err := repo.Get(*createdID, filters.TimeRecordFilters{StudentID: studentID})
		assert.Nil(t, err)
		assert.Equal(t, date.Format("2006-01-02"), found.Date().Format("2006-01-02"))
		assert.Equal(t, entryTime.Format("15:04:05"), found.EntryTime().Format("15:04:05"))
	})

	t.Run("List", func(t *testing.T) {
		CleanDB(t)
		studentID, _ := createStudent(t)
		internshipID := createInternship(t, studentID)

		// Create time record
		date := time.Now()
		entryTime := time.Date(date.Year(), date.Month(), date.Day(), 9, 0, 0, 0, time.UTC)
		tr, _ := timeRecordDomain.NewBuilder().
			WithDate(date).
			WithEntryTime(entryTime).
			WithLocation("Test Location 2").
			WithStudentID(*studentID).
			WithInternshipID(*internshipID).
			Build()

		_, err := repo.Create(tr)
		assert.Nil(t, err)

		// List time records
		records, err := repo.List(filters.TimeRecordFilters{StudentID: studentID})
		assert.Nil(t, err)
		assert.NotEmpty(t, records)
	})

	t.Run("Update", func(t *testing.T) {
		CleanDB(t)
		studentID, _ := createStudent(t)
		internshipID := createInternship(t, studentID)

		// Create time record
		date := time.Now()
		entryTime := time.Date(date.Year(), date.Month(), date.Day(), 8, 0, 0, 0, time.UTC)
		tr, _ := timeRecordDomain.NewBuilder().
			WithDate(date).
			WithEntryTime(entryTime).
			WithLocation("Test Location 3").
			WithStudentID(*studentID).
			WithInternshipID(*internshipID).
			Build()

		createdID, err := repo.Create(tr)
		assert.Nil(t, err)

		// Update time record
		updated, err := repo.Get(*createdID, filters.TimeRecordFilters{StudentID: studentID})
		assert.Nil(t, err)

		newEntryTime := time.Date(date.Year(), date.Month(), date.Day(), 9, 30, 0, 0, time.UTC)
		exitTime := time.Date(date.Year(), date.Month(), date.Day(), 18, 0, 0, 0, time.UTC)
		updatedTR, _ := timeRecordDomain.NewBuilder().
			WithID(updated.ID()).
			WithDate(updated.Date()).
			WithEntryTime(newEntryTime).
			WithExitTime(&exitTime).
			WithLocation("Updated Location").
			WithStudentID(updated.StudentID()).
			WithInternshipID(updated.InternshipID()).
			Build()

		err = repo.Update(updatedTR)
		assert.Nil(t, err)

		// Verify update
		found, err := repo.Get(*createdID, filters.TimeRecordFilters{StudentID: studentID})
		assert.Nil(t, err)
		assert.Equal(t, newEntryTime.Format("15:04:05"), found.EntryTime().Format("15:04:05"))
		assert.NotNil(t, found.ExitTime())
	})

	t.Run("Delete", func(t *testing.T) {
		CleanDB(t)
		studentID, _ := createStudent(t)
		internshipID := createInternship(t, studentID)

		// Create time record
		date := time.Now()
		entryTime := time.Date(date.Year(), date.Month(), date.Day(), 8, 0, 0, 0, time.UTC)
		tr, _ := timeRecordDomain.NewBuilder().
			WithDate(date).
			WithEntryTime(entryTime).
			WithLocation("Test Location 4").
			WithStudentID(*studentID).
			WithInternshipID(*internshipID).
			Build()

		createdID, err := repo.Create(tr)
		assert.Nil(t, err)

		// Delete time record
		err = repo.Delete(*createdID)
		assert.Nil(t, err)

		// Verify deletion
		_, err = repo.Get(*createdID, filters.TimeRecordFilters{StudentID: studentID})
		assert.NotNil(t, err)
	})

	t.Run("UpdateStatus", func(t *testing.T) {
		CleanDB(t)
		studentID, teacherID := createStudent(t)
		internshipID := createInternship(t, studentID)

		// Create time record
		date := time.Now()
		entryTime := time.Date(date.Year(), date.Month(), date.Day(), 8, 0, 0, 0, time.UTC)
		tr, _ := timeRecordDomain.NewBuilder().
			WithDate(date).
			WithEntryTime(entryTime).
			WithLocation("Test Location 5").
			WithStudentID(*studentID).
			WithInternshipID(*internshipID).
			Build()

		createdID, err := repo.Create(tr)
		assert.Nil(t, err)

		// Get approved status
		statuses, err := statusRepo.List()
		assert.Nil(t, err)
		var approvedStatusID *uuid.UUID
		for _, s := range statuses {
			if s.Name() == "approved" {
				id := s.ID()
				approvedStatusID = &id
				break
			}
		}
		if approvedStatusID == nil {
			t.Skip("approved status not found")
		}

		// Update status using the same teacher
		err = repo.UpdateStatus(*createdID, *teacherID, *approvedStatusID)
		assert.Nil(t, err)

		// Verify status update (by getting the record)
		found, err := repo.Get(*createdID, filters.TimeRecordFilters{StudentID: studentID})
		assert.Nil(t, err)
		assert.NotNil(t, found)
	})
}
