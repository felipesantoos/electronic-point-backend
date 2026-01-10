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

type InternshipLocationHandlers interface {
	Create(RichContext) error
	Update(RichContext) error
	Delete(RichContext) error
	List(RichContext) error
	Get(RichContext) error
}

type internshipLocationHandlers struct {
	services primary.InternshipLocationPort
}

func NewInternshipLocationHandlers(services primary.InternshipLocationPort) InternshipLocationHandlers {
	return &internshipLocationHandlers{services}
}

// Create
// @ID InternshipLocation.Create
// @Summary Crie um novo local de estágio.
// @Description Cria um novo local de estágio no sistema com os dados fornecidos.
// @Tags Locais de estágio
// @Security BearerAuth
// @Accept application/json
// @Produce json
// @Param body body request.InternshipLocation true "Dados do local de estágio"
// @Success 201 {object} response.ID "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /internship-locations [post]
func (this *internshipLocationHandlers) Create(ctx RichContext) error {
	if ctx.RoleName() != role.TEACHER_ROLE_CODE {
		return forbiddenErrorWithMessage(messages.YouDoNotHaveAccessToThisResource)
	}
	var internshipLocationDTO request.InternshipLocation
	if err := ctx.Bind(&internshipLocationDTO); err != nil {
		logger.Error().Msg(err.Error())
		return badRequestErrorWithMessage(err.Error())
	}
	_internshipLocation, validationError := internshipLocationDTO.ToDomain()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return unprocessableEntityErrorWithMessage(validationError.String())
	}
	id, err := this.services.Create(_internshipLocation)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusCreated, response.IDBuilder().FromUUID(*id))
}

// Update
// @ID InternshipLocation.Update
// @Summary Atualizar um local de estágio existente.
// @Description Atualiza os dados de um local de estágio existente no sistema.
// @Tags Locais de estágio
// @Security BearerAuth
// @Accept application/json
// @Produce json
// @Param id path string true "ID do local de estágio" default(8c6b88c0-d123-45f6-9a10-1d8c5f7b9e75)
// @Param body body request.InternshipLocation true "Dados do local de estágio"
// @Success 204 {object} nil "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /internship-locations/{id} [put]
func (this *internshipLocationHandlers) Update(ctx RichContext) error {
	if ctx.RoleName() != role.TEACHER_ROLE_CODE {
		return forbiddenErrorWithMessage(messages.YouDoNotHaveAccessToThisResource)
	}
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		logger.Error().Msg(conversionError.Error())
		return badRequestErrorWithMessage(conversionError.Error())
	}
	var internshipLocationDTO request.InternshipLocation
	if err := ctx.Bind(&internshipLocationDTO); err != nil {
		logger.Error().Msg(err.Error())
		return badRequestErrorWithMessage(err.Error())
	}
	_internshipLocation, validationError := internshipLocationDTO.ToDomain()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return unprocessableEntityErrorWithMessage(validationError.String())
	}
	validationError = _internshipLocation.SetID(id)
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return unprocessableEntityErrorWithMessage(validationError.String())
	}
	err := this.services.Update(_internshipLocation)
	if err != nil {
		return responseFromError(err)
	}
	return successNoContent(ctx)
}

// Delete
// @ID InternshipLocation.Delete
// @Summary Excluir um local de estágio por ID.
// @Description Exclui o local de estágio especificado do sistema.
// @Tags Locais de estágio
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID do local de estágio" default(8c6b88c0-d123-45f6-9a10-1d8c5f7b9e75)
// @Success 204 {object} nil "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /internship-locations/{id} [delete]
func (this *internshipLocationHandlers) Delete(ctx RichContext) error {
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		logger.Error().Msg(conversionError.Error())
		return badRequestErrorWithMessage(conversionError.Error())
	}
	err := this.services.Delete(id)
	if err != nil {
		return responseFromError(err)
	}
	return successNoContent(ctx)
}

// List
// @ID InternshipLocation.List
// @Summary Listar todos os locais de estágio.
// @Description Recupera uma lista de todos os locais de estágio no sistema.
// @Tags Locais de estágio
// @Security BearerAuth
// @Produce json
// @Param studentID query string false "ID do estudante"
// @Success 200 {array} response.InternshipLocation "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /internship-locations [get]
func (this *internshipLocationHandlers) List(ctx RichContext) error {
	var studentID *uuid.UUID
	if !checkers.IsEmpty(ctx.QueryParam(params.StudentID)) {
		value, conversionError := getUUIDQueryParamValue(ctx, params.StudentID)
		if conversionError != nil {
			logger.Error().Msg(conversionError.String())
			return responseFromError(conversionError)
		}
		studentID = value
	}
	_filters := filters.InternshipLocationFilters{StudentID: studentID}
	internshipLocations, err := this.services.List(_filters)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.InternshipLocationBuilder().BuildFromDomainList(internshipLocations))
}

// Get
// @ID InternshipLocation.Get
// @Summary Recuperar um local de estágio por ID.
// @Description Recupera um local de estágio específico com base no ID fornecido.
// @Tags Locais de estágio
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID do local de estágio" default(8c6b88c0-d123-45f6-9a10-1d8c5f7b9e75)
// @Success 200 {object} response.InternshipLocation "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /internship-locations/{id} [get]
func (this *internshipLocationHandlers) Get(ctx RichContext) error {
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		logger.Error().Msg(conversionError.Error())
		return badRequestErrorWithMessage(conversionError.Error())
	}
	_internshipLocation, err := this.services.Get(id)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.InternshipLocationBuilder().BuildFromDomain(_internshipLocation))
}
