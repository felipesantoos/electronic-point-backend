package views

import (
	"eletronic_point/src/apps/api/handlers/mocks"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAccountViewHandler_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAccountPort(ctrl)
	mockResources := mocks.NewMockResourcesPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)

	handler := NewAccountViewHandlers(mockService, mockResources)

	mockCtx.EXPECT().IsAdmin().Return(true).AnyTimes()
	mockCtx.EXPECT().AccountID().Return(nil).AnyTimes()
	mockCtx.EXPECT().ProfileID().Return(nil).AnyTimes()
	mockCtx.EXPECT().RoleName().Return("admin").AnyTimes()
	mockCtx.EXPECT().QueryParam(gomock.Any()).Return("").AnyTimes()
	mockCtx.EXPECT().Request().Return(&http.Request{}).AnyTimes()
	mockService.EXPECT().List(gomock.Any()).Return(nil, nil)
	mockResources.EXPECT().ListAccountRoles().Return(nil, nil)
	mockCtx.EXPECT().Cookie("ep_flash").Return(nil, http.ErrNoCookie).AnyTimes()
	mockCtx.EXPECT().Render(http.StatusOK, "accounts/list.html", gomock.Any()).Return(nil)

	err := handler.List(mockCtx)
	assert.Nil(t, err)
}
