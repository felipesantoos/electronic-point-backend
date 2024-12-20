package handlers

import (
	"eletronic_point/src/apps/api/handlers/checkers"
	"eletronic_point/src/apps/api/handlers/dto/request"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/handlers/params"
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/messages"
	"eletronic_point/src/core/services/filters"
	"net/http"

	"github.com/google/uuid"
)

type InternshipHandlers interface {
	Create(RichContext) error
	Update(RichContext) error
	Delete(RichContext) error
	List(RichContext) error
	Get(RichContext) error
}

type internshipHandlers struct {
	services primary.InternshipPort
}

func NewInternshipHandlers(services primary.InternshipPort) InternshipHandlers {
	return &internshipHandlers{services}
}

// Create
// @ID Internship.Create
// @Summary Crie um novo estágio.
// @Description Cria um novo estágio no sistema com os dados fornecidos.
// @Tags Estágios
// @Security BearerAuth
// @Accept application/json
// @Produce json
// @Param body body request.Internship true "Dados de estágio"
// @Success 201 {object} response.ID "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /internships [post]
func (this *internshipHandlers) Create(ctx RichContext) error {
	if ctx.RoleName() != role.TEACHER_ROLE_CODE {
		return unauthorizedErrorWithMessage(messages.YouDoNotHaveAccessToThisResource)
	}
	var internshipDTO request.Internship
	if err := ctx.Bind(&internshipDTO); err != nil {
		logger.Error().Msg(err.Error())
		return badRequestErrorWithMessage(err.Error())
	}
	_internship, validationError := internshipDTO.ToDomain()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return unprocessableEntityErrorWithMessage(validationError.String())
	}
	id, err := this.services.Create(_internship)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusCreated, response.IDBuilder().FromUUID(*id))
}

// Update
// @ID Internship.Update
// @Summary Atualizar um estágio existente.
// @Description Atualiza os dados de um estágio existente no sistema.
// @Tags Estágios
// @Security BearerAuth
// @Accept application/json
// @Produce json
// @Param id path string true "ID do estágio" default(9ec95529-06d5-47e2-b617-1606088ac9e6)
// @Param body body request.Internship true "Dados de estágio"
// @Success 204 {object} nil "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /internships/{id} [put]
func (this *internshipHandlers) Update(ctx RichContext) error {
	if ctx.RoleName() != role.TEACHER_ROLE_CODE {
		return unauthorizedErrorWithMessage(messages.YouDoNotHaveAccessToThisResource)
	}
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		logger.Error().Msg(conversionError.Error())
		return badRequestErrorWithMessage(conversionError.Error())
	}
	var internshipDTO request.Internship
	if err := ctx.Bind(&internshipDTO); err != nil {
		logger.Error().Msg(err.Error())
		return badRequestErrorWithMessage(err.Error())
	}
	_internship, validationError := internshipDTO.ToDomain()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return unprocessableEntityErrorWithMessage(validationError.String())
	}
	_internship.SetID(id)
	err := this.services.Update(_internship)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.NoContent(http.StatusNoContent)
}

// Delete
// @ID Internship.Delete
// @Summary Excluir um estágio por ID.
// @Description Exclui o estágio especificado do sistema.
// @Tags Estágios
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID do estágio" default(9ec95529-06d5-47e2-b617-1606088ac9e6)
// @Success 204 {object} nil "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /internships/{id} [delete]
func (this *internshipHandlers) Delete(ctx RichContext) error {
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		logger.Error().Msg(conversionError.Error())
		return badRequestErrorWithMessage(conversionError.Error())
	}
	err := this.services.Delete(id)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.NoContent(http.StatusNoContent)
}

// List
// @ID Internship.List
// @Summary Listar todos os estágios.
// @Description Recupera uma lista de todos os estágios no sistema.
// @Tags Estágios
// @Security BearerAuth
// @Produce json
// @Param studentID query string false "ID do estudante"
// @Success 200 {array} response.Internship "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /internships [get]
func (this *internshipHandlers) List(ctx RichContext) error {
	var studentID *uuid.UUID
	if !checkers.IsEmpty(ctx.QueryParam(params.StudentID)) {
		value, conversionError := getUUIDQueryParamValue(ctx, params.StudentID)
		if conversionError != nil {
			logger.Error().Msg(conversionError.String())
			return badRequestErrorWithMessage(conversionError.String())
		}
		studentID = value
	}
	_filters := filters.InternshipFilters{StudentID: studentID}
	result, err := this.services.List(_filters)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.InternshipBuilder().BuildFromDomainList(result))
}

// Get
// @ID Internship.Get
// @Summary Recuperar um estágio.
// @Description Recupera um estágio específico pelo ID.
// @Tags Estágios
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID do estágio" default(9ec95529-06d5-47e2-b617-1606088ac9e6)
// @Success 200 {object} response.Internship "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /internships/{id} [get]
func (this *internshipHandlers) Get(ctx RichContext) error {
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		logger.Error().Msg(conversionError.Error())
		return badRequestErrorWithMessage(conversionError.Error())
	}
	result, err := this.services.Get(id)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.InternshipBuilder().BuildFromDomain(result))
}
