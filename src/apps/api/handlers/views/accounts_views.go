package views

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/handlers/dto/request"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/handlers/views/helpers"
	"eletronic_point/src/core/domain/person"
	updatepassword "eletronic_point/src/core/domain/updatePassword"
	"eletronic_point/src/core/interfaces/primary"
	"net/http"

	"github.com/google/uuid"
)

type AccountViewHandlers interface {
	List(handlers.RichContext) error
	CreatePage(handlers.RichContext) error
	Create(handlers.RichContext) error
	Show(handlers.RichContext) error
	EditPage(handlers.RichContext) error
	Update(handlers.RichContext) error
	ProfilePage(handlers.RichContext) error
	UpdateProfile(handlers.RichContext) error
	UpdatePassword(handlers.RichContext) error
}

type accountViewHandlers struct {
	service          primary.AccountPort
	resourcesService primary.ResourcesPort
}

func NewAccountViewHandlers(service primary.AccountPort, resourcesService primary.ResourcesPort) AccountViewHandlers {
	return &accountViewHandlers{service, resourcesService}
}

func (h *accountViewHandlers) List(ctx handlers.RichContext) error {
	if !ctx.IsAdmin() {
		return ctx.Redirect(http.StatusFound, "/login")
	}

	f := helpers.GetAccountFilters(ctx)
	accounts, err := h.service.List(f)
	if err != nil {
		return ctx.Render(http.StatusOK, "accounts/list.html", helpers.PageData{
			Title:  "Contas",
			Errors: []string{err.String()},
		})
	}

	roles, _ := h.resourcesService.ListAccountRoles()

	data := map[string]interface{}{
		"Accounts": response.AccountBuilder().BuildFromDomainList(accounts),
		"Roles":    helpers.ToOptions(roles),
		"Filters":  f,
	}

	return ctx.Render(http.StatusOK, "accounts/list.html", helpers.NewPageData(ctx, "Contas", "accounts", data).
		WithBreadcrumbs(helpers.Breadcrumb{Label: "Contas", URL: "/admin/accounts"}))
}

func (h *accountViewHandlers) CreatePage(ctx handlers.RichContext) error {
	if !ctx.IsAdmin() {
		return ctx.Redirect(http.StatusFound, "/")
	}

	roles, _ := h.resourcesService.ListAccountRoles()

	data := map[string]interface{}{
		"Roles": helpers.ToOptions(roles),
	}

	return ctx.Render(http.StatusOK, "accounts/create.html", helpers.NewPageData(ctx, "Nova Conta", "accounts", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Contas", URL: "/admin/accounts"},
			helpers.Breadcrumb{Label: "Nova", URL: "/admin/accounts/new"},
		))
}

func (h *accountViewHandlers) Create(ctx handlers.RichContext) error {
	if !ctx.IsAdmin() {
		return ctx.NoContent(http.StatusForbidden)
	}

	var body request.CreateAccount
	if err := ctx.Bind(&body); err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{"Dados inválidos"}})
	}

	// Manual conversion to domain for simplicity here, or use request DTO
	acc, dErr := body.ToDomain()
	if dErr != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{dErr.String()}})
	}

	_, err := h.service.Create(acc)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	ctx.Response().Header().Set("HX-Redirect", "/admin/accounts")
	return ctx.NoContent(http.StatusCreated)
}

func (h *accountViewHandlers) Show(ctx handlers.RichContext) error {
	if !ctx.IsAdmin() {
		return ctx.Redirect(http.StatusFound, "/")
	}

	id, _ := uuid.Parse(ctx.Param("id"))
	acc, err := h.service.FindByID(&id)
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/admin/accounts")
	}

	data := map[string]interface{}{
		"Account": response.AccountBuilder().BuildFromDomain(acc),
	}

	return ctx.Render(http.StatusOK, "accounts/show.html", helpers.NewPageData(ctx, acc.Person().Name(), "accounts", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Contas", URL: "/admin/accounts"},
			helpers.Breadcrumb{Label: acc.Person().Name(), URL: "/admin/accounts/" + id.String()},
		))
}

