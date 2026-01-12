package handlers

import (
	"eletronic_point/src/apps/api/handlers/mocks"
	"eletronic_point/src/core/domain/authorization"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestAuthHandler_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mocks.NewMockAuthPort(ctrl)
	mockCtx := mocks.NewMockRichContext(ctrl)
	handler := NewAuthHandler(mockService)

	mockCtx.EXPECT().AccountID().Return(nil)
	mockCtx.EXPECT().Bind(gomock.Any()).Return(nil)
	mockCtx.EXPECT().JSON(gomock.Any(), gomock.Any()).Return(nil)

	expTime := time.Now().Add(time.Hour)
	auth := authorization.NewFromToken("test-token", &expTime)
	mockService.EXPECT().Login(gomock.Any()).Return(auth, nil)

	err := handler.Login(mockCtx)
	assert.Nil(t, err)
}
