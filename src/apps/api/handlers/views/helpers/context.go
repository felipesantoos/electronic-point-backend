package helpers

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"strings"

	"github.com/google/uuid"
)

// Breadcrumb represents a single step in the navigation path
type Breadcrumb struct {
	Label string
	URL   string
}

// PageData is the common structure for all view templates
type PageData struct {
	Title         string
	User          UserInfo
	ActiveMenu    string
	Data          interface{}
	FlashMessages []FlashMessage
	Errors        []string
	Breadcrumbs   []Breadcrumb
}

// UserInfo contains basic information about the logged-in user for templates
type UserInfo struct {
	ID        *uuid.UUID
	Name      string
	FirstName string
	ProfileID *uuid.UUID
	RoleName  string
	IsAdmin   bool
	IsTeacher bool
	IsStudent bool
}

// FlashMessage represents a temporary message to be displayed to the user
type FlashMessage struct {
	Type    string // success, error, warning, info
	Content string
}

// NewPageData creates a new PageData with user info and flash messages from context
func NewPageData(ctx handlers.RichContext, title string, activeMenu string, data interface{}) PageData {
	name := ctx.Name()
	firstName := name
	if strings.Contains(name, " ") {
		firstName = strings.Split(name, " ")[0]
	}

	return PageData{
		Title: title,
		User: UserInfo{
			ID:        ctx.AccountID(),
			Name:      name,
			FirstName: firstName,
			ProfileID: ctx.ProfileID(),
			RoleName:  ctx.RoleName(),
			IsAdmin:   ctx.IsAdmin(),
			IsTeacher: strings.ToLower(ctx.RoleName()) == "teacher" || strings.ToLower(ctx.RoleName()) == "professor",
			IsStudent: strings.ToLower(ctx.RoleName()) == "student" || strings.ToLower(ctx.RoleName()) == "estudante",
		},
		ActiveMenu:    activeMenu,
		Data:          data,
		FlashMessages: GetFlashMessages(ctx),
		Breadcrumbs:   []Breadcrumb{{Label: "Dashboard", URL: "/"}},
	}
}

// WithBreadcrumbs adds custom breadcrumbs to the PageData
func (p PageData) WithBreadcrumbs(breadcrumbs ...Breadcrumb) PageData {
	// Only add Dashboard if not already there and if there are other breadcrumbs
	if len(breadcrumbs) > 0 && breadcrumbs[0].URL != "/" {
		p.Breadcrumbs = append([]Breadcrumb{{Label: "Dashboard", URL: "/"}}, breadcrumbs...)
	} else {
		p.Breadcrumbs = breadcrumbs
	}
	return p
}

// ErrorResponse attempts to convert an API error response to a string list
func GetErrorsFromResponse(errResponse interface{}) []string {
	if err, ok := errResponse.(response.ErrorMessage); ok {
		return []string{err.Message}
	}
	return nil
}