func (h *accountViewHandlers) EditPage(ctx handlers.RichContext) error {
	if !ctx.IsAdmin() {
		return ctx.Redirect(http.StatusFound, "/")
	}

	id, _ := uuid.Parse(ctx.Param("id"))
	acc, err := h.service.FindByID(&id)
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/admin/accounts")
	}

	roles, _ := h.resourcesService.ListAccountRoles()

	data := map[string]interface{}{
		"Account": response.AccountBuilder().BuildFromDomain(acc),
		"Roles":   helpers.ToOptions(roles),
	}

	return ctx.Render(http.StatusOK, "accounts/edit.html", helpers.NewPageData(ctx, "Editar Conta", "accounts", data).
		WithBreadcrumbs(
			helpers.Breadcrumb{Label: "Contas", URL: "/admin/accounts"},
			helpers.Breadcrumb{Label: acc.Person().Name(), URL: "/admin/accounts/" + id.String()},
			helpers.Breadcrumb{Label: "Editar", URL: "/admin/accounts/" + id.String() + "/edit"},
		))
}

func (h *accountViewHandlers) Update(ctx handlers.RichContext) error {
	if !ctx.IsAdmin() {
		return ctx.NoContent(http.StatusForbidden)
	}

	id, _ := uuid.Parse(ctx.Param("id"))
	acc, err := h.service.FindByID(&id)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{"Conta não encontrada"}})
	}

	var body struct {
		Email  string `form:"email"`
		RoleID string `form:"role_id"`
	}
	if err := ctx.Bind(&body); err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{"Dados inválidos"}})
	}

	roleUUID, _ := uuid.Parse(body.RoleID)
	allRoles, _ := h.resourcesService.ListAccountRoles()
	for _, r := range allRoles {
		if r.ID() != nil && *r.ID() == roleUUID {
			acc.SetRole(r)
			break
		}
	}

	acc.SetEmail(body.Email)

	if err := h.service.Update(acc); err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	ctx.Response().Header().Set("HX-Redirect", "/admin/accounts/"+id.String())
	return ctx.NoContent(http.StatusOK)
}

func (h *accountViewHandlers) ProfilePage(ctx handlers.RichContext) error {
	acc, err := h.service.FindByID(ctx.AccountID())
	if err != nil {
		return ctx.Redirect(http.StatusFound, "/login")
	}

	data := struct {
		Profile response.Account
	}{
		Profile: response.AccountBuilder().BuildFromDomain(acc),
	}

	return ctx.Render(http.StatusOK, "accounts/profile.html", helpers.NewPageData(ctx, "Meu Perfil", "profile", data))
}

func (h *accountViewHandlers) UpdateProfile(ctx handlers.RichContext) error {
	var body struct {
		Name      string `form:"name"`
		CPF       string `form:"cpf"`
		Phone     string `form:"phone"`
		BirthDate string `form:"birth_date"`
	}
	if err := ctx.Bind(&body); err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{"Dados inválidos"}})
	}

	p, dErr := person.NewBuilder().
		WithID(*ctx.ProfileID()).
		WithName(body.Name).
		WithCPF(body.CPF).
		WithPhone(body.Phone).
		WithBirthDate(body.BirthDate).
		Build()
	
	if dErr != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{dErr.String()}})
	}

	err := h.service.UpdateAccountProfile(p)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{
		FlashMessages: []helpers.FlashMessage{{Type: "success", Content: "Perfil atualizado com sucesso!"}},
	})
}

func (h *accountViewHandlers) UpdatePassword(ctx handlers.RichContext) error {
	var body struct {
		CurrentPassword string `form:"current_password"`
		NewPassword     string `form:"new_password"`
		ConfirmPassword string `form:"confirm_password"`
	}
	if err := ctx.Bind(&body); err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{"Dados inválidos"}})
	}

	if body.NewPassword != body.ConfirmPassword {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{"As senhas não coincidem"}})
	}

	data := updatepassword.New(body.CurrentPassword, body.NewPassword)
	err := h.service.UpdateAccountPassword(ctx.AccountID(), data)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{Errors: []string{err.String()}})
	}

	return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{
		FlashMessages: []helpers.FlashMessage{{Type: "success", Content: "Senha atualizada com sucesso!"}},
	})
}
