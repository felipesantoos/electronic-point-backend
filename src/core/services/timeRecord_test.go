package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/domain/internship"
	"eletronic_point/src/core/domain/internshipLocation"
	"eletronic_point/src/core/domain/simplifiedStudent"
	"eletronic_point/src/core/domain/timeRecord"
	"eletronic_point/src/core/domain/timeRecordStatus"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/core/services/mocks"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestTimeRecordService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTimeRecordPort(ctrl)
	mockInternshipRepo := mocks.NewMockInternshipPort(ctrl)
	service := NewTimeRecordService(mockRepo, mockInternshipRepo)

	studentID := uuid.New()
	internshipID := uuid.New()
	date := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	entryTime := time.Date(2024, 1, 15, 8, 0, 0, 0, time.UTC)
	scheduleEntryTime := time.Date(2024, 1, 15, 8, 0, 0, 0, time.UTC)

	// Setup internship with valid student
	location, _ := internshipLocation.NewBuilder().
		WithID(uuid.New()).
		WithName("Location").
		WithNumber("123").
		WithStreet("Street").
		WithNeighborhood("Neighborhood").
		WithCity("City").
		WithZipCode("12345").
		Build()
	simplifiedStudent, _ := simplifiedStudent.NewBuilder().
		WithID(studentID).
		WithName("John Doe").
		Build()
	intern, _ := internship.NewBuilder().
		WithID(internshipID).
		WithStartedIn(time.Now()).
		WithLocation(location).
		WithStudent(simplifiedStudent).
		WithScheduleEntryTime(&scheduleEntryTime).
		Build()

	// Test successful create
	t.Run("Successful create", func(t *testing.T) {
		tr, _ := timeRecord.NewBuilder().
			WithID(uuid.New()).
			WithDate(date).
			WithEntryTime(entryTime).
			WithStudentID(studentID).
			WithInternshipID(internshipID).
			WithLocation("Office").
			Build()

		mockInternshipRepo.EXPECT().Get(internshipID).Return(intern, nil)
		mockRepo.EXPECT().List(gomock.Any()).Return([]timeRecord.TimeRecord{}, nil)
		id := uuid.New()
		mockRepo.EXPECT().Create(gomock.Any()).Return(&id, nil)

		createdID, err := service.Create(tr)
		assert.Nil(t, err)
		assert.NotNil(t, createdID)
	})

	// Test ownership validation failure
	t.Run("Ownership validation failure", func(t *testing.T) {
		differentStudentID := uuid.New()
		tr, _ := timeRecord.NewBuilder().
			WithID(uuid.New()).
			WithDate(date).
			WithEntryTime(entryTime).
			WithStudentID(differentStudentID).
			WithInternshipID(internshipID).
			WithLocation("Office").
			Build()

		mockInternshipRepo.EXPECT().Get(internshipID).Return(intern, nil)

		id, err := service.Create(tr)
		assert.NotNil(t, err)
		assert.Nil(t, id)
		assert.Contains(t, err.Messages(), messages.InternshipNotFoundErrorMessage)
	})

	// Test tolerance exceeded (more than 30 minutes)
	t.Run("Tolerance exceeded", func(t *testing.T) {
		lateEntryTime := scheduleEntryTime.Add(31 * time.Minute)
		tr, _ := timeRecord.NewBuilder().
			WithID(uuid.New()).
			WithDate(date).
			WithEntryTime(lateEntryTime).
			WithStudentID(studentID).
			WithInternshipID(internshipID).
			WithLocation("Office").
			Build()

		mockInternshipRepo.EXPECT().Get(internshipID).Return(intern, nil)

		id, err := service.Create(tr)
		assert.NotNil(t, err)
		assert.Nil(t, id)
		assert.Contains(t, err.Messages(), messages.TimeRecordToleranceErrorMessage)
	})

	// Test daily limit exceeded
	t.Run("Daily limit exceeded", func(t *testing.T) {
		tr, _ := timeRecord.NewBuilder().
			WithID(uuid.New()).
			WithDate(date).
			WithEntryTime(entryTime).
			WithStudentID(studentID).
			WithInternshipID(internshipID).
			WithLocation("Office").
			Build()

		// Create existing records that total 5 hours or more
		existingEntry := entryTime.Add(-5 * time.Hour)
		existingExit := entryTime // 5 hours exactly
		existingRecord, _ := timeRecord.NewBuilder().
			WithID(uuid.New()).
			WithDate(date).
			WithEntryTime(existingEntry).
			WithExitTime(&existingExit).
			WithStudentID(studentID).
			WithInternshipID(internshipID).
			WithLocation("Office").
			Build()

		mockInternshipRepo.EXPECT().Get(internshipID).Return(intern, nil)
		mockRepo.EXPECT().List(gomock.Any()).Return([]timeRecord.TimeRecord{
			existingRecord,
		}, nil)

		createdID, err := service.Create(tr)
		assert.NotNil(t, err)
		assert.Nil(t, createdID)
		assert.Contains(t, err.Messages(), messages.TimeRecordDailyLimitErrorMessage)
	})

	// Test automatic cutting of excess time
	t.Run("Automatic cutting of excess time", func(t *testing.T) {
		tr, _ := timeRecord.NewBuilder().
			WithID(uuid.New()).
			WithDate(date).
			WithEntryTime(entryTime).
			WithStudentID(studentID).
			WithInternshipID(internshipID).
			WithLocation("Office").
			Build()

		// Existing records total 4 hours (240 minutes)
		existingEntry := entryTime.Add(-5 * time.Hour)
		existingExit := entryTime.Add(-1 * time.Hour) // 4 hours
		existingRecord, _ := timeRecord.NewBuilder().
			WithID(uuid.New()).
			WithDate(date).
			WithEntryTime(existingEntry).
			WithExitTime(&existingExit).
			WithStudentID(studentID).
			WithInternshipID(internshipID).
			WithLocation("Office").
			Build()

		// New record tries to add 2 hours (would be 6h total)
		newExitTime := entryTime.Add(2 * time.Hour)
		tr.SetExitTime(&newExitTime)

		mockInternshipRepo.EXPECT().Get(internshipID).Return(intern, nil)
		mockRepo.EXPECT().List(gomock.Any()).Return([]timeRecord.TimeRecord{existingRecord}, nil)
		id := uuid.New()
		mockRepo.EXPECT().Create(gomock.Any()).DoAndReturn(func(tr timeRecord.TimeRecord) (*uuid.UUID, errors.Error) {
			// Verify exit time was cut to 1 hour (60 minutes allowed)
			exitTime := tr.ExitTime()
			expectedExitTime := entryTime.Add(1 * time.Hour)
			assert.NotNil(t, exitTime)
			assert.Equal(t, expectedExitTime.Hour(), exitTime.Hour())
			assert.Equal(t, expectedExitTime.Minute(), exitTime.Minute())
			return &id, nil
		})

		createdID, err := service.Create(tr)
		assert.Nil(t, err)
		assert.NotNil(t, createdID)
	})

	// Test error getting internship
	t.Run("Error getting internship", func(t *testing.T) {
		tr, _ := timeRecord.NewBuilder().
			WithID(uuid.New()).
			WithDate(date).
			WithEntryTime(entryTime).
			WithStudentID(studentID).
			WithInternshipID(internshipID).
			WithLocation("Office").
			Build()

		expectedErr := errors.NewUnexpected()
		mockInternshipRepo.EXPECT().Get(internshipID).Return(nil, expectedErr)

		id, err := service.Create(tr)
		assert.NotNil(t, err)
		assert.Nil(t, id)
		assert.Equal(t, expectedErr, err)
	})
}

