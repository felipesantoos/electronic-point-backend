package postgres

import (
	"eletronic_point/src/core/domain/account"
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/domain/person"
	"eletronic_point/src/core/domain/role"
	studentDomain "eletronic_point/src/core/domain/student"
	"eletronic_point/src/core/services/filters"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestStudentRepository_Integration(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	SetupTestDB(t)

	repo := NewStudentRepository()
	instRepo := NewInstitutionRepository()
	campusRepo := NewCampusRepository()
	courseRepo := NewCourseRepository()
	accRepo := NewAccountRepository()

	// Helper function to create dependencies
	createDependencies := func(t *testing.T) (*uuid.UUID, *uuid.UUID, *uuid.UUID, uuid.UUID) {
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

		// Create course
		courseID := uuid.New()
		courseObj, _ := course.NewBuilder().WithID(courseID).WithName("Test Course").Build()
		createdCourseID, err := courseRepo.Create(courseObj)
		assert.Nil(t, err)

		return createdInstID, createdCampusID, createdCourseID, realTeacherID
	}

	t.Run("Create and Get", func(t *testing.T) {
		CleanDB(t)
		instID, campusID, courseID, teacherID := createDependencies(t)

		// Get entities from repo
		instFromRepo, err := instRepo.Get(*instID)
		assert.Nil(t, err)
		campusFromRepo, err := campusRepo.Get(*campusID)
		assert.Nil(t, err)
		courseFromRepo, err := courseRepo.Get(*courseID)
		assert.Nil(t, err)

		// Create student
		p, _ := person.NewBuilder().
			WithName("Student Test").
			WithEmail("student1@example.com").
			WithBirthDate("1990-01-01").
			WithCPF("11144477735").
			WithPhone("82999999999").
			Build()
		
		student, _ := studentDomain.NewBuilder().
			WithPerson(p).
			WithRegistration("REG001").
			WithInstitution(instFromRepo).
			WithCampus(campusFromRepo).
			WithCourse(courseFromRepo).
			WithTotalWorkload(100).
			WithResponsibleTeacherID(teacherID).
			Build()

		createdID, err := repo.Create(student)
		assert.Nil(t, err)
		assert.NotNil(t, createdID)

		// Get student
		found, err := repo.Get(*createdID, filters.StudentFilters{})
		assert.Nil(t, err)
		assert.Equal(t, "Student Test", found.Name())
		assert.Equal(t, "REG001", found.Registration())
	})

	t.Run("List", func(t *testing.T) {
		CleanDB(t)
		instID, campusID, courseID, teacherID := createDependencies(t)

		// Get entities from repo
		instFromRepo, err := instRepo.Get(*instID)
		assert.Nil(t, err)
		campusFromRepo, err := campusRepo.Get(*campusID)
		assert.Nil(t, err)
		courseFromRepo, err := courseRepo.Get(*courseID)
		assert.Nil(t, err)

		// Create student
		p, _ := person.NewBuilder().
			WithName("Student Test 2").
			WithEmail("student2@example.com").
			WithBirthDate("1990-01-01").
			WithCPF("11144477735").
			WithPhone("82999999999").
			Build()
		
		student, _ := studentDomain.NewBuilder().
			WithPerson(p).
			WithRegistration("REG002").
			WithInstitution(instFromRepo).
			WithCampus(campusFromRepo).
			WithCourse(courseFromRepo).
			WithTotalWorkload(100).
			WithResponsibleTeacherID(teacherID).
			Build()

		_, err = repo.Create(student)
		assert.Nil(t, err)

		// List students
		students, err := repo.List(filters.StudentFilters{})
		assert.Nil(t, err)
		assert.NotEmpty(t, students)
	})

	t.Run("Update", func(t *testing.T) {
		CleanDB(t)
		instID, campusID, courseID, teacherID := createDependencies(t)

		// Get entities from repo
		instFromRepo, err := instRepo.Get(*instID)
		assert.Nil(t, err)
		campusFromRepo, err := campusRepo.Get(*campusID)
		assert.Nil(t, err)
		courseFromRepo, err := courseRepo.Get(*courseID)
		assert.Nil(t, err)

		// Create student
		p, _ := person.NewBuilder().
			WithName("Student Test 3").
			WithEmail("student3@example.com").
			WithBirthDate("1990-01-01").
			WithCPF("11144477735").
			WithPhone("82999999999").
			Build()
		
		student, _ := studentDomain.NewBuilder().
			WithPerson(p).
			WithRegistration("REG003").
			WithInstitution(instFromRepo).
			WithCampus(campusFromRepo).
			WithCourse(courseFromRepo).
			WithTotalWorkload(100).
			WithResponsibleTeacherID(teacherID).
			Build()

		createdID, err := repo.Create(student)
		assert.Nil(t, err)

		// Update student
		updated, err := repo.Get(*createdID, filters.StudentFilters{})
		assert.Nil(t, err)

		updatedPerson, _ := person.NewBuilder().
			WithID(*updated.ID()).
			WithName("Updated Student").
			WithEmail(updated.Email()).
			WithBirthDate(updated.BirthDate()).
			WithCPF(updated.CPF()).
			WithPhone("82888888888").
			Build()

		updatedStudent, _ := studentDomain.NewBuilder().
			WithPerson(updatedPerson).
			WithRegistration("REG003UPD").
			WithInstitution(updated.Institution()).
			WithCampus(updated.Campus()).
			WithCourse(updated.Course()).
			WithTotalWorkload(150).
			WithResponsibleTeacherID(teacherID).
			Build()

		err = repo.Update(updatedStudent)
		assert.Nil(t, err)

		// Verify update
		found, err := repo.Get(*createdID, filters.StudentFilters{})
		assert.Nil(t, err)
		assert.Equal(t, "Updated Student", found.Name())
		assert.Equal(t, "REG003UPD", found.Registration())
		assert.Equal(t, 150, found.TotalWorkload())
	})

	t.Run("Delete", func(t *testing.T) {
		CleanDB(t)
		instID, campusID, courseID, teacherID := createDependencies(t)

		// Get entities from repo
		instFromRepo, err := instRepo.Get(*instID)
		assert.Nil(t, err)
		campusFromRepo, err := campusRepo.Get(*campusID)
		assert.Nil(t, err)
		courseFromRepo, err := courseRepo.Get(*courseID)
		assert.Nil(t, err)

		// Create student
		p, _ := person.NewBuilder().
			WithName("Student Test 4").
			WithEmail("student4@example.com").
			WithBirthDate("1990-01-01").
			WithCPF("11144477735").
			WithPhone("82999999999").
			Build()
		
		student, _ := studentDomain.NewBuilder().
			WithPerson(p).
			WithRegistration("REG004").
			WithInstitution(instFromRepo).
			WithCampus(campusFromRepo).
			WithCourse(courseFromRepo).
			WithTotalWorkload(100).
			WithResponsibleTeacherID(teacherID).
			Build()

		createdID, err := repo.Create(student)
		assert.Nil(t, err)

		// Delete student
		err = repo.Delete(*createdID)
		assert.Nil(t, err)

		// Verify deletion
		_, err = repo.Get(*createdID, filters.StudentFilters{})
		assert.NotNil(t, err)
	})
}
