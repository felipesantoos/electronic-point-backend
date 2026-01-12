package handlers

import (
	"eletronic_point/src/apps/api/handlers/mocks"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAccountHandler_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAccountPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewAccountHandler(mockService)

	mockCtx.EXPECT().QueryParam("search").Return("test")
	mockCtx.EXPECT().QueryParam("role_id").Return("")
	mockService.EXPECT().List(gomock.Any()).Return(nil, nil)
	mockCtx.EXPECT().JSON(http.StatusOK, gomock.Any()).Return(nil)

	err := handler.List(mockCtx)
	assert.Nil(t, err)
}

func TestAccountHandler_Delete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAccountPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewAccountHandler(mockService)

	id := uuid.New()
	mockCtx.EXPECT().Param("id").Return(id.String())
	mockService.EXPECT().Delete(id).Return(nil)
	mockCtx.EXPECT().NoContent(http.StatusOK).Return(nil)

	err := handler.Delete(mockCtx)
	assert.Nil(t, err)
}
