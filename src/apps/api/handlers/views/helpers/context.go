package helpers

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/handlers/dto/response"

	"github.com/google/uuid"
)

// PageData is the common structure for all view templates
type PageData struct {
	Title         string
	User          UserInfo
	ActiveMenu    string
	Data          interface{}
	FlashMessages []FlashMessage
	Errors        []string
}

// UserInfo contains basic information about the logged-in user for templates
type UserInfo struct {
	ID        *uuid.UUID
	ProfileID *uuid.UUID
	RoleName  string
	IsAdmin   bool
}

// FlashMessage represents a temporary message to be displayed to the user
type FlashMessage struct {
	Type    string // success, error, warning, info
	Content string
}

// NewPageData creates a new PageData with user info and flash messages from context
func NewPageData(ctx handlers.RichContext, title string, activeMenu string, data interface{}) PageData {
	return PageData{
		Title: title,
		User: UserInfo{
			ID:        ctx.AccountID(),
			ProfileID: ctx.ProfileID(),
			RoleName:  ctx.RoleName(),
			IsAdmin:   ctx.IsAdmin(),
		},
		ActiveMenu:    activeMenu,
		Data:          data,
		FlashMessages: GetFlashMessages(ctx),
	}
}

// ErrorResponse attempts to convert an API error response to a string list
func GetErrorsFromResponse(errResponse interface{}) []string {
	if err, ok := errResponse.(response.ErrorMessage); ok {
		return []string{err.Message}
	}
	return nil
}
