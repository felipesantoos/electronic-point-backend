package handlers

import (
	"backend_template/src/core/interfaces/usecases"
	"backend_template/src/ui/api/handlers/dto/request"
	"backend_template/src/ui/api/handlers/dto/response"
	"net/http"

	"github.com/wallrony/go-validator/validator"
)

type AccountHandler interface {
	List(EnhancedContext) error
	FindProfile(EnhancedContext) error
	Create(EnhancedContext) error
	UpdatePassword(EnhancedContext) error
	UpdateProfile(EnhancedContext) error
}

type accountHandler struct {
	service usecases.AccountUseCase
}

func NewAccountHandler(service usecases.AccountUseCase) AccountHandler {
	return &accountHandler{service}
}

// List
// @ID Accounts.List
// @Summary Listar todas as contas existentes do banco de dados.
// @Description Esta rota retorna todas as informações de todas as contas cadastradas no banco de dados.
// @Description Dados como "professional" irão somente aparecer caso a role da conta for própria para contenção desses.
// @Security	bearerAuth
// @Tags Administrador
// @Produce json
// @Success 200 {array} response.Account "Requisição realizada com sucesso."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /admin/accounts [get]
func (h *accountHandler) List(context EnhancedContext) error {
	accounts, err := h.service.List()
	if err != nil {
		return responseFromError(err)
	}
	serializedAccounts := []response.Account{}
	for _, account := range accounts {
		serializedAccounts = append(serializedAccounts, response.AccountBuilder().BuildFromDomain(account))
	}
	return context.JSON(http.StatusOK, serializedAccounts)
}

// FindProfile
// @ID Accounts.FindProfile
// @Summary Pesquisar dados do perfil de uma conta.
// @Description Esta rota retorna todas as informações de todas as contas cadastradas no banco de dados.
// @Description Dados como "professional" irão somente aparecer caso a role da conta for própria para contenção desses.
// @Security	bearerAuth
// @Tags Geral
// @Produce json
// @Success 200 {object} response.Account "Requisição realizada com sucesso."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /accounts/profile [get]
func (h *accountHandler) FindProfile(context EnhancedContext) error {
	account, err := h.service.FindByID(context.AccountID())
	if err != nil {
		return responseFromError(err)
	}
	return context.JSON(http.StatusOK, response.AccountBuilder().BuildFromDomain(account))
}

// Create
// @ID Accounts.Create
// @Summary Cadastrar uma nova conta de usuário
// @Description Ao enviar dados para cadastro de uma nova conta, os dados relacionados à "Profissional"
// @Description são facultativos, tendo somente que enviar os dados que são relacionados à role definida.
// @Security	bearerAuth
// @Accept json
// @Param json body request.CreateAccount true "JSON com todos os dados necessários para o cadastro de uma conta de usuário."
// @Tags Administrador
// @Produce json
// @Success 201 {object} response.ID "Requisição realizada com sucesso."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /admin/accounts [post]
func (h *accountHandler) Create(context EnhancedContext) error {
	var body interface{}
	if err := context.Bind(&body); err != nil {
		return unsupportedMediaTypeError
	}
	if dto, err := validator.ValidateDTO[request.CreateAccount](body); err != nil {
		return responseFromValidationError((err))
	} else if data, err := dto.ToDomain(); err != nil {
		return responseFromError(err)
	} else if id, err := h.service.Create(data); err != nil {
		return responseFromError(err)
	} else {
		return context.JSON(http.StatusCreated, map[string]interface{}{
			"id": id.String(),
		})
	}
}

// UpdateProfile
// @ID Account.UpdateProfile
// @Summary Atualizar dados do perfil de uma conta.
// @Security	bearerAuth
// @Accept json
// @Tags Geral
// @Param json  body request.UpdateAccountProfile true "JSON com todos os dados necessários para o processo de atualização de dados do perfil."
// @Success 200
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /accounts/profile [put]
func (h *accountHandler) UpdateProfile(context EnhancedContext) error {
	var body interface{}
	if err := context.Bind(&body); err != nil {
		return unsupportedMediaTypeError
	}
	data, vErr := validator.ValidateDTO[request.UpdateAccountProfile](body)
	if vErr != nil {
		return responseFromValidationError((vErr))
	}
	profile, err := data.ToDomain()
	if err != nil {
		return responseFromError(err)
	}
	profile.SetID(context.ProfileID())
	if err := h.service.UpdateAccountProfile(profile); err != nil {
		return responseFromError(err)
	}
	return context.NoContent(http.StatusOK)
}

// UpdateAccountPassword
// @ID Account.UpdateAccountPassword
// @Summary Realizar a atualização de senha de uma conta.
// @Security	bearerAuth
// @Accept json
// @Tags Geral
// @Param json  body request.UpdatePassword true "JSON com todos os dados necessários para a atualização da senha da conta."
// @Success 200
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /accounts/update-password [put]
func (h *accountHandler) UpdatePassword(context EnhancedContext) error {
	var body = map[string]interface{}{}
	if bindErr := context.Bind(&body); bindErr != nil {
		return unsupportedMediaTypeError
	}
	data, vErr := validator.ValidateDTO[request.UpdatePassword](body)
	if vErr != nil {
		return responseFromValidationError((vErr))
	}
	err := h.service.UpdateAccountPassword(context.AccountID(), data.ToDomain())
	if err != nil {
		return responseFromError(err)
	}
	return context.NoContent(http.StatusOK)
}
