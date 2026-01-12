package handlers

import (
	"eletronic_point/src/apps/api/handlers/mocks"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestTimeRecordStatusHandler_List(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockTimeRecordStatusPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewTimeRecordStatusHandlers(mockService)

	mockService.EXPECT().List().Return(nil, nil)
	mockCtx.EXPECT().JSON(http.StatusOK, gomock.Any()).Return(nil)

	err := handler.List(mockCtx)
	assert.Nil(t, err)
}
