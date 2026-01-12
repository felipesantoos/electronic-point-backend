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
	"time"

	"github.com/google/uuid"
)

type TimeRecordHandlers interface {
	Create(RichContext) error
	Update(RichContext) error
	Delete(RichContext) error
	List(RichContext) error
	Get(RichContext) error
	Approve(RichContext) error
	Disapprove(RichContext) error
}

type timeRecordHandlers struct {
	services primary.TimeRecordPort
}

func NewTimeRecordHandlers(services primary.TimeRecordPort) TimeRecordHandlers {
	return &timeRecordHandlers{services}
}

// Create
// @ID TimeRecord.Create
// @Summary Crie um novo registro de tempo.
// @Description Cria um novo registro de tempo no sistema com os dados fornecidos.
// @Tags Registros de tempo
// @Security BearerAuth
// @Accept application/json
// @Produce json
// @Param body body request.TimeRecord true "Dados de registro de tempo"
// @Success 201 {object} response.ID "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /time-records [post]
func (this *timeRecordHandlers) Create(ctx RichContext) error {
	if ctx.RoleName() != role.STUDENT_ROLE_CODE && ctx.RoleName() != role.ADMIN_ROLE_CODE {
		return forbiddenErrorWithMessage(messages.YouDoNotHaveAccessToThisResource)
	}
	var timeRecordDTO request.TimeRecord
	if err := ctx.Bind(&timeRecordDTO); err != nil {
		logger.Error().Msg(err.Error())
		return badRequestErrorWithMessage(err.Error())
	}
	_timeRecord, validationError := timeRecordDTO.ToDomain()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return unprocessableEntityErrorWithMessage(validationError.String())
	}
	var studentID uuid.UUID
	if ctx.RoleName() == role.ADMIN_ROLE_CODE {
		if timeRecordDTO.StudentID != nil {
			studentID = *timeRecordDTO.StudentID
		} else {
			return badRequestErrorWithMessage(messages.StudentIDErrorMessage)
		}
	} else if ctx.ProfileID() != nil {
		studentID = *ctx.ProfileID()
	}
	err := _timeRecord.SetStudentID(studentID)
	if err != nil {
		logger.Error().Msg(err.String())
		return unprocessableEntityErrorWithMessage(messages.StudentIDErrorMessage)
	}
	id, err := this.services.Create(_timeRecord)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusCreated, response.IDBuilder().FromUUID(*id))
}

// Update
// @ID TimeRecord.Update
// @Summary Atualizar um registro de tempo existente.
// @Description Atualiza os dados de um registro de tempo existente no sistema.
// @Tags Registros de tempo
// @Security BearerAuth
// @Accept application/json
// @Produce json
// @Param id path string true "ID do registro de tempo" default(ea11bb4b-9aed-4444-9c00-f80bde564063)
// @Param body body request.TimeRecord true "Dados de registro de tempo"
// @Success 204 {object} nil "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /time-records/{id} [put]
func (this *timeRecordHandlers) Update(ctx RichContext) error {
	if ctx.RoleName() != role.STUDENT_ROLE_CODE && ctx.RoleName() != role.ADMIN_ROLE_CODE {
		return forbiddenErrorWithMessage(messages.YouDoNotHaveAccessToThisResource)
	}
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		logger.Error().Msg(conversionError.Error())
		return badRequestErrorWithMessage(conversionError.Error())
	}
	var timeRecordDTO request.TimeRecord
	if err := ctx.Bind(&timeRecordDTO); err != nil {
		logger.Error().Msg(err.Error())
		return badRequestErrorWithMessage(err.Error())
	}
	_timeRecord, validationError := timeRecordDTO.ToDomain()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return unprocessableEntityErrorWithMessage(validationError.String())
	}
	_timeRecord.SetID(id)
	var studentID uuid.UUID
	if ctx.RoleName() == role.ADMIN_ROLE_CODE {
		if timeRecordDTO.StudentID != nil {
			studentID = *timeRecordDTO.StudentID
		} else {
			return badRequestErrorWithMessage(messages.StudentIDErrorMessage)
		}
	} else if ctx.ProfileID() != nil {
		studentID = *ctx.ProfileID()
	}
	err := _timeRecord.SetStudentID(studentID)
	if err != nil {
		logger.Error().Msg(err.String())
		return unprocessableEntityErrorWithMessage(messages.StudentIDErrorMessage)
	}
	err = this.services.Update(_timeRecord)
	if err != nil {
		return responseFromError(err)
	}
	return successNoContent(ctx)
}

