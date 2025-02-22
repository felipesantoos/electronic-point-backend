package handlers

import (
	"eletronic_point/src/apps/api/handlers/checkers"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/handlers/params"
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/services/filters"
	"net/http"
)

type ReportsHandlers interface {
	GetTimeRecordsByStudent(RichContext) error
}

type reportsHandlers struct {
	services primary.ReportsPort
}

func NewReportsHandlers(services primary.ReportsPort) ReportsHandlers {
	return &reportsHandlers{services}
}

// GetTimeRecordsByStudent
// @ID Reports.GetTimeRecordsByStudent
// @Summary Listar todos os registros de tempo agrupados por estudante.
// @Description Recupera uma lista de todos os registros de tempo agrupados por estudante no sistema.
// @Tags Relatórios
// @Security BearerAuth
// @Produce json
// @Param institutionID query string false "ID da instituição"
// @Param campusID query string false "ID do campus"
// @Param studentID query string false "ID do estudante"
// @Param startDate query string false "Data inicial"
// @Param endDate query string false "Data de término"
// @Param statusID query string false "ID do status"
// @Success 200 {array} response.TimeRecordsByStudent "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /reports/time-records-by-student [get]
func (this *reportsHandlers) GetTimeRecordsByStudent(ctx RichContext) error {
	studentFilters := filters.StudentFilters{}
	timeRecordFilters := filters.TimeRecordFilters{}
	if ctx.RoleName() == role.TEACHER_ROLE_CODE {
		studentFilters.TeacherID = ctx.ProfileID()
		timeRecordFilters.TeacherID = ctx.ProfileID()
	} else {
		return forbiddenError
	}
	if !checkers.IsEmpty(ctx.QueryParam(params.InstitutionID)) {
		value, conversionError := getUUIDQueryParamValue(ctx, params.InstitutionID)
		if conversionError != nil {
			logger.Error().Msg(conversionError.String())
			return responseFromError(conversionError)
		}
		studentFilters.InstitutionID = value
	}
	if !checkers.IsEmpty(ctx.QueryParam(params.CampusID)) {
		value, conversionError := getUUIDQueryParamValue(ctx, params.CampusID)
		if conversionError != nil {
			logger.Error().Msg(conversionError.String())
			return responseFromError(conversionError)
		}
		studentFilters.CampusID = value
	}
	if !checkers.IsEmpty(ctx.QueryParam(params.StudentID)) {
		value, conversionError := getUUIDQueryParamValue(ctx, params.StudentID)
		if conversionError != nil {
			logger.Error().Msg(conversionError.String())
			return responseFromError(conversionError)
		}
		studentFilters.StudentID = value
		timeRecordFilters.StudentID = value
	}
	if !checkers.IsEmpty(ctx.QueryParam(params.StartDate)) {
		value, conversionError := getTimeQueryParamValue(ctx, params.StartDate)
		if conversionError != nil {
			logger.Error().Msg(conversionError.String())
			return responseFromError(conversionError)
		}
		timeRecordFilters.StartDate = value
	}
	if !checkers.IsEmpty(ctx.QueryParam(params.EndDate)) {
		value, conversionError := getTimeQueryParamValue(ctx, params.EndDate)
		if conversionError != nil {
			logger.Error().Msg(conversionError.String())
			return responseFromError(conversionError)
		}
		timeRecordFilters.EndDate = value
	}
	if !checkers.IsEmpty(ctx.QueryParam(params.StatusID)) {
		value, conversionError := getUUIDQueryParamValue(ctx, params.StatusID)
		if conversionError != nil {
			logger.Error().Msg(conversionError.String())
			return responseFromError(conversionError)
		}
		timeRecordFilters.StatusID = value
	}

	timeRecordsByStudentFilters := filters.TimeRecordsByStudentFilters{
		StudentFilters:    studentFilters,
		TimeRecordFilters: timeRecordFilters,
	}
	reports, err := this.services.GetTimeRecordsByStudent(timeRecordsByStudentFilters)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.TimeRecordsByStudentBuilder().BuildFromDomainList(reports))
}
