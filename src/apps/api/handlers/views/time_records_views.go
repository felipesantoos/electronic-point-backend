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

type TimeRecordViewHandlers interface {
	List(handlers.RichContext) error
	CreatePage(handlers.RichContext) error
	Create(handlers.RichContext) error
	Show(handlers.RichContext) error
}

type timeRecordViewHandlers struct {
	service        primary.TimeRecordPort
	studentService primary.StudentPort
	statusService  primary.TimeRecordStatusPort
}

func NewTimeRecordViewHandlers(
	service primary.TimeRecordPort,
	studentService primary.StudentPort,
	statusService primary.TimeRecordStatusPort,
) TimeRecordViewHandlers {
	return &timeRecordViewHandlers{service, studentService, statusService}
}

func (h *timeRecordViewHandlers) List(ctx handlers.RichContext) error {
	f := helpers.GetTimeRecordFilters(ctx)
	records, err := h.service.List(f)
	if err != nil {
		return ctx.Render(http.StatusOK, "time-records/list.html", helpers.PageData{Errors: []string{err.String()}})
	}

	students, _ := h.studentService.List(filters.StudentFilters{})
	statuses, _ := h.statusService.List()

	data := map[string]interface{}{
		"TimeRecords": response.TimeRecordBuilder().BuildFromDomainList(records),
		"Students":    helpers.ToOptions(students),
		"Statuses":    helpers.ToOptions(statuses),
		"Filters":     f,
	}

	return ctx.Render(http.StatusOK, "time-records/list.html", helpers.NewPageData(ctx, "Registros de Ponto", "time-records", data).
		WithBreadcrumbs(helpers.Breadcrumb{Label: "Registros de Ponto", URL: "/time-records"}))
}

func (h *timeRecordViewHandlers) CreatePage(ctx handlers.RichContext) error {
	students, _ := h.studentService.List(filters.StudentFilters{})

	data := map[string]interface{}{
		"Students": helpers.ToOptions(students),
	}

	return ctx.Render(http.StatusOK, "time-records/create.html", helpers.NewPageData(ctx, "Novo Registro", "time-records", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Registros de Ponto", URL: "/time-records"},
			helpers.Breadcrumb{Label: "Novo", URL: "/time-records/new"},
		))
}

func (h *timeRecordViewHandlers) Create(ctx handlers.RichContext) error {
	var body struct {
		StudentID     string `form:"student_id"`
		Date          string `form:"date"`
		EntryTime     string `form:"entry_time"`
		ExitTime      string `form:"exit_time"`
		Location      string `form:"location"`
		IsOffSite     bool   `form:"is_off_site"`
		Justification string `form:"justification"`
	}
	if err := ctx.Bind(&body); err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{"Dados inválidos"}})
	}

	studentID, _ := uuid.Parse(body.StudentID)
	date, _ := time.Parse("2006-01-02", body.Date)
	entryTime, _ := time.Parse("2006-01-02T15:04", body.EntryTime)

	var exitTimePtr *time.Time
	if body.ExitTime != "" {
		exitTime, _ := time.Parse("2006-01-02T15:04", body.ExitTime)
		exitTimePtr = &exitTime
	}

	var justificationPtr *string
	if body.Justification != "" {
		justificationPtr = &body.Justification
	}

	dto := request.TimeRecord{
		Date:          date,
		EntryTime:     entryTime,
		ExitTime:      exitTimePtr,
		Location:      body.Location,
		IsOffSite:     body.IsOffSite,
		Justification: justificationPtr,
	}

	tr, dErr := dto.ToDomain()
	if dErr != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{dErr.String()}})
	}
	tr.SetStudentID(studentID)

	_, err := h.service.Create(tr)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	ctx.Response().Header().Set("HX-Redirect", "/time-records")
	return ctx.NoContent(http.StatusCreated)
}

func (h *timeRecordViewHandlers) Show(ctx handlers.RichContext) error {
	id, _ := uuid.Parse(ctx.Param("id"))
	tr, err := h.service.Get(id, filters.TimeRecordFilters{})
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/time-records")
	}

	trResponse := response.TimeRecordBuilder().BuildFromDomain(tr)
	data := map[string]interface{}{
		"TimeRecord": trResponse,
	}

	return ctx.Render(http.StatusOK, "time-records/show.html", helpers.NewPageData(ctx, "Detalhes do Registro", "time-records", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Registros de Ponto", URL: "/time-records"},
			helpers.Breadcrumb{Label: trResponse.Student.Name, URL: "/time-records/" + id.String()},
		))
}
