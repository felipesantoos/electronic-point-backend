package views

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/handlers/dto/request"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/handlers/views/helpers"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/services/filters"
	"net/http"

	"github.com/google/uuid"
)

type InternshipLocationViewHandlers interface {
	List(handlers.RichContext) error
	CreatePage(handlers.RichContext) error
	Create(handlers.RichContext) error
	EditPage(handlers.RichContext) error
	Update(handlers.RichContext) error
}

type internshipLocationViewHandlers struct {
	service primary.InternshipLocationPort
}

func NewInternshipLocationViewHandlers(service primary.InternshipLocationPort) InternshipLocationViewHandlers {
	return &internshipLocationViewHandlers{service}
}

func (h *internshipLocationViewHandlers) List(ctx handlers.RichContext) error {
	locations, err := h.service.List(filters.InternshipLocationFilters{})
	if err != nil {
		return ctx.Render(http.StatusOK, "internship-locations/list.html", helpers.PageData{Errors: []string{err.String()}})
	}

	data := struct {
		Locations []response.InternshipLocation
	}{
		Locations: response.InternshipLocationBuilder().BuildFromDomainList(locations),
	}

	return ctx.Render(http.StatusOK, "internship-locations/list.html", helpers.NewPageData(ctx, "Locais de Estágio", "internship-locations", data))
}

func (h *internshipLocationViewHandlers) CreatePage(ctx handlers.RichContext) error {
	return ctx.Render(http.StatusOK, "internship-locations/create.html", helpers.NewPageData(ctx, "Novo Local", "internship-locations", nil))
}

func (h *internshipLocationViewHandlers) Create(ctx handlers.RichContext) error {
	var dto request.InternshipLocation
	if err := ctx.Bind(&dto); err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{"Dados inválidos"}})
	}

	loc, dErr := dto.ToDomain()
	if dErr != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{dErr.String()}})
	}

	_, err := h.service.Create(loc)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	ctx.Response().Header().Set("HX-Redirect", "/internship-locations")
	return ctx.NoContent(http.StatusCreated)
}

func (h *internshipLocationViewHandlers) EditPage(ctx handlers.RichContext) error {
	id, _ := uuid.Parse(ctx.Param("id"))
	loc, err := h.service.Get(id)
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/internship-locations")
	}

	data := struct {
		Location response.InternshipLocation
	}{
		Location: response.InternshipLocationBuilder().BuildFromDomain(loc),
	}

	return ctx.Render(http.StatusOK, "internship-locations/edit.html", helpers.NewPageData(ctx, "Editar Local", "internship-locations", data))
}

func (h *internshipLocationViewHandlers) Update(ctx handlers.RichContext) error {
	id, _ := uuid.Parse(ctx.Param("id"))
	var dto request.InternshipLocation
	if err := ctx.Bind(&dto); err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{"Dados inválidos"}})
	}

	loc, dErr := dto.ToDomain()
	if dErr != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{dErr.String()}})
	}
	loc.SetID(id)

	err := h.service.Update(loc)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	ctx.Response().Header().Set("HX-Redirect", "/internship-locations")
	return ctx.NoContent(http.StatusOK)
}
