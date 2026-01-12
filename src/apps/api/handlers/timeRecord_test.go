package handlers

import (
	"eletronic_point/src/apps/api/handlers/dto/request"
	"eletronic_point/src/apps/api/handlers/mocks"
	"eletronic_point/src/core/domain/role"
	"net/http"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestTimeRecordHandler_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockTimeRecordPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewTimeRecordHandlers(mockService)

	mockCtx.EXPECT().QueryParam(gomock.Any()).Return("").AnyTimes()
	mockCtx.EXPECT().RoleName().Return(role.ADMIN_ROLE_CODE).AnyTimes()
	mockService.EXPECT().List(gomock.Any()).Return(nil, nil)
	mockCtx.EXPECT().JSON(http.StatusOK, gomock.Any()).Return(nil)

	err := handler.List(mockCtx)
	assert.Nil(t, err)
}

func TestTimeRecordHandler_Create_Admin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockTimeRecordPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewTimeRecordHandlers(mockService)

	studentID := uuid.New()
	internshipID := uuid.New()
	dto := request.TimeRecord{
		InternshipID: internshipID,
		StudentID:    &studentID,
		Date:         time.Now(),
		EntryTime:    time.Now(),
		Location:     "Test",
	}

	mockCtx.EXPECT().RoleName().Return(role.ADMIN_ROLE_CODE).AnyTimes()
	mockCtx.EXPECT().Bind(gomock.Any()).DoAndReturn(func(i interface{}) error {
		*(i.(*request.TimeRecord)) = dto
		return nil
	})
	id := uuid.New()
	mockService.EXPECT().Create(gomock.Any()).Return(&id, nil)
	mockCtx.EXPECT().JSON(http.StatusCreated, gomock.Any()).Return(nil)

	err := handler.Create(mockCtx)
	assert.Nil(t, err)
}

func TestTimeRecordHandler_Create_Student(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockTimeRecordPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewTimeRecordHandlers(mockService)

	studentID := uuid.New()
	internshipID := uuid.New()
	dto := request.TimeRecord{
		InternshipID: internshipID,
		Date:         time.Now(),
		EntryTime:    time.Now(),
		Location:     "Test",
	}

	mockCtx.EXPECT().RoleName().Return(role.STUDENT_ROLE_CODE).AnyTimes()
	mockCtx.EXPECT().ProfileID().Return(&studentID).AnyTimes()
	mockCtx.EXPECT().Bind(gomock.Any()).DoAndReturn(func(i interface{}) error {
		*(i.(*request.TimeRecord)) = dto
		return nil
	})
	id := uuid.New()
	mockService.EXPECT().Create(gomock.Any()).Return(&id, nil)
	mockCtx.EXPECT().JSON(http.StatusCreated, gomock.Any()).Return(nil)

	err := handler.Create(mockCtx)
	assert.Nil(t, err)
}

func TestTimeRecordHandler_Approve_Admin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockTimeRecordPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewTimeRecordHandlers(mockService)

	id := uuid.New()
	teacherID := uuid.New()

	mockCtx.EXPECT().Param("id").Return(id.String())
	mockCtx.EXPECT().RoleName().Return(role.ADMIN_ROLE_CODE).AnyTimes()
	mockCtx.EXPECT().Request().Return(&http.Request{}).AnyTimes()
	mockCtx.EXPECT().Bind(gomock.Any()).DoAndReturn(func(i interface{}) error {
		*(i.(*request.TimeRecord)) = request.TimeRecord{TeacherID: &teacherID}
		return nil
	})
	mockService.EXPECT().Approve(id, teacherID).Return(nil)
	mockCtx.EXPECT().NoContent(http.StatusNoContent).Return(nil)

	err := handler.Approve(mockCtx)
	assert.Nil(t, err)
}

func TestTimeRecordHandler_Approve_Teacher(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockTimeRecordPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewTimeRecordHandlers(mockService)

	id := uuid.New()
	teacherID := uuid.New()

	mockCtx.EXPECT().Param("id").Return(id.String())
	mockCtx.EXPECT().RoleName().Return(role.TEACHER_ROLE_CODE).AnyTimes()
	mockCtx.EXPECT().Request().Return(&http.Request{}).AnyTimes()
	mockCtx.EXPECT().ProfileID().Return(&teacherID).AnyTimes()

	mockService.EXPECT().Approve(id, teacherID).Return(nil)
	mockCtx.EXPECT().NoContent(http.StatusNoContent).Return(nil)

	err := handler.Approve(mockCtx)
	assert.Nil(t, err)
}
