package views

import (
	"eletronic_point/src/apps/api/handlers"
	"eletronic_point/src/apps/api/handlers/dto/request"
	"eletronic_point/src/apps/api/handlers/views/helpers"
	"eletronic_point/src/core/interfaces/primary"
	"net/http"
)

type AuthViewHandlers interface {
	LoginPage(handlers.RichContext) error
	Login(handlers.RichContext) error
	Logout(handlers.RichContext) error
	ResetPasswordPage(handlers.RichContext) error
	AskPasswordResetMail(handlers.RichContext) error
	ResetPasswordConfirmPage(handlers.RichContext) error
	UpdatePasswordByPasswordReset(handlers.RichContext) error
}

type authViewHandlers struct {
	service primary.AuthPort
}

func NewAuthViewHandlers(service primary.AuthPort) AuthViewHandlers {
	return &authViewHandlers{service}
}

func (h *authViewHandlers) LoginPage(ctx handlers.RichContext) error {
	if ctx.AccountID() != nil {
		return ctx.Redirect(http.StatusFound, "/")
	}
	return ctx.Render(http.StatusOK, "auth/login", helpers.NewPageData(ctx, "Login", "", nil))
}

func (h *authViewHandlers) Login(ctx handlers.RichContext) error {
	var body = struct {
		Email    string `form:"email"`
		Password string `form:"password"`
	}{}
	if err := ctx.Bind(&body); err != nil {
		return ctx.Render(http.StatusBadRequest, "auth/login", helpers.PageData{
			Title:  "Login",
			Errors: []string{"Dados inválidos"},
		})
	}

	credentials := request.Credentials{Email: body.Email, Password: body.Password}
	authorization, err := h.service.Login(credentials.ToDomain())
	if err != nil {
		return ctx.Render(http.StatusUnauthorized, "auth/login", helpers.PageData{
			Title:  "Login",
			Errors: []string{"Email ou senha incorretos"},
		})
	}

	ctx.SetCookie(helpers.PrepareTokenCookie(authorization))

	if ctx.Request().Header.Get("HX-Request") == "true" {
		ctx.Response().Header().Set("HX-Redirect", "/")
		return ctx.NoContent(http.StatusOK)
	}

	return ctx.Redirect(http.StatusFound, "/")
}

func (h *authViewHandlers) Logout(ctx handlers.RichContext) error {
	if ctx.AccountID() != nil {
		h.service.Logout(ctx.AccountID())
	}
	ctx.SetCookie(helpers.PrepareLogoutCookie())

	if ctx.Request().Header.Get("HX-Request") == "true" {
		ctx.Response().Header().Set("HX-Redirect", "/login")
		return ctx.NoContent(http.StatusOK)
	}

	return ctx.Redirect(http.StatusFound, "/login")
}

func (h *authViewHandlers) ResetPasswordPage(ctx handlers.RichContext) error {
	return ctx.Render(http.StatusOK, "auth/reset-password", helpers.NewPageData(ctx, "Recuperar Senha", "", nil))
}

func (h *authViewHandlers) AskPasswordResetMail(ctx handlers.RichContext) error {
	var body = struct {
		Email string `form:"email"`
	}{}
	if err := ctx.Bind(&body); err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{
			Errors: []string{"E-mail inválido"},
		})
	}

	err := h.service.AskPasswordResetMail(body.Email)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{
			Errors: []string{err.String()},
		})
	}

	return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{
		FlashMessages: []helpers.FlashMessage{
			{Type: "success", Content: "E-mail de recuperação enviado com sucesso!"},
		},
	})
}

func (h *authViewHandlers) ResetPasswordConfirmPage(ctx handlers.RichContext) error {
	token := ctx.Param("token")
	if err := h.service.FindPasswordResetByToken(token); err != nil {
		return ctx.Redirect(http.StatusFound, "/login")
	}
	return ctx.Render(http.StatusOK, "auth/reset-password-confirm", helpers.NewPageData(ctx, "Nova Senha", "", map[string]string{"Token": token}))
}

func (h *authViewHandlers) UpdatePasswordByPasswordReset(ctx handlers.RichContext) error {
	token := ctx.Param("token")
	var body = struct {
		NewPassword     string `form:"new_password"`
		ConfirmPassword string `form:"confirm_password"`
	}{}
	if err := ctx.Bind(&body); err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{
			Errors: []string{"Dados inválidos"},
		})
	}

	if body.NewPassword != body.ConfirmPassword {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{
			Errors: []string{"As senhas não coincidem"},
		})
	}

	err := h.service.UpdatePasswordByPasswordReset(token, body.NewPassword)
	if err != nil {
		return ctx.Render(http.StatusOK, "components/alerts", helpers.PageData{
			Errors: []string{err.String()},
		})
	}

	if ctx.Request().Header.Get("HX-Request") == "true" {
		ctx.Response().Header().Set("HX-Redirect", "/login")
		return ctx.NoContent(http.StatusOK)
	}

	return ctx.Redirect(http.StatusFound, "/login")
}
