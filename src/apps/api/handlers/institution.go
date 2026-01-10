package handlers

import (
	"eletronic_point/src/apps/api/handlers/checkers"
	"eletronic_point/src/apps/api/handlers/dto/request"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/handlers/params"
	"eletronic_point/src/core/domain/institution"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/services/filters"
	"net/http"

	"github.com/google/uuid"
)

type InstitutionHandlers interface {
	List(RichContext) error
	Get(RichContext) error
	Create(RichContext) error
	Update(RichContext) error
	Delete(RichContext) error
}

type institutionHandlers struct {
	services primary.InstitutionPort
}

func NewInstitutionHandlers(services primary.InstitutionPort) InstitutionHandlers {
	return &institutionHandlers{services}
}

// List
// @ID Institution.List
// @Summary Listar todas as instituições.
// @Description Recupera uma lista de todas as instituições no sistema.
// @Tags Instituições
// @Security BearerAuth
// @Produce json
// @Param name query string false "Nome da instituição"
// @Success 200 {array} response.Institution "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /institutions [get]
func (this *institutionHandlers) List(ctx RichContext) error {
	var name *string
	if !checkers.IsEmpty(ctx.QueryParam(params.Name)) {
		value := ctx.QueryParam(params.Name)
		name = &value
	}
	_filters := filters.InstitutionFilters{Name: name}
	institutions, err := this.services.List(_filters)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.InstitutionBuilder().BuildFromDomainList(institutions))
}

// Get
func (this *institutionHandlers) Get(ctx RichContext) error {
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		return badRequestErrorWithMessage(conversionError.Error())
	}
	_institution, err := this.services.Get(id)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.InstitutionBuilder().BuildFromDomain(_institution))
}

// Create
func (this *institutionHandlers) Create(ctx RichContext) error {
	var institutionDTO request.Institution
	if err := ctx.Bind(&institutionDTO); err != nil {
		return badRequestErrorWithMessage(err.Error())
	}
	_institution, err := institution.NewBuilder().WithName(institutionDTO.Name).Build()
	if err != nil {
		return responseFromError(err)
	}
	id, err := this.services.Create(_institution)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusCreated, response.IDBuilder().FromUUID(*id))
}

// Update
func (this *institutionHandlers) Update(ctx RichContext) error {
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		return badRequestErrorWithMessage(conversionError.Error())
	}
	var institutionDTO request.Institution
	if err := ctx.Bind(&institutionDTO); err != nil {
		return badRequestErrorWithMessage(err.Error())
	}
	_institution, err := institution.NewBuilder().WithID(id).WithName(institutionDTO.Name).Build()
	if err != nil {
		return responseFromError(err)
	}
	err = this.services.Update(_institution)
	if err != nil {
		return responseFromError(err)
	}
	return successNoContent(ctx)
}

// Delete
func (this *institutionHandlers) Delete(ctx RichContext) error {
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