// Delete
// @ID TimeRecord.Delete
// @Summary Excluir um registro de tempo por ID.
// @Description Exclui o registro de tempo especificado do sistema.
// @Tags Registros de tempo
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID do registro de tempo" default(ea11bb4b-9aed-4444-9c00-f80bde564063)
// @Success 204 {object} nil "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /time-records/{id} [delete]
func (this *timeRecordHandlers) Delete(ctx RichContext) error {
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
// @ID TimeRecord.List
// @Summary Listar todos os registros de tempo.
// @Description Recupera uma lista de todos os registros de tempo no sistema.
// @Tags Registros de tempo
// @Security BearerAuth
// @Produce json
// @Param studentID query string false "ID do estudante"
// @Param startDate query string false "Data inicial"
// @Param endDate query string false "Data de término"
// @Param statusID query string false "ID do status"
// @Success 200 {array} response.TimeRecord "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /time-records [get]
func (this *timeRecordHandlers) List(ctx RichContext) error {
	var studentID *uuid.UUID
	if !checkers.IsEmpty(ctx.QueryParam(params.StudentID)) {
		value, conversionError := getUUIDQueryParamValue(ctx, params.StudentID)
		if conversionError != nil {
			logger.Error().Msg(conversionError.String())
			return responseFromError(conversionError)
		}
		studentID = value
	}
	var startDate *time.Time
	if !checkers.IsEmpty(ctx.QueryParam(params.StartDate)) {
		value, conversionError := getTimeQueryParamValue(ctx, params.StartDate)
		if conversionError != nil {
			logger.Error().Msg(conversionError.String())
			return responseFromError(conversionError)
		}
		startDate = value
	}
	var endDate *time.Time
	if !checkers.IsEmpty(ctx.QueryParam(params.EndDate)) {
		value, conversionError := getTimeQueryParamValue(ctx, params.EndDate)
		if conversionError != nil {
			logger.Error().Msg(conversionError.String())
			return responseFromError(conversionError)
		}
		endDate = value
	}
	var statusID *uuid.UUID
	if !checkers.IsEmpty(ctx.QueryParam(params.StatusID)) {
		value, conversionError := getUUIDQueryParamValue(ctx, params.StatusID)
		if conversionError != nil {
			logger.Error().Msg(conversionError.String())
			return responseFromError(conversionError)
		}
		statusID = value
	}
	if ctx.RoleName() == role.STUDENT_ROLE_CODE {
		studentID = ctx.ProfileID()
	}
	var teacherID *uuid.UUID
	if ctx.RoleName() == role.TEACHER_ROLE_CODE {
		teacherID = ctx.ProfileID()
	}
	_filters := filters.TimeRecordFilters{
		StudentID: studentID,
		StartDate: startDate,
		EndDate:   endDate,
		TeacherID: teacherID,
		StatusID:  statusID,
	}
	timeRecords, err := this.services.List(_filters)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.TimeRecordBuilder().BuildFromDomainList(timeRecords))
}