func TestTimeRecordService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTimeRecordPort(ctrl)
	mockInternshipRepo := mocks.NewMockInternshipPort(ctrl)
	service := NewTimeRecordService(mockRepo, mockInternshipRepo)

	studentID := uuid.New()
	internshipID := uuid.New()
	recordID := uuid.New()
	date := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)
	entryTime := time.Date(2024, 1, 15, 8, 0, 0, 0, time.UTC)

	location, _ := internshipLocation.NewBuilder().
		WithID(uuid.New()).
		WithName("Location").
		WithNumber("123").
		WithStreet("Street").
		WithNeighborhood("Neighborhood").
		WithCity("City").
		WithZipCode("12345").
		Build()
	simplifiedStudent, _ := simplifiedStudent.NewBuilder().
		WithID(studentID).
		WithName("John Doe").
		Build()
	intern, _ := internship.NewBuilder().
		WithID(internshipID).
		WithStartedIn(time.Now()).
		WithLocation(location).
		WithStudent(simplifiedStudent).
		Build()

	// Test successful update
	t.Run("Successful update", func(t *testing.T) {
		tr, _ := timeRecord.NewBuilder().
			WithID(recordID).
			WithDate(date).
			WithEntryTime(entryTime).
			WithStudentID(studentID).
			WithInternshipID(internshipID).
			WithLocation("Office").
			Build()

		mockInternshipRepo.EXPECT().Get(internshipID).Return(intern, nil)
		mockRepo.EXPECT().List(gomock.Any()).Return([]timeRecord.TimeRecord{}, nil)
		mockRepo.EXPECT().Update(tr).Return(nil)

		err := service.Update(tr)
		assert.Nil(t, err)
	})

	// Test ownership validation failure
	t.Run("Ownership validation failure", func(t *testing.T) {
		differentStudentID := uuid.New()
		tr, _ := timeRecord.NewBuilder().
			WithID(recordID).
			WithDate(date).
			WithEntryTime(entryTime).
			WithStudentID(differentStudentID).
			WithInternshipID(internshipID).
			WithLocation("Office").
			Build()

		mockInternshipRepo.EXPECT().Get(internshipID).Return(intern, nil)

		err := service.Update(tr)
		assert.NotNil(t, err)
		assert.Contains(t, err.Messages(), messages.InternshipNotFoundErrorMessage)
	})
}

