package views

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/handlers/dto/request"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/handlers/views/helpers"
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/services/filters"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

type InternshipViewHandlers interface {
	List(handlers.RichContext) error
	CreatePage(handlers.RichContext) error
	Create(handlers.RichContext) error
	Show(handlers.RichContext) error
}

type internshipViewHandlers struct {
	service          primary.InternshipPort
	studentService   primary.StudentPort
	locationService  primary.InternshipLocationPort
	accountService   primary.AccountPort
	resourcesService primary.ResourcesPort
}

func NewInternshipViewHandlers(
	service primary.InternshipPort,
	studentService primary.StudentPort,
	locationService primary.InternshipLocationPort,
	accountService primary.AccountPort,
	resourcesService primary.ResourcesPort,
) InternshipViewHandlers {
	return &internshipViewHandlers{service, studentService, locationService, accountService, resourcesService}
}

func (h *internshipViewHandlers) List(ctx handlers.RichContext) error {
	f := helpers.GetInternshipFilters(ctx)
	internships, err := h.service.List(f)
	if err != nil {
		return ctx.Render(http.StatusOK, "internships/list.html", helpers.PageData{Errors: []string{err.String()}})
	}

	students, _ := h.studentService.List(filters.StudentFilters{})

	data := map[string]interface{}{
		"Internships": response.InternshipBuilder().BuildFromDomainList(internships),
		"Students":    helpers.ToOptions(students),
		"Filters":     f,
	}

	return ctx.Render(http.StatusOK, "internships/list.html", helpers.NewPageData(ctx, "Estágios", "internships", data).
		WithBreadcrumbs(helpers.Breadcrumb{Label: "Estágios", URL: "/internships"}))
}

func (h *internshipViewHandlers) CreatePage(ctx handlers.RichContext) error {
	students, _ := h.studentService.List(filters.StudentFilters{})
	locations, _ := h.locationService.List(filters.InternshipLocationFilters{})

	selectedStudentID := ctx.QueryParam("student_id")

	data := map[string]interface{}{
		"Students":          helpers.ToOptions(students),
		"Locations":         helpers.ToOptions(locations),
		"SelectedStudentID": selectedStudentID,
	}

	if ctx.IsAdmin() {
		roles, _ := h.resourcesService.ListAccountRoles()
		var teacherRoleID *uuid.UUID
		for _, r := range roles {
			if strings.ToLower(r.Code()) == role.TEACHER_ROLE_CODE {
				teacherRoleID = r.ID()
				break
			}
		}
		if teacherRoleID != nil {
			teachers, _ := h.accountService.List(filters.AccountFilters{RoleID: teacherRoleID})
			data["Teachers"] = helpers.ToOptions(teachers)
		}
	}

	return ctx.Render(http.StatusOK, "internships/create.html", helpers.NewPageData(ctx, "Vincular Estágio", "internships", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Estágios", URL: "/internships"},
			helpers.Breadcrumb{Label: "Novo", URL: "/internships/new"},
		))
}

func (h *internshipViewHandlers) Create(ctx handlers.RichContext) error {
	var body struct {
		StudentID         string `form:"student_id"`
		TeacherID         string `form:"teacher_id"`
		LocationID        string `form:"location_id"`
		StartedIn         string `form:"started_in"`
		EndedIn           string `form:"ended_in"`
		ScheduleEntryTime string `form:"schedule_entry_time"`
		ScheduleExitTime  string `form:"schedule_exit_time"`
	}
	if err := ctx.Bind(&body); err != nil {
		return helpers.HTMXError(ctx, http.StatusBadRequest, "Dados inválidos")
	}

	studentID, _ := uuid.Parse(body.StudentID)
	locationID, _ := uuid.Parse(body.LocationID)
	startedIn, _ := time.Parse("2006-01-02", body.StartedIn)

	var teacherID *uuid.UUID
	if ctx.RoleName() == role.ADMIN_ROLE_CODE {
		if val := body.TeacherID; val != "" {
			if uid, err := uuid.Parse(val); err == nil {
				teacherID = &uid
			}
		}
	} else {
		teacherID = ctx.ProfileID()
	}

	var endedInPtr *time.Time
	if body.EndedIn != "" {
		endedIn, _ := time.Parse("2006-01-02", body.EndedIn)
		endedInPtr = &endedIn
	}

	var entryTimePtr *time.Time
	if body.ScheduleEntryTime != "" {
		t, err := time.Parse("15:04", body.ScheduleEntryTime)
		if err == nil {
			entryTimePtr = &t
		}
	}

	var exitTimePtr *time.Time
	if body.ScheduleExitTime != "" {
		t, err := time.Parse("15:04", body.ScheduleExitTime)
		if err == nil {
			exitTimePtr = &t
		}
	}

	dto := request.Internship{
		StudentID:         studentID,
		TeacherID:         teacherID,
		LocationID:        locationID,
		StartedIn:         startedIn,
		EndedIn:           endedInPtr,
		ScheduleEntryTime: entryTimePtr,
		ScheduleExitTime:  exitTimePtr,
	}

	intern, dErr := dto.ToDomain()
	if dErr != nil {
		return helpers.HTMXError(ctx, http.StatusUnprocessableEntity, dErr.String())
	}

	_, err := h.service.Create(intern)
	if err != nil {
		return helpers.HTMXError(ctx, http.StatusBadRequest, err.String())
	}

	ctx.Response().Header().Set("HX-Redirect", "/internships")
	return ctx.NoContent(http.StatusCreated)
}

func (h *internshipViewHandlers) Show(ctx handlers.RichContext) error {
	id, _ := uuid.Parse(ctx.Param("id"))
	intern, err := h.service.Get(id)
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/internships")
	}

	data := map[string]interface{}{
		"Internship": response.InternshipBuilder().BuildFromDomain(intern),
	}

	return ctx.Render(http.StatusOK, "internships/show.html", helpers.NewPageData(ctx, "Detalhes do Estágio", "internships", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Estágios", URL: "/internships"},
			helpers.Breadcrumb{Label: "Detalhes", URL: "/internships/" + id.String()},
		))
}
