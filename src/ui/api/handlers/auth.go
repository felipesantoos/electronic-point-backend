package handlers

import (
	"backend_template/src/core/interfaces/usecases"
	"backend_template/src/ui/api/handlers/dto/request"
	"backend_template/src/ui/api/handlers/dto/response"
	"encoding/hex"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/wallrony/go-validator/validator"
)

type AuthHandler interface {
	Login(echo.Context) error
	Logout(echo.Context) error
	AskPasswordResetMail(echo.Context) error
	FindPasswordResetByToken(echo.Context) error
	UpdatePasswordByPasswordReset(echo.Context) error
}

type authHandler struct {
	service usecases.AuthUseCase
}

func NewAuthHandler(service usecases.AuthUseCase) AuthHandler {
	return &authHandler{service}
}

// Login
// @ID Auth.Login
// @Summary Adquirir autorização de acesso aos recursos da API através de credenciais de uma conta.
// @Accept json
// @Param json body request.Credentials true "JSON com todos os dados necessários para o processo de autenticação."
// @Tags Anônimo
// @Produce json
// @Success 200 {object} response.Authorization "Requisição realizada com sucesso."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /auth/login [post]
func (h *authHandler) Login(context echo.Context) error {
	var body map[string]interface{}
	if bindErr := context.Bind(&body); bindErr != nil {
		return unsupportedMediaTypeError
	}
	dto, vErr := validator.ValidateDTO[request.Credentials](body)
	if vErr != nil {
		return responseFromValidationError(vErr)
	}
	authorization, err := h.service.Login(dto.ToDomain())
	if err != nil {
		return responseFromError(err)
	}
	return context.JSON(http.StatusOK, response.NewAuthorizationBuilder().BuildFromDomain(authorization))
}

// Logout
// @ID Auth.Logout
// @Summary Remove a sessão do registro de sessões permitidas.
// @Tags Geral
// @Produce json
// @Success 204 "Requisição realizada com sucesso."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /auth/logout [post]
func (h *authHandler) Logout(context echo.Context) error {
	accountId, err := getAccountIDFromAuthorization(context)
	if err != nil {
		return responseFromError(err)
	}
	err = h.service.Logout(*accountId)
	if err != nil {
		return responseFromError(err)
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
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /auth/reset-password [post]
func (h *authHandler) AskPasswordResetMail(context echo.Context) error {
	var body map[string]interface{}
	if bindErr := context.Bind(&body); bindErr != nil {
		return context.NoContent(http.StatusUnsupportedMediaType)
	}
	dto, err := validator.ValidateDTO[request.CreatePasswordReset](body)
	if err != nil {
		return responseFromValidationError(err)
	}
	if err := h.service.AskPasswordResetMail(dto.Email); err != nil {
		return responseFromError(err)
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
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /auth/reset-password/{token} [get]
func (h *authHandler) FindPasswordResetByToken(context echo.Context) error {
	if token, err := h.getPasswordResetToken(context); err != nil {
		return err
	} else if err := h.service.FindPasswordResetByToken(token); err != nil {
		return responseFromError(err)
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
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /auth/reset-password/{token} [put]
func (h *authHandler) UpdatePasswordByPasswordReset(context echo.Context) error {
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
		return responseFromValidationError(vErr)
	}
	if err := h.service.UpdatePasswordByPasswordReset(token, dto.NewPassword); err != nil {
		return responseFromError(err)
	}
	return context.NoContent(http.StatusOK)
}

func (h *authHandler) getPasswordResetToken(context echo.Context) (string, error) {
	token := context.Param("token")
	if _, err := hex.DecodeString(token); err != nil {
		return "", &echo.HTTPError{
			Code:    400,
			Message: "the provided token is invalid!",
		}
	}
	return token, nil
}
