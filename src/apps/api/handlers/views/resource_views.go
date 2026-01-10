package views

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/handlers/views/helpers"
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/domain/course"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/services/filters"
	"net/http"

	"github.com/google/uuid"
)

type ResourceViewHandlers interface {
	Institutions(handlers.RichContext) error
	CreateInstitutionPage(handlers.RichContext) error
	CreateInstitution(handlers.RichContext) error
	EditInstitutionPage(handlers.RichContext) error
	UpdateInstitution(handlers.RichContext) error

	Campus(handlers.RichContext) error
	CreateCampusPage(handlers.RichContext) error
	CreateCampus(handlers.RichContext) error
	EditCampusPage(handlers.RichContext) error
	UpdateCampus(handlers.RichContext) error

	Courses(handlers.RichContext) error
	CreateCoursePage(handlers.RichContext) error
	CreateCourse(handlers.RichContext) error
	EditCoursePage(handlers.RichContext) error
	UpdateCourse(handlers.RichContext) error
}

type resourceViewHandlers struct {
	institutionService primary.InstitutionPort
	campusService      primary.CampusPort
	courseService      primary.CoursePort
}

func NewResourceViewHandlers(
	institutionService primary.InstitutionPort,
	campusService primary.CampusPort,
	courseService primary.CoursePort,
) ResourceViewHandlers {
	return &resourceViewHandlers{
		institutionService,
		campusService,
		courseService,
	}
}

func (h *resourceViewHandlers) Institutions(ctx handlers.RichContext) error {
	name := ctx.QueryParam("name")
	var namePtr *string
	if name != "" {
		namePtr = &name
	}

	institutions, _ := h.institutionService.List(filters.InstitutionFilters{Name: namePtr})

	data := struct {
		Institutions []response.Institution
		Filters      map[string]string
	}{
		Institutions: response.InstitutionBuilder().BuildFromDomainList(institutions),
		Filters:      map[string]string{"name": name},
	}

	return ctx.Render(http.StatusOK, "institutions/list.html", helpers.NewPageData(ctx, "Instituições", "institutions", data).
		WithBreadcrumbs(helpers.Breadcrumb{Label: "Instituições", URL: "/institutions"}))
}

func (h *resourceViewHandlers) CreateInstitutionPage(ctx handlers.RichContext) error {
	return ctx.Render(http.StatusOK, "institutions/create.html", helpers.NewPageData(ctx, "Nova Instituição", "institutions", nil).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Instituições", URL: "/institutions"},
			helpers.Breadcrumb{Label: "Nova", URL: "/institutions/new"},
		))
}

func (h *resourceViewHandlers) CreateInstitution(ctx handlers.RichContext) error {
	name := ctx.FormValue("name")
	if name == "" {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{"O nome é obrigatório"}})
	}

	inst, err := institution.NewBuilder().WithName(name).Build()
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	_, err = h.institutionService.Create(inst)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	ctx.Response().Header().Set("HX-Redirect", "/institutions")
	return ctx.NoContent(http.StatusCreated)
}

func (h *resourceViewHandlers) EditInstitutionPage(ctx handlers.RichContext) error {
	id, _ := uuid.Parse(ctx.Param("id"))
	inst, err := h.institutionService.Get(id)
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/institutions")
	}

	data := struct {
		Institution response.Institution
	}{
		Institution: response.InstitutionBuilder().BuildFromDomain(inst),
	}

	return ctx.Render(http.StatusOK, "institutions/edit.html", helpers.NewPageData(ctx, "Editar Instituição", "institutions", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Instituições", URL: "/institutions"},
			helpers.Breadcrumb{Label: "Editar", URL: "/institutions/" + id.String() + "/edit"},
		))
}

func (h *resourceViewHandlers) UpdateInstitution(ctx handlers.RichContext) error {
	id, _ := uuid.Parse(ctx.Param("id"))
	name := ctx.FormValue("name")

	inst, err := institution.NewBuilder().WithID(id).WithName(name).Build()
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	err = h.institutionService.Update(inst)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	ctx.Response().Header().Set("HX-Redirect", "/institutions")
	return ctx.NoContent(http.StatusOK)
}

func (h *resourceViewHandlers) Campus(ctx handlers.RichContext) error {
	name := ctx.QueryParam("name")
	var namePtr *string
	if name != "" {
		namePtr = &name
	}

	campus, _ := h.campusService.List(filters.CampusFilters{Name: namePtr})

	data := struct {
		Campus  []response.Campus
		Filters map[string]string
	}{
		Campus:  response.CampusBuilder().BuildFromDomainList(campus),
		Filters: map[string]string{"name": name},
	}

	return ctx.Render(http.StatusOK, "campus/list.html", helpers.NewPageData(ctx, "Campus", "campus", data).
		WithBreadcrumbs(helpers.Breadcrumb{Label: "Campus", URL: "/campus"}))
}