func TestTimeRecordService_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTimeRecordPort(ctrl)
	mockInternshipRepo := mocks.NewMockInternshipPort(ctrl)
	service := NewTimeRecordService(mockRepo, mockInternshipRepo)

	id := uuid.New()
	mockRepo.EXPECT().Delete(id).Return(nil)

	err := service.Delete(id)
	assert.Nil(t, err)
}

func TestTimeRecordService_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTimeRecordPort(ctrl)
	mockInternshipRepo := mocks.NewMockInternshipPort(ctrl)
	service := NewTimeRecordService(mockRepo, mockInternshipRepo)

	f := filters.TimeRecordFilters{}
	mockRepo.EXPECT().List(f).Return(nil, nil)

	_, err := service.List(f)
	assert.Nil(t, err)
}

func TestTimeRecordService_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTimeRecordPort(ctrl)
	mockInternshipRepo := mocks.NewMockInternshipPort(ctrl)
	service := NewTimeRecordService(mockRepo, mockInternshipRepo)

	id := uuid.New()
	f := filters.TimeRecordFilters{}
	expectedRecord, _ := timeRecord.NewBuilder().
		WithID(id).
		WithDate(time.Now()).
		WithEntryTime(time.Now()).
		WithStudentID(uuid.New()).
		WithInternshipID(uuid.New()).
		WithLocation("Office").
		Build()

	mockRepo.EXPECT().Get(id, f).Return(expectedRecord, nil)

	tr, err := service.Get(id, f)
	assert.Nil(t, err)
	assert.Equal(t, expectedRecord, tr)
}

func TestTimeRecordService_Approve(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTimeRecordPort(ctrl)
	mockInternshipRepo := mocks.NewMockInternshipPort(ctrl)
	service := NewTimeRecordService(mockRepo, mockInternshipRepo)

	timeRecordID := uuid.New()
	approvedBy := uuid.New()
	approvedStatusID := timeRecordStatus.Approved.ID()

	mockRepo.EXPECT().UpdateStatus(timeRecordID, approvedBy, approvedStatusID).Return(nil)

	err := service.Approve(timeRecordID, approvedBy)
	assert.Nil(t, err)
}

func TestTimeRecordService_Disapprove(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockTimeRecordPort(ctrl)
	mockInternshipRepo := mocks.NewMockInternshipPort(ctrl)
	service := NewTimeRecordService(mockRepo, mockInternshipRepo)

	timeRecordID := uuid.New()
	disapprovedBy := uuid.New()
	disapprovedStatusID := timeRecordStatus.Disapproved.ID()

	mockRepo.EXPECT().UpdateStatus(timeRecordID, disapprovedBy, disapprovedStatusID).Return(nil)

	err := service.Disapprove(timeRecordID, disapprovedBy)
	assert.Nil(t, err)
}
