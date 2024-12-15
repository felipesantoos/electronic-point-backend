package handlers

import (
	"eletronic_point/src/apps/api/handlers/dto/request"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/utils"
	"eletronic_point/src/core/domain/authorization"
	"eletronic_point/src/core/interfaces/primary"
	"encoding/hex"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wallrony/go-validator/validator"
)

type AuthHandler interface {
	Login(RichContext) error
	Logout(RichContext) error
	AskPasswordResetMail(RichContext) error
	FindPasswordResetByToken(RichContext) error
	UpdatePasswordByPasswordReset(RichContext) error
}

type authHandler struct {
	service primary.AuthPort
}

func NewAuthHandler(service primary.AuthPort) AuthHandler {
	return &authHandler{service}
}

// Login
// @ID Auth.Login
// @Summary Adquirir autorização de acesso aos recursos da API através de credenciais de uma conta.
// @Description | E-mail         | Senha  | Função    |
// @Description |----------------|--------|-----------|
// @Description | jose@gmail.com | 123456 | Professor |
// @Description | ana@gmail.com  | 123456 | Esudante  |
// @Accept json
// @Param json body request.Credentials true "JSON com todos os dados necessários para o processo de autenticação."
// @Tags Anônimo
// @Produce json
// @Success 201 {object} response.Authorization "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /auth/login [post]
func (h *authHandler) Login(context RichContext) error {
	if context.AccountID() != nil {
		return context.NoContent(http.StatusOK)
	}
	var body map[string]interface{}
	if bindErr := context.Bind(&body); bindErr != nil {
		return response.ErrorBuilder().NewUnsupportedMediaTypeError()
	}
	dto, vErr := validator.ValidateDTO[request.Credentials](body)
	if vErr != nil {
		return response.ErrorBuilder().NewFromValidationError(vErr)
	}
	authorization, err := h.service.Login(dto.ToDomain())
	if err != nil {
		return response.ErrorBuilder().NewFromDomain(err)
	}
	tokenCookie := h.prepareTokenCookie(authorization)
	context.SetCookie(tokenCookie)
	return context.NoContent(http.StatusCreated)
}

// Logout
// @ID Auth.Logout
// @Summary Remove a sessão do registro de sessões permitidas.
// @Tags Geral
// @Produce json
// @Success 204 "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /auth/logout [post]
func (h *authHandler) Logout(context RichContext) error {
	err := h.service.Logout(context.AccountID())
	if err != nil {
		return response.ErrorBuilder().NewFromDomain(err)
	}
	return context.NoContent(http.StatusNoContent)
}

// PasswordReset
// @ID Auth.PasswordReset
// @Summary Solicitar email com token para atualização de senha.
// @Description cadastra uma nova entrada para a entidade `password_reset` vinculada à conta da sessão
// @Description e envia um e-mail para o email dessa.
// @Accept json
// @Param json body request.CreatePasswordReset true "JSON com todos os dados necessários para resetar a senha por email."
// @Tags Anônimo
// @Produce json
// @Success 201
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /auth/reset-password [post]
func (h *authHandler) AskPasswordResetMail(context RichContext) error {
	var body map[string]interface{}
	if bindErr := context.Bind(&body); bindErr != nil {
		return response.ErrorBuilder().NewUnsupportedMediaTypeError()
	}
	dto, err := validator.ValidateDTO[request.CreatePasswordReset](body)
	if err != nil {
		return response.ErrorBuilder().NewFromValidationError(err)
	}
	if err := h.service.AskPasswordResetMail(dto.Email); err != nil {
		return response.ErrorBuilder().NewFromDomain(err)
	}
	return context.NoContent(http.StatusCreated)
}

// FindPasswordResetByToken
// @ID Auth.FindPasswordResetByToken
// @Summary Verificar a existência de uma solicitação de atualização de senha por token.
// @Accept json
// @Tags Anônimo
// @Param token path string true "Token recebido pelo email da conta do usuário da plataforma."
// @Produce json
// @Success 200
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /auth/reset-password/{token} [get]
func (h *authHandler) FindPasswordResetByToken(context RichContext) error {
	if token, err := h.getPasswordResetToken(context); err != nil {
		return err
	} else if err := h.service.FindPasswordResetByToken(token); err != nil {
		return response.ErrorBuilder().NewFromDomain(err)
	}
	return context.NoContent(http.StatusOK)
}

// UpdatePasswordByPasswordReset
// @ID Auth.UpdatePasswordByPasswordReset
// @Summary Atualizar a senha de uma conta a partir de um token de atualização de senha.
// @Accept json
// @Tags Anônimo
// @Param token path string true "Token recebido pelo email da conta do usuário da plataforma."
// @Param json body request.UpdatePasswordByPasswordReset true "JSON com todos os dados necessários para resetar a senha por email."
// @Produce json
// @Success 200
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /auth/reset-password/{token} [put]
func (h *authHandler) UpdatePasswordByPasswordReset(context RichContext) error {
	token, err := h.getPasswordResetToken(context)
	if err != nil {
		return err
	}
	var body = map[string]interface{}{}
	if bindErr := context.Bind(&body); bindErr != nil {
		return context.NoContent(http.StatusUnsupportedMediaType)
	}
	dto, vErr := validator.ValidateDTO[request.UpdatePasswordByPasswordReset](body)
	if vErr != nil {
		return response.ErrorBuilder().NewFromValidationError(vErr)
	}
	if err := h.service.UpdatePasswordByPasswordReset(token, dto.NewPassword); err != nil {
		return response.ErrorBuilder().NewFromDomain(err)
	}
	return context.NoContent(http.StatusOK)
}

func (h *authHandler) getPasswordResetToken(context RichContext) (string, error) {
	token := context.Param("token")
	if _, err := hex.DecodeString(token); err != nil {
		return "", &echo.HTTPError{
			Code:    400,
			Message: "the provided token is invalid!",
		}
	}
	return token, nil
}

func (h *authHandler) prepareTokenCookie(auth authorization.Authorization) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = COOKIE_TOKEN_NAME
	cookie.Value = auth.Token()
	cookie.Path = "/"
	cookie.HttpOnly = false
	cookie.Secure = utils.IsAPIInProdMode()
	cookie.Expires = *auth.ExpirationTime()
	return cookie
}
