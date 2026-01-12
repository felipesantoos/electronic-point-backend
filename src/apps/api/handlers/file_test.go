package handlers

import (
	"eletronic_point/src/apps/api/handlers/mocks"
	"io"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestFileHandler_Get(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockFilePort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewFileHandlers(mockService)

	// Create a temporary file to mock a real file
	tmpFile, _ := os.CreateTemp("", "test.txt")
	defer os.Remove(tmpFile.Name())
	tmpFile.WriteString("test content")
	tmpFile.Seek(0, io.SeekStart)

	rec := httptest.NewRecorder()
	echoRes := echo.NewResponse(rec, echo.New())

	mockCtx.EXPECT().Param("name").Return("test.txt")
	mockService.EXPECT().Get("test.txt").Return(tmpFile, nil)
	mockCtx.EXPECT().Response().Return(echoRes).AnyTimes()

	err := handler.Get(mockCtx)
	assert.Nil(t, err)
}
