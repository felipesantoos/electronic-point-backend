package handlers

import (
	"bytes"
	"eletronic_point/src/apps/api/handlers/mocks"
	"eletronic_point/src/core/domain/role"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestStudentHandler_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockStudentPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewStudentHandlers(mockService)

	mockCtx.EXPECT().QueryParam(gomock.Any()).Return("").AnyTimes()
	mockCtx.EXPECT().RoleName().Return(role.ADMIN_ROLE_CODE).AnyTimes()
	mockService.EXPECT().List(gomock.Any()).Return(nil, nil)
	mockCtx.EXPECT().JSON(http.StatusOK, gomock.Any()).Return(nil)

	err := handler.List(mockCtx)
	assert.Nil(t, err)
}

func TestStudentHandler_Create_Admin(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockStudentPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewStudentHandlers(mockService)

	teacherID := uuid.New()
	campusID := uuid.New()
	courseID := uuid.New()

	// Setup mock context for multipart form data
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()
	req := httptest.NewRequest(http.MethodPost, "/students", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	
	mockCtx.EXPECT().RoleName().Return(role.ADMIN_ROLE_CODE).AnyTimes()
	mockCtx.EXPECT().Request().Return(req).AnyTimes()
	mockCtx.EXPECT().FormValue("name").Return("Test Student")
	mockCtx.EXPECT().FormValue("birth_date").Return("2000-01-01")
	mockCtx.EXPECT().FormValue("cpf").Return("11144477735")
	mockCtx.EXPECT().FormValue("email").Return("test@test.com")
	mockCtx.EXPECT().FormValue("phone").Return("123456789")
	mockCtx.EXPECT().FormValue("registration").Return("20211234")
	mockCtx.EXPECT().FormValue("campus_id").Return(campusID.String())
	mockCtx.EXPECT().FormValue("course_id").Return(courseID.String())
	mockCtx.EXPECT().FormValue("total_workload").Return("100")
	mockCtx.EXPECT().FormValue("responsible_teacher_id").Return(teacherID.String())

	id := uuid.New()
	mockService.EXPECT().Create(gomock.Any()).Return(&id, nil)
	mockCtx.EXPECT().JSON(http.StatusCreated, gomock.Any()).Return(nil)

	err := handler.Create(mockCtx)
	assert.Nil(t, err)
}

func TestStudentHandler_Create_Teacher(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockStudentPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewStudentHandlers(mockService)

	teacherID := uuid.New()
	campusID := uuid.New()
	courseID := uuid.New()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	writer.Close()
	req := httptest.NewRequest(http.MethodPost, "/students", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	
	mockCtx.EXPECT().RoleName().Return(role.TEACHER_ROLE_CODE).AnyTimes()
	mockCtx.EXPECT().ProfileID().Return(&teacherID).AnyTimes()
	mockCtx.EXPECT().Request().Return(req).AnyTimes()
	mockCtx.EXPECT().FormValue("name").Return("Test Student")
	mockCtx.EXPECT().FormValue("birth_date").Return("2000-01-01")
	mockCtx.EXPECT().FormValue("cpf").Return("11144477735")
	mockCtx.EXPECT().FormValue("email").Return("test@test.com")
	mockCtx.EXPECT().FormValue("phone").Return("123456789")
	mockCtx.EXPECT().FormValue("registration").Return("20211234")
	mockCtx.EXPECT().FormValue("campus_id").Return(campusID.String())
	mockCtx.EXPECT().FormValue("course_id").Return(courseID.String())
	mockCtx.EXPECT().FormValue("total_workload").Return("100")

	id := uuid.New()
	mockService.EXPECT().Create(gomock.Any()).Return(&id, nil)
	mockCtx.EXPECT().JSON(http.StatusCreated, gomock.Any()).Return(nil)

	err := handler.Create(mockCtx)
	assert.Nil(t, err)
}
