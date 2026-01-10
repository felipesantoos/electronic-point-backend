package handlers

import (
	"eletronic_point/src/apps/api/handlers/checkers"
	"eletronic_point/src/apps/api/handlers/dto/request"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/handlers/params"
	"eletronic_point/src/core/domain/campus"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/services/filters"
	"net/http"

	"github.com/google/uuid"
)

type CampusHandlers interface {
	List(RichContext) error
	Get(RichContext) error
	Create(RichContext) error
	Update(RichContext) error
	Delete(RichContext) error
}

type campusHandlers struct {
	services primary.CampusPort
}

func NewCampusHandlers(services primary.CampusPort) CampusHandlers {
	return &campusHandlers{services}
}

// List
// @ID Campus.List
// @Summary Listar todos os campi.
// @Description Recupera uma lista de todos os campi no sistema.
// @Tags Campi
// @Security BearerAuth
// @Produce json
// @Param name query string false "Nome do campus"
// @Success 200 {array} response.Campus "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /campus [get]
func (this *campusHandlers) List(ctx RichContext) error {
	var name *string
	if !checkers.IsEmpty(ctx.QueryParam(params.Name)) {
		value := ctx.QueryParam(params.Name)
		name = &value
	}
	_filters := filters.CampusFilters{Name: name}
	campuss, err := this.services.List(_filters)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.CampusBuilder().BuildFromDomainList(campuss))
}

// Get
func (this *campusHandlers) Get(ctx RichContext) error {
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		return badRequestErrorWithMessage(conversionError.Error())
	}
	_campus, err := this.services.Get(id)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.CampusBuilder().BuildFromDomain(_campus))
}

// Create
func (this *campusHandlers) Create(ctx RichContext) error {
	var campusDTO request.Campus
	if err := ctx.Bind(&campusDTO); err != nil {
		return badRequestErrorWithMessage(err.Error())
	}
	_campus, err := campus.NewBuilder().WithName(campusDTO.Name).WithInstitutionID(campusDTO.InstitutionID).Build()
	if err != nil {
		return responseFromError(err)
	}
	id, err := this.services.Create(_campus)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusCreated, response.IDBuilder().FromUUID(*id))
}

// Update
func (this *campusHandlers) Update(ctx RichContext) error {
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		return badRequestErrorWithMessage(conversionError.Error())
	}
	var campusDTO request.Campus
	if err := ctx.Bind(&campusDTO); err != nil {
		return badRequestErrorWithMessage(err.Error())
	}
	_campus, err := campus.NewBuilder().WithID(id).WithName(campusDTO.Name).WithInstitutionID(campusDTO.InstitutionID).Build()
	if err != nil {
		return responseFromError(err)
	}
	err = this.services.Update(_campus)
	if err != nil {
		return responseFromError(err)
	}
	return successNoContent(ctx)
}

// Delete
func (this *campusHandlers) Delete(ctx RichContext) error {
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		return badRequestErrorWithMessage(conversionError.Error())
	}
	err := this.services.Delete(id)
	if err != nil {
		return responseFromError(err)
	}
	return successNoContent(ctx)
}
