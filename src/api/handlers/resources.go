package handlers

import (
	"dit_backend/src/api/handlers/dto/response"
	"dit_backend/src/core/interfaces/usecases"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ResourcesHandler interface {
	ListAccountRoles(echo.Context) error
}

type resourcesHandler struct {
	usecase usecases.ResourcesUseCase
}

func NewResourcesHandler(usecase usecases.ResourcesUseCase) ResourcesHandler {
	return &resourcesHandler{usecase}
}

// List Account Roles
// @ID Resources.ListAccountRoles
// @Summary Listar todas as funções de conta existentes do banco de dados.
// @Description Pode ser utilizada para visualizar as funções de conta cadastradas no banco de dados e adquirir o
// @Description identificador da função desejada para a criação de uma nova conta.
// @Security	bearerAuth
// @Tags Resources
// @Produce json
// @Success 200 {array} response.Role "Requisição realizada com sucesso."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /res/account-roles [get]
func (instance *resourcesHandler) ListAccountRoles(context echo.Context) error {
	roles, err := instance.usecase.ListAccountRoles()
	if err != nil {
		return responseFromError(err)
	}
	var serializedRoles = []response.Role{}
	for _, role := range roles {
		serializedRoles = append(serializedRoles, response.AccountRoleBuilder().FromDomain(role))
	}
	return context.JSON(http.StatusOK, serializedRoles)
}
