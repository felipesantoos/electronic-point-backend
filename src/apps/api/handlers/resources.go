package handlers

import (
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/core/interfaces/primary"
	"net/http"
)

type ResourcesHandler interface {
	ListAccountRoles(RichContext) error
}

type resourcesHandler struct {
	usecase primary.ResourcesPort
}

func NewResourcesHandler(usecase primary.ResourcesPort) ResourcesHandler {
	return &resourcesHandler{usecase}
}

// List Account Roles
// @ID Resources.ListAccountRoles
// @Summary Listar todas as funções de conta existentes do banco de dados.
// @Description Pode ser utilizada para visualizar as funções de conta cadastradas no banco de dados e adquirir o
// @Description identificador da função desejada para a criação de uma nova conta.
// @Tags Recursos
// @Produce json
// @Success 200 {array} response.Role "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /res/account-roles [get]
func (h *resourcesHandler) ListAccountRoles(context RichContext) error {
	result, err := h.usecase.ListAccountRoles()
	if err != nil {
		return response.ErrorBuilder().NewFromDomain(err)
	}
	return context.JSON(http.StatusOK, response.AccountRoleBuilder().BuildFromDomainList(result))
}
