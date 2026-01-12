package handlers

import (
	"eletronic_point/src/apps/api/handlers/mocks"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestInstitutionHandler_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockInstitutionPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewInstitutionHandlers(mockService)

	mockCtx.EXPECT().QueryParam("name").Return("")
	mockService.EXPECT().List(gomock.Any()).Return(nil, nil)
	mockCtx.EXPECT().JSON(http.StatusOK, gomock.Any()).Return(nil)

	err := handler.List(mockCtx)
	assert.Nil(t, err)
}