func (h *resourceViewHandlers) CreateCampusPage(ctx handlers.RichContext) error {
	institutions, _ := h.institutionService.List(filters.InstitutionFilters{})

	data := struct {
		Institutions interface{}
	}{
		Institutions: helpers.ToOptions(institutions),
	}

	return ctx.Render(http.StatusOK, "campus/create.html", helpers.NewPageData(ctx, "Novo Campus", "campus", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Campus", URL: "/campus"},
			helpers.Breadcrumb{Label: "Novo", URL: "/campus/new"},
		))
}

func (h *resourceViewHandlers) CreateCampus(ctx handlers.RichContext) error {
	name := ctx.FormValue("name")
	institutionID, _ := uuid.Parse(ctx.FormValue("institution_id"))

	camp, err := campus.NewBuilder().WithName(name).WithInstitutionID(institutionID).Build()
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	_, err = h.campusService.Create(camp)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	ctx.Response().Header().Set("HX-Redirect", "/campus")
	return ctx.NoContent(http.StatusCreated)
}

func (h *resourceViewHandlers) EditCampusPage(ctx handlers.RichContext) error {
	id, _ := uuid.Parse(ctx.Param("id"))
	camp, err := h.campusService.Get(id)
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/campus")
	}

	institutions, _ := h.institutionService.List(filters.InstitutionFilters{})

	data := struct {
		Campus       response.Campus
		Institutions interface{}
	}{
		Campus:       response.CampusBuilder().BuildFromDomain(camp),
		Institutions: helpers.ToOptions(institutions),
	}

	return ctx.Render(http.StatusOK, "campus/edit.html", helpers.NewPageData(ctx, "Editar Campus", "campus", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Campus", URL: "/campus"},
			helpers.Breadcrumb{Label: "Editar", URL: "/campus/" + id.String() + "/edit"},
		))
}

func (h *resourceViewHandlers) UpdateCampus(ctx handlers.RichContext) error {
	id, _ := uuid.Parse(ctx.Param("id"))
	name := ctx.FormValue("name")
	institutionID, _ := uuid.Parse(ctx.FormValue("institution_id"))

	camp, err := campus.NewBuilder().WithID(id).WithName(name).WithInstitutionID(institutionID).Build()
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	err = h.campusService.Update(camp)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	ctx.Response().Header().Set("HX-Redirect", "/campus")
	return ctx.NoContent(http.StatusOK)
}

func (h *resourceViewHandlers) Courses(ctx handlers.RichContext) error {
	name := ctx.QueryParam("name")
	var namePtr *string
	if name != "" {
		namePtr = &name
	}

	courses, _ := h.courseService.List(filters.CourseFilters{Name: namePtr})

	data := struct {
		Courses []response.Course
		Filters map[string]string
	}{
		Courses: response.CourseBuilder().BuildFromDomainList(courses),
		Filters: map[string]string{"name": name},
	}

	return ctx.Render(http.StatusOK, "courses/list.html", helpers.NewPageData(ctx, "Cursos", "courses", data).
		WithBreadcrumbs(helpers.Breadcrumb{Label: "Cursos", URL: "/courses"}))
}

func (h *resourceViewHandlers) CreateCoursePage(ctx handlers.RichContext) error {
	return ctx.Render(http.StatusOK, "courses/create.html", helpers.NewPageData(ctx, "Novo Curso", "courses", nil).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Cursos", URL: "/courses"},
			helpers.Breadcrumb{Label: "Novo", URL: "/courses/new"},
		))
}

func (h *resourceViewHandlers) CreateCourse(ctx handlers.RichContext) error {
	name := ctx.FormValue("name")
	if name == "" {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{"O nome é obrigatório"}})
	}

	c, err := course.NewBuilder().WithName(name).Build()
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	_, err = h.courseService.Create(c)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	ctx.Response().Header().Set("HX-Redirect", "/courses")
	return ctx.NoContent(http.StatusCreated)
}

func (h *resourceViewHandlers) EditCoursePage(ctx handlers.RichContext) error {
	id, _ := uuid.Parse(ctx.Param("id"))
	c, err := h.courseService.Get(id)
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/courses")
	}

	data := struct {
		Course response.Course
	}{
		Course: response.CourseBuilder().BuildFromDomain(c),
	}

	return ctx.Render(http.StatusOK, "courses/edit.html", helpers.NewPageData(ctx, "Editar Curso", "courses", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Cursos", URL: "/courses"},
			helpers.Breadcrumb{Label: "Editar", URL: "/courses/" + id.String() + "/edit"},
		))
}

func (h *resourceViewHandlers) UpdateCourse(ctx handlers.RichContext) error {
	id, _ := uuid.Parse(ctx.Param("id"))
	name := ctx.FormValue("name")

	c, err := course.NewBuilder().WithID(id).WithName(name).Build()
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	err = h.courseService.Update(c)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	ctx.Response().Header().Set("HX-Redirect", "/courses")
	return ctx.NoContent(http.StatusOK)
}
