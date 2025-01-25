package handlers

import (
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/handlers/params"
	"eletronic_point/src/core/interfaces/primary"
	"net/http"

	"github.com/google/uuid"
)

type TimeRecordStatusHandlers interface {
	List(RichContext) error
	Get(RichContext) error
}

type timeRecordStatusHandlers struct {
	services primary.TimeRecordStatusPort
}

func NewTimeRecordStatusHandlers(services primary.TimeRecordStatusPort) TimeRecordStatusHandlers {
	return &timeRecordStatusHandlers{services}
}

// List
// @ID TimeRecordStatus.List
// @Summary Listar status de registro de tempo.
// @Description Recupera uma lista de todos os status ou filtra por nome.
// @Tags Status de Registro de Tempo
// @Security BearerAuth
// @Produce json
// @Success 200 {array} response.TimeRecordStatus "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Router /time-record-status [get]
func (this *timeRecordStatusHandlers) List(ctx RichContext) error {
	timeRecordStatuses, err := this.services.List()
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.TimeRecordStatusBuilder().BuildFromDomainList(timeRecordStatuses))
}

// Get
// @ID TimeRecordStatus.Get
// @Summary Obtenha um status de registro de tempo por ID.
// @Description Recupera os detalhes de um status de registro de tempo específico por ID.
// @Tags Status de Registro de Tempo
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID do status de registro de tempo" default(52613242-6b50-490a-9b4c-90cc3f263e9a)
// @Success 200 {object} response.TimeRecordStatus "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Router /time-record-status/{id} [get]
func (this *timeRecordStatusHandlers) Get(ctx RichContext) error {
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		logger.Error().Msg(conversionError.Error())
		return badRequestErrorWithMessage(conversionError.Error())
	}
	timeRecordStatus, err := this.services.Get(id)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.TimeRecordStatusBuilder().BuildFromDomain(timeRecordStatus))
}
