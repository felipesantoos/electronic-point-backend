package handlers

import (
	"eletronic_point/src/apps/api/handlers/mocks"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestInternshipLocationHandler_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockInternshipLocationPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewInternshipLocationHandlers(mockService)

	mockCtx.EXPECT().QueryParam(gomock.Any()).Return("").AnyTimes()
	mockService.EXPECT().List(gomock.Any()).Return(nil, nil)
	mockCtx.EXPECT().JSON(http.StatusOK, gomock.Any()).Return(nil)

	err := handler.List(mockCtx)
	assert.Nil(t, err)
}
