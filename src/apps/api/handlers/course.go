package handlers

import (
	"eletronic_point/src/apps/api/handlers/checkers"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/handlers/params"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/services/filters"
	"net/http"
)

type CourseHandlers interface {
	List(RichContext) error
}

type courseHandlers struct {
	services primary.CoursePort
}

func NewCourseHandlers(services primary.CoursePort) CourseHandlers {
	return &courseHandlers{services}
}

// List
// @ID Course.List
// @Summary Listar todos os cursos.
// @Description Recupera uma lista de todos os cursos no sistema.
// @Tags Cursos
// @Security BearerAuth
// @Produce json
// @Param name query string false "Nome do curso"
// @Success 200 {array} response.Course "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /courses [get]
func (this *courseHandlers) List(ctx RichContext) error {
	var name *string
	if !checkers.IsEmpty(ctx.QueryParam(params.Name)) {
		value := ctx.QueryParam(params.Name)
		name = &value
	}
	_filters := filters.CourseFilters{Name: name}
	courses, err := this.services.List(_filters)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.CourseBuilder().BuildFromDomainList(courses))
}
