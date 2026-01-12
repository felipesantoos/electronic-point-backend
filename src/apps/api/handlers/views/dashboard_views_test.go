package views

import (
	"eletronic_point/src/apps/api/handlers/mocks"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestDashboardViewHandler_Dashboard(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStudentService := mocks.NewMockStudentPort(ctrl)
	mockInternshipService := mocks.NewMockInternshipPort(ctrl)
	mockTimeRecordService := mocks.NewMockTimeRecordPort(ctrl)
	mockLocationService := mocks.NewMockInternshipLocationPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)

	handler := NewDashboardViewHandlers(mockStudentService, mockInternshipService, mockTimeRecordService, mockLocationService)

	mockCtx.EXPECT().RoleName().Return("admin").AnyTimes()
	mockCtx.EXPECT().AccountID().Return(nil).AnyTimes()
	mockCtx.EXPECT().ProfileID().Return(nil).AnyTimes()
	mockCtx.EXPECT().IsAdmin().Return(true).AnyTimes()
	mockCtx.EXPECT().Request().Return(&http.Request{}).AnyTimes()
	mockStudentService.EXPECT().List(gomock.Any()).Return(nil, nil)
	mockInternshipService.EXPECT().List(gomock.Any()).Return(nil, nil)
	mockTimeRecordService.EXPECT().List(gomock.Any()).Return(nil, nil)
	mockLocationService.EXPECT().List(gomock.Any()).Return(nil, nil)
	mockCtx.EXPECT().Cookie("ep_flash").Return(nil, http.ErrNoCookie).AnyTimes()
	mockCtx.EXPECT().Render(http.StatusOK, "dashboard", gomock.Any()).Return(nil)

	err := handler.Dashboard(mockCtx)
	assert.Nil(t, err)
}
