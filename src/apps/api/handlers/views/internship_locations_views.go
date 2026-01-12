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
	Show(handlers.RichContext) error
}

type internshipLocationViewHandlers struct {
	service           primary.InternshipLocationPort
	internshipService primary.InternshipPort
}

func NewInternshipLocationViewHandlers(
	service primary.InternshipLocationPort,
	internshipService primary.InternshipPort,
) InternshipLocationViewHandlers {
	return &internshipLocationViewHandlers{service, internshipService}
}

func (h *internshipLocationViewHandlers) List(ctx handlers.RichContext) error {
	f := helpers.GetInternshipLocationFilters(ctx)

	locations, err := h.service.List(f)
	if err != nil {
		return ctx.Render(http.StatusOK, "internship-locations/list.html", helpers.PageData{Errors: []string{err.String()}})
	}

	data := map[string]interface{}{
		"Locations": response.InternshipLocationBuilder().BuildFromDomainList(locations),
		"Filters":   f,
	}

	return ctx.Render(http.StatusOK, "internship-locations/list.html", helpers.NewPageData(ctx, "Locais de Estágio", "internship-locations", data).
		WithBreadcrumbs(helpers.Breadcrumb{Label: "Locais de Estágio", URL: "/internship-locations"}))
}

func (h *internshipLocationViewHandlers) CreatePage(ctx handlers.RichContext) error {
	return ctx.Render(http.StatusOK, "internship-locations/create.html", helpers.NewPageData(ctx, "Novo Local", "internship-locations", nil))
}

func (h *internshipLocationViewHandlers) Create(ctx handlers.RichContext) error {
	var dto request.InternshipLocation
	if err := ctx.Bind(&dto); err != nil {
		return helpers.HTMXError(ctx, http.StatusBadRequest, "Dados inválidos")
	}

	loc, dErr := dto.ToDomain()
	if dErr != nil {
		return helpers.HTMXError(ctx, http.StatusUnprocessableEntity, dErr.String())
	}

	_, err := h.service.Create(loc)
	if err != nil {
		return helpers.HTMXError(ctx, http.StatusBadRequest, err.String())
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
		return helpers.HTMXError(ctx, http.StatusBadRequest, "Dados inválidos")
	}

	loc, dErr := dto.ToDomain()
	if dErr != nil {
		return helpers.HTMXError(ctx, http.StatusUnprocessableEntity, dErr.String())
	}
	loc.SetID(id)

	err := h.service.Update(loc)
	if err != nil {
		return helpers.HTMXError(ctx, http.StatusBadRequest, err.String())
	}

	ctx.Response().Header().Set("HX-Redirect", "/internship-locations")
	return ctx.NoContent(http.StatusOK)
}

func (h *internshipLocationViewHandlers) Show(ctx handlers.RichContext) error {
	id, _ := uuid.Parse(ctx.Param("id"))
	loc, err := h.service.Get(id)
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/internship-locations")
	}

	// Fetch related internships
	internships, _ := h.internshipService.List(filters.InternshipFilters{})
	// Filter internships manually for this location if the service doesn't support it directly
	// Actually, let's see if InternshipFilters supports LocationID
	// Checking src/core/services/filters/internship.go

	locationInternships := make([]response.Internship, 0)
	allInternships := response.InternshipBuilder().BuildFromDomainList(internships)
	for _, intern := range allInternships {
		if intern.Location.ID == id {
			locationInternships = append(locationInternships, intern)
		}
	}

	data := map[string]interface{}{
		"Location":    response.InternshipLocationBuilder().BuildFromDomain(loc),
		"Internships": locationInternships,
	}

	return ctx.Render(http.StatusOK, "internship-locations/show.html", helpers.NewPageData(ctx, "Detalhes do Local", "internship-locations", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Locais de Estágio", URL: "/internship-locations"},
			helpers.Breadcrumb{Label: "Detalhes", URL: "/internship-locations/" + id.String()},
		))
}
