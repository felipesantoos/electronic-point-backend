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

func TestInternshipHandler_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockInternshipPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewInternshipHandlers(mockService)

	mockCtx.EXPECT().QueryParam(gomock.Any()).Return("").AnyTimes()
	mockService.EXPECT().List(gomock.Any()).Return(nil, nil)
	mockCtx.EXPECT().JSON(http.StatusOK, gomock.Any()).Return(nil)

	err := handler.List(mockCtx)
	assert.Nil(t, err)
}

func TestInternshipHandler_Create_Admin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockInternshipPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewInternshipHandlers(mockService)

	studentID := uuid.New()
	teacherID := uuid.New()
	locationID := uuid.New()
	dto := request.Internship{
		StudentID:  studentID,
		TeacherID:  &teacherID,
		LocationID: locationID,
		StartedIn:  time.Now(),
	}

	mockCtx.EXPECT().RoleName().Return(role.ADMIN_ROLE_CODE).AnyTimes()
	mockCtx.EXPECT().Bind(gomock.Any()).DoAndReturn(func(i interface{}) error {
		*(i.(*request.Internship)) = dto
		return nil
	})
	id := uuid.New()
	mockService.EXPECT().Create(gomock.Any()).Return(&id, nil)
	mockCtx.EXPECT().JSON(http.StatusCreated, gomock.Any()).Return(nil)

	err := handler.Create(mockCtx)
	assert.Nil(t, err)
}

func TestInternshipHandler_Create_Teacher(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockInternshipPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewInternshipHandlers(mockService)

	studentID := uuid.New()
	locationID := uuid.New()
	dto := request.Internship{
		StudentID:  studentID,
		LocationID: locationID,
		StartedIn:  time.Now(),
	}

	mockCtx.EXPECT().RoleName().Return(role.TEACHER_ROLE_CODE).AnyTimes()
	mockCtx.EXPECT().Bind(gomock.Any()).DoAndReturn(func(i interface{}) error {
		*(i.(*request.Internship)) = dto
		return nil
	})
	id := uuid.New()
	mockService.EXPECT().Create(gomock.Any()).Return(&id, nil)
	mockCtx.EXPECT().JSON(http.StatusCreated, gomock.Any()).Return(nil)

	err := handler.Create(mockCtx)
	assert.Nil(t, err)
}
