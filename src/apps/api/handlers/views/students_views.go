package views

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/handlers/dto/request"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/handlers/formData"
	"eletronic_point/src/apps/api/handlers/views/helpers"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/services/filters"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

type StudentViewHandlers interface {
	List(handlers.RichContext) error
	CreatePage(handlers.RichContext) error
	Create(handlers.RichContext) error
	EditPage(handlers.RichContext) error
	Update(handlers.RichContext) error
	Show(handlers.RichContext) error
}

type studentViewHandlers struct {
	service            primary.StudentPort
	institutionService primary.InstitutionPort
	campusService      primary.CampusPort
	courseService      primary.CoursePort
}

func NewStudentViewHandlers(
	service primary.StudentPort,
	institutionService primary.InstitutionPort,
	campusService primary.CampusPort,
	courseService primary.CoursePort,
) StudentViewHandlers {
	return &studentViewHandlers{service, institutionService, campusService, courseService}
}

func (h *studentViewHandlers) List(ctx handlers.RichContext) error {
	f := helpers.GetStudentFilters(ctx)
	students, err := h.service.List(f)
	if err != nil {
		return ctx.Render(http.StatusOK, "students/list.html", helpers.PageData{Errors: []string{err.String()}})
	}

	institutions, _ := h.institutionService.List(filters.InstitutionFilters{})
	campus, _ := h.campusService.List(filters.CampusFilters{})

	data := map[string]interface{}{
		"Students":     response.StudentBuilder().BuildFromDomainList(students),
		"Institutions": helpers.ToOptions(institutions),
		"Campus":       helpers.ToOptions(campus),
		"Filters":      f,
	}

	return ctx.Render(http.StatusOK, "students/list.html", helpers.NewPageData(ctx, "Estudantes", "students", data).
		WithBreadcrumbs(helpers.Breadcrumb{Label: "Estudantes", URL: "/students"}))
}

func (h *studentViewHandlers) CreatePage(ctx handlers.RichContext) error {
	campus, _ := h.campusService.List(filters.CampusFilters{})
	courses, _ := h.courseService.List(filters.CourseFilters{})

	data := map[string]interface{}{
		"Campus":  helpers.ToOptions(campus),
		"Courses": helpers.ToOptions(courses),
	}

	return ctx.Render(http.StatusOK, "students/create.html", helpers.NewPageData(ctx, "Novo Estudante", "students", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Estudantes", URL: "/students"},
			helpers.Breadcrumb{Label: "Novo", URL: "/students/new"},
		))
}

func (h *studentViewHandlers) Create(ctx handlers.RichContext) error {
	if err := ctx.Request().ParseMultipartForm(10 << 20); err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{"Erro ao processar formulário"}})
	}

	campusID, _ := uuid.Parse(ctx.FormValue(formData.StudentCampusID))
	courseID, _ := uuid.Parse(ctx.FormValue(formData.StudentCourseID))
	totalWorkload, _ := strconv.Atoi(ctx.FormValue(formData.StudentTotalWorkload))

	fileName, _ := helpers.SaveUploadedFile(ctx, formData.StudentProfilePicture)

	dto := request.Student{
		Name:           ctx.FormValue(formData.StudentName),
		BirthDate:      ctx.FormValue(formData.StudentBirthDate),
		CPF:            ctx.FormValue(formData.StudentCPF),
		Email:          ctx.FormValue(formData.StudentEmail),
		Phone:          ctx.FormValue(formData.StudentPhone),
		Registration:   ctx.FormValue(formData.StudentRegistration),
		ProfilePicture: fileName,
		CampusID:       campusID,
		CourseID:       courseID,
		TotalWorkload:  totalWorkload,
	}

	s, dErr := dto.ToDomain()
	if dErr != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{dErr.String()}})
	}

	s.SetResponsibleTeacherID(*ctx.ProfileID())

	_, err := h.service.Create(s)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	ctx.Response().Header().Set("HX-Redirect", "/students")
	return ctx.NoContent(http.StatusCreated)
}

func (h *studentViewHandlers) EditPage(ctx handlers.RichContext) error {
	id, _ := uuid.Parse(ctx.Param("id"))
	s, err := h.service.Get(id, filters.StudentFilters{})
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/students")
	}

	campus, _ := h.campusService.List(filters.CampusFilters{})
	courses, _ := h.courseService.List(filters.CourseFilters{})

	data := map[string]interface{}{
		"Student": response.StudentBuilder().BuildFromDomain(s),
		"Campus":  helpers.ToOptions(campus),
		"Courses": helpers.ToOptions(courses),
	}

	return ctx.Render(http.StatusOK, "students/edit.html", helpers.NewPageData(ctx, "Editar Estudante", "students", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Estudantes", URL: "/students"},
			helpers.Breadcrumb{Label: s.Name(), URL: "/students/" + id.String()},
			helpers.Breadcrumb{Label: "Editar", URL: "/students/" + id.String() + "/edit"},
		))
}

func (h *studentViewHandlers) Update(ctx handlers.RichContext) error {
	id, _ := uuid.Parse(ctx.Param("id"))
	if err := ctx.Request().ParseMultipartForm(10 << 20); err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{"Erro ao processar formulário"}})
	}

	campusID, _ := uuid.Parse(ctx.FormValue(formData.StudentCampusID))
	courseID, _ := uuid.Parse(ctx.FormValue(formData.StudentCourseID))
	totalWorkload, _ := strconv.Atoi(ctx.FormValue(formData.StudentTotalWorkload))

	fileName, _ := helpers.SaveUploadedFile(ctx, formData.StudentProfilePicture)

	dto := request.Student{
		Name:           ctx.FormValue(formData.StudentName),
		BirthDate:      ctx.FormValue(formData.StudentBirthDate),
		CPF:            ctx.FormValue(formData.StudentCPF),
		Email:          ctx.FormValue(formData.StudentEmail),
		Phone:          ctx.FormValue(formData.StudentPhone),
		Registration:   ctx.FormValue(formData.StudentRegistration),
		ProfilePicture: fileName,
		CampusID:       campusID,
		CourseID:       courseID,
		TotalWorkload:  totalWorkload,
	}

	s, dErr := dto.ToDomain()
	if dErr != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{dErr.String()}})
	}
	s.SetID(&id)

	err := h.service.Update(s)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	ctx.Response().Header().Set("HX-Redirect", "/students/"+id.String())
	return ctx.NoContent(http.StatusOK)
}

func (h *studentViewHandlers) Show(ctx handlers.RichContext) error {
	id, _ := uuid.Parse(ctx.Param("id"))
	s, err := h.service.Get(id, helpers.GetStudentFilters(ctx))
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/students")
	}

	data := map[string]interface{}{
		"Student": response.StudentBuilder().BuildFromDomain(s),
	}

	return ctx.Render(http.StatusOK, "students/show.html", helpers.NewPageData(ctx, s.Name(), "students", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Estudantes", URL: "/students"},
			helpers.Breadcrumb{Label: s.Name(), URL: "/students/" + id.String()},
		))
}
