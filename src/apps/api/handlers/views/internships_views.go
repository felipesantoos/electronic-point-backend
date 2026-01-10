package views

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/handlers/dto/request"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/handlers/views/helpers"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/services/filters"
	"net/http"
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
	locationService primary.InternshipLocationPort
}

func NewInternshipViewHandlers(
	service primary.InternshipPort,
	studentService primary.StudentPort,
	locationService primary.InternshipLocationPort,
) InternshipViewHandlers {
	return &internshipViewHandlers{service, studentService, locationService}
}

func (h *internshipViewHandlers) List(ctx handlers.RichContext) error {
	f := helpers.GetInternshipFilters(ctx)
	internships, err := h.service.List(f)
	if err != nil {
		return ctx.Render(http.StatusOK, "internships/list.html", helpers.PageData{Errors: []string{err.String()}})
	}

	students, _ := h.studentService.List(filters.StudentFilters{})

	data := struct {
		Internships []response.Internship
		Students    interface{}
		Filters     filters.InternshipFilters
	}{
		Internships: response.InternshipBuilder().BuildFromDomainList(internships),
		Students:    helpers.ToOptions(students),
		Filters:     f,
	}

	return ctx.Render(http.StatusOK, "internships/list.html", helpers.NewPageData(ctx, "Estágios", "internships", data).
		WithBreadcrumbs(helpers.Breadcrumb{Label: "Estágios", URL: "/internships"}))
}

func (h *internshipViewHandlers) CreatePage(ctx handlers.RichContext) error {
	students, _ := h.studentService.List(filters.StudentFilters{})
	locations, _ := h.locationService.List(filters.InternshipLocationFilters{})

	selectedStudentID := ctx.QueryParam("student_id")

	data := struct {
		Students          interface{}
		Locations         interface{}
		SelectedStudentID string
	}{
		Students:          helpers.ToOptions(students),
		Locations:         helpers.ToOptions(locations),
		SelectedStudentID: selectedStudentID,
	}

	return ctx.Render(http.StatusOK, "internships/create.html", helpers.NewPageData(ctx, "Vincular Estágio", "internships", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Estágios", URL: "/internships"},
			helpers.Breadcrumb{Label: "Novo", URL: "/internships/new"},
		))
}

func (h *internshipViewHandlers) Create(ctx handlers.RichContext) error {
	var body struct {
		StudentID  string `form:"student_id"`
		LocationID string `form:"location_id"`
		StartedIn  string `form:"started_in"`
		EndedIn    string `form:"ended_in"`
	}
	if err := ctx.Bind(&body); err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{"Dados inválidos"}})
	}

	studentID, _ := uuid.Parse(body.StudentID)
	locationID, _ := uuid.Parse(body.LocationID)
	startedIn, _ := time.Parse("2006-01-02", body.StartedIn)
	
	var endedInPtr *time.Time
	if body.EndedIn != "" {
		endedIn, _ := time.Parse("2006-01-02", body.EndedIn)
		endedInPtr = &endedIn
	}

	dto := request.Internship{
		StudentID:  studentID,
		LocationID: locationID,
		StartedIn:  startedIn,
		EndedIn:    endedInPtr,
	}

	intern, dErr := dto.ToDomain()
	if dErr != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{dErr.String()}})
	}

	_, err := h.service.Create(intern)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
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

	data := struct {
		Internship response.Internship
	}{
		Internship: response.InternshipBuilder().BuildFromDomain(intern),
	}

	return ctx.Render(http.StatusOK, "internships/show.html", helpers.NewPageData(ctx, "Detalhes do Estágio", "internships", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Estágios", URL: "/internships"},
			helpers.Breadcrumb{Label: "Detalhes", URL: "/internships/" + id.String()},
		))
}