// Get
// @ID TimeRecord.Get
// @Summary Obtenha um registro de tempo por ID.
// @Description Recupera os detalhes de um registro de tempo específico por ID.
// @Tags Registros de tempo
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID do registro de tempo" default(ea11bb4b-9aed-4444-9c00-f80bde564063)
// @Success 200 {array} response.TimeRecord "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /time-records/{id} [get]
func (this *timeRecordHandlers) Get(ctx RichContext) error {
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		logger.Error().Msg(conversionError.Error())
		return badRequestErrorWithMessage(conversionError.Error())
	}
	var studentID *uuid.UUID
	if ctx.RoleName() == role.STUDENT_ROLE_CODE {
		studentID = ctx.ProfileID()
	}
	var teacherID *uuid.UUID
	if ctx.RoleName() == role.TEACHER_ROLE_CODE {
		teacherID = ctx.ProfileID()
	}
	_filters := filters.TimeRecordFilters{
		StudentID: studentID,
		TeacherID: teacherID,
	}
	_timeRecord, err := this.services.Get(id, _filters)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.TimeRecordBuilder().BuildFromDomain(_timeRecord))
}

// Approve
// @ID TimeRecord.Approve
// @Summary Aprovar um registro de tempo.
// @Description Atualiza o status de um registro de tempo para "Aprovado".
// @Tags Registros de tempo
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID do registro de tempo" default(ea11bb4b-9aed-4444-9c00-f80bde564063)
// @Success 204 {object} nil "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /time-records/{id}/approve [patch]
func (this *timeRecordHandlers) Approve(ctx RichContext) error {
	if ctx.RoleName() != role.TEACHER_ROLE_CODE && ctx.RoleName() != role.ADMIN_ROLE_CODE {
		return forbiddenErrorWithMessage(messages.YouDoNotHaveAccessToThisResource)
	}
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		logger.Error().Msg(conversionError.Error())
		return badRequestErrorWithMessage(conversionError.Error())
	}
	var approvedBy *uuid.UUID
	if ctx.RoleName() == role.ADMIN_ROLE_CODE {
		var timeRecordDTO request.TimeRecord
		if err := ctx.Bind(&timeRecordDTO); err == nil && timeRecordDTO.TeacherID != nil {
			approvedBy = timeRecordDTO.TeacherID
		} else if val := ctx.QueryParam("teacher_id"); val != "" {
			if uid, err := uuid.Parse(val); err == nil {
				approvedBy = &uid
			}
		}
		if approvedBy == nil {
			return badRequestErrorWithMessage("teacher_id is required for admin")
		}
	} else {
		approvedBy = ctx.ProfileID()
	}
	err := this.services.Approve(id, *approvedBy)
	if err != nil {
		return responseFromError(err)
	}
	return successNoContent(ctx)
}

// Disapprove
// @ID TimeRecord.Disapprove
// @Summary Desaprovar um registro de tempo.
// @Description Atualiza os status de um registro de tempo para "Desaprovado".
// @Tags Registros de tempo
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID do registro de tempo" default(ea11bb4b-9aed-4444-9c00-f80bde564063)
// @Success 204 {object} nil "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /time-records/{id}/disapprove [patch]
func (this *timeRecordHandlers) Disapprove(ctx RichContext) error {
	if ctx.RoleName() != role.TEACHER_ROLE_CODE && ctx.RoleName() != role.ADMIN_ROLE_CODE {
		return forbiddenErrorWithMessage(messages.YouDoNotHaveAccessToThisResource)
	}
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		logger.Error().Msg(conversionError.Error())
		return badRequestErrorWithMessage(conversionError.Error())
	}
	var disapprovedBy *uuid.UUID
	if ctx.RoleName() == role.ADMIN_ROLE_CODE {
		var timeRecordDTO request.TimeRecord
		if err := ctx.Bind(&timeRecordDTO); err == nil && timeRecordDTO.TeacherID != nil {
			disapprovedBy = timeRecordDTO.TeacherID
		} else if val := ctx.QueryParam("teacher_id"); val != "" {
			if uid, err := uuid.Parse(val); err == nil {
				disapprovedBy = &uid
			}
		}
		if disapprovedBy == nil {
			return badRequestErrorWithMessage("teacher_id is required for admin")
		}
	} else {
		disapprovedBy = ctx.ProfileID()
	}
	err := this.services.Disapprove(id, *disapprovedBy)
	if err != nil {
		return responseFromError(err)
	}
	return successNoContent(ctx)
}
