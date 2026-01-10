package views

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/handlers/views/helpers"
	"eletronic_point/src/core/domain/timeRecordStatus"
	"eletronic_point/src/core/interfaces/primary"
	"net/http"

	"github.com/google/uuid"
)

type TimeRecordStatusViewHandlers interface {
	List(handlers.RichContext) error
	CreatePage(handlers.RichContext) error
	Create(handlers.RichContext) error
	EditPage(handlers.RichContext) error
	Update(handlers.RichContext) error
	Show(handlers.RichContext) error
}

type timeRecordStatusViewHandlers struct {
	service primary.TimeRecordStatusPort
}

func NewTimeRecordStatusViewHandlers(service primary.TimeRecordStatusPort) TimeRecordStatusViewHandlers {
	return &timeRecordStatusViewHandlers{service}
}

func (h *timeRecordStatusViewHandlers) List(ctx handlers.RichContext) error {
	name := ctx.QueryParam("name")
	statuses, err := h.service.List()
	if err != nil {
		return ctx.Render(http.StatusOK, "time-record-status/list.html", helpers.PageData{Errors: []string{err.String()}})
	}

	data := map[string]interface{}{
		"Statuses": response.TimeRecordStatusBuilder().BuildFromDomainList(statuses),
		"Filters":  map[string]string{"name": name},
	}

	return ctx.Render(http.StatusOK, "time-record-status/list.html", helpers.NewPageData(ctx, "Status de Ponto", "time-record-status", data).
		WithBreadcrumbs(helpers.Breadcrumb{Label: "Status de Ponto", URL: "/time-record-status"}))
}

func (h *timeRecordStatusViewHandlers) CreatePage(ctx handlers.RichContext) error {
	return ctx.Render(http.StatusOK, "time-record-status/create.html", helpers.NewPageData(ctx, "Novo Status", "time-record-status", nil).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Status de Ponto", URL: "/time-record-status"},
			helpers.Breadcrumb{Label: "Novo", URL: "/time-record-status/new"},
		))
}

func (h *timeRecordStatusViewHandlers) Create(ctx handlers.RichContext) error {
	name := ctx.FormValue("name")
	if name == "" {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{"O nome é obrigatório"}})
	}

	status, err := timeRecordStatus.NewBuilder().WithName(name).Build()
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	_, err = h.service.Create(status)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	ctx.Response().Header().Set("HX-Redirect", "/time-record-status")
	return ctx.NoContent(http.StatusCreated)
}

func (h *timeRecordStatusViewHandlers) EditPage(ctx handlers.RichContext) error {
	id, _ := uuid.Parse(ctx.Param("id"))
	status, err := h.service.Get(id)
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/time-record-status")
	}

	data := struct {
		Status response.TimeRecordStatus
	}{
		Status: response.TimeRecordStatusBuilder().BuildFromDomain(status),
	}

	return ctx.Render(http.StatusOK, "time-record-status/edit.html", helpers.NewPageData(ctx, "Editar Status", "time-record-status", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Status de Ponto", URL: "/time-record-status"},
			helpers.Breadcrumb{Label: status.Name(), URL: "/time-record-status/" + id.String()},
			helpers.Breadcrumb{Label: "Editar", URL: "/time-record-status/" + id.String() + "/edit"},
		))
}

func (h *timeRecordStatusViewHandlers) Update(ctx handlers.RichContext) error {
	id, _ := uuid.Parse(ctx.Param("id"))
	name := ctx.FormValue("name")

	status, err := timeRecordStatus.NewBuilder().WithID(id).WithName(name).Build()
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	err = h.service.Update(status)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	ctx.Response().Header().Set("HX-Redirect", "/time-record-status/"+id.String())
	return ctx.NoContent(http.StatusOK)
}

func (h *timeRecordStatusViewHandlers) Show(ctx handlers.RichContext) error {
	id, _ := uuid.Parse(ctx.Param("id"))
	status, err := h.service.Get(id)
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/time-record-status")
	}

	data := struct {
		Status response.TimeRecordStatus
	}{
		Status: response.TimeRecordStatusBuilder().BuildFromDomain(status),
	}

	return ctx.Render(http.StatusOK, "time-record-status/show.html", helpers.NewPageData(ctx, status.Name(), "time-record-status", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Status de Ponto", URL: "/time-record-status"},
			helpers.Breadcrumb{Label: status.Name(), URL: "/time-record-status/" + id.String()},
		))
}
