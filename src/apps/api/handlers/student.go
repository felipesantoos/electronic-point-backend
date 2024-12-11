package handlers

import (
	"eletronic_point/src/apps/api/handlers/dto/request"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/handlers/formData"
	"eletronic_point/src/apps/api/handlers/params"
	"eletronic_point/src/core/interfaces/primary"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/google/uuid"
)

type StudentHandlers interface {
	Create(RichContext) error
	Update(RichContext) error
	Delete(RichContext) error
	List(RichContext) error
	Get(RichContext) error
}

type studentHandlers struct {
	services primary.StudentPort
}

func NewStudentHandlers(services primary.StudentPort) StudentHandlers {
	return &studentHandlers{services}
}

// Create
// @ID Student.Create
// @Summary Criar um novo estudante.
// @Description Cria um novo estudante no sistema com os dados fornecidos. O campo `profile_picture` deve ser enviado como um arquivo em um formulário.
// @Tags Estudantes
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Nome do estudante" default(Nome 1)
// @Param registration formData string true "Matrícula do estudante" default(0000000001)
// @Param profile_picture formData file false "Foto de perfil do estudante (arquivo de imagem)"
// @Param institution formData string true "Instituição do estudante" default(IFAL 1)
// @Param course formData string true "Curso do estudante" default(Curso 1)
// @Param internship_location_name formData string true "Nome do local de estágio" default(Nome do local 1)
// @Param internship_address formData string true "Endereço do local de estágio" default(Endereço 1)
// @Param internship_location formData string true "Localização do estágio" default(Localização 1)
// @Param total_workload formData int true "Carga horária total do estágio" default(100)
// @Success 201 {object} response.ID "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /students [post]
func (this *studentHandlers) Create(ctx RichContext) error {
	if formDataError := ctx.Request().ParseMultipartForm(10 << 20); formDataError != nil {
		logger.Error().Msg(formDataError.Error())
		return badRequestErrorWithMessage(formDataError.Error())
	}
	name := ctx.FormValue(formData.StudentName)
	registration := ctx.FormValue(formData.StudentRegistration)
	institution := ctx.FormValue(formData.StudentInstitution)
	course := ctx.FormValue(formData.StudentCourse)
	internshipLocationName := ctx.FormValue(formData.StudentInternshipLocationName)
	internshipAddress := ctx.FormValue(formData.StudentInternshipAddress)
	internshipLocation := ctx.FormValue(formData.StudentInternshipLocation)
	totalWorkload, conversionError := strconv.Atoi(ctx.FormValue(formData.StudentTotalWorkload))
	if conversionError != nil {
		logger.Error().Msg(conversionError.Error())
		return badRequestErrorWithMessage(conversionError.Error())
	}
	var filePath *string
	file, header, formFileError := ctx.Request().FormFile(formData.StudentProfilePicture)
	if formFileError == nil {
		defer file.Close()
		path := "uploads/" + header.Filename
		out, err := os.Create(path)
		if err != nil {
			logger.Error().Msg(err.Error())
			return unprocessableEntityErrorWithMessage(err.Error())
		}
		defer out.Close()
		if _, err := io.Copy(out, file); err != nil {
			logger.Error().Msg(err.Error())
			return unprocessableEntityErrorWithMessage(err.Error())
		}
		filePath = &path
	} else if formFileError != http.ErrMissingFile {
		logger.Error().Msg(formFileError.Error())
		return badRequestErrorWithMessage(formFileError.Error())
	}
	studentDTO := request.Student{
		Name:                   name,
		Registration:           registration,
		ProfilePicture:         filePath,
		Institution:            institution,
		Course:                 course,
		InternshipLocationName: internshipLocationName,
		InternshipAddress:      internshipAddress,
		InternshipLocation:     internshipLocation,
		TotalWorkload:          totalWorkload,
	}
	_student, validationError := studentDTO.ToDomain()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return unprocessableEntityErrorWithMessage(validationError.String())
	}
	id, err := this.services.Create(_student)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusCreated, response.IDBuilder().FromUUID(*id))
}

// Update
// @ID Student.Update
// @Summary Atualizar informações de um estudante.
// @Description Atualiza os dados de um estudante existente no sistema.
// @Tags Estudantes
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "ID do estudante"
// @Param name formData string true "Nome do estudante" default(Nome 1)
// @Param registration formData string true "Matrícula do estudante" default(0000000001)
// @Param profile_picture formData file false "Foto de perfil do estudante (arquivo de imagem)"
// @Param institution formData string true "Instituição do estudante" default(IFAL 1)
// @Param course formData string true "Curso do estudante" default(Curso 1)
// @Param internship_location_name formData string true "Nome do local de estágio" default(Nome do local 1)
// @Param internship_address formData string true "Endereço do local de estágio" default(Endereço 1)
// @Param internship_location formData string true "Localização do estágio" default(Localização 1)
// @Param total_workload formData int true "Carga horária total do estágio" default(100)
// @Success 204 {object} nil "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /students/{id} [put]
func (this *studentHandlers) Update(ctx RichContext) error {
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		return badRequestErrorWithMessage(conversionError.Error())
	}
	if formDataError := ctx.Request().ParseMultipartForm(10 << 20); formDataError != nil {
		logger.Error().Msg(formDataError.Error())
		return badRequestErrorWithMessage(formDataError.Error())
	}
	name := ctx.FormValue(formData.StudentName)
	registration := ctx.FormValue(formData.StudentRegistration)
	institution := ctx.FormValue(formData.StudentInstitution)
	course := ctx.FormValue(formData.StudentCourse)
	internshipLocationName := ctx.FormValue(formData.StudentInternshipLocationName)
	internshipAddress := ctx.FormValue(formData.StudentInternshipAddress)
	internshipLocation := ctx.FormValue(formData.StudentInternshipLocation)
	totalWorkload, conversionError := strconv.Atoi(ctx.FormValue(formData.StudentTotalWorkload))
	if conversionError != nil {
		logger.Error().Msg(conversionError.Error())
		return badRequestErrorWithMessage(conversionError.Error())
	}
	var filePath *string
	file, header, formFileError := ctx.Request().FormFile(formData.StudentProfilePicture)
	if formFileError == nil {
		defer file.Close()
		path := "uploads/" + header.Filename
		out, err := os.Create(path)
		if err != nil {
			logger.Error().Msg(err.Error())
			return unprocessableEntityErrorWithMessage(err.Error())
		}
		defer out.Close()
		if _, err := io.Copy(out, file); err != nil {
			logger.Error().Msg(err.Error())
			return unprocessableEntityErrorWithMessage(err.Error())
		}
		filePath = &path
	} else if formFileError != http.ErrMissingFile {
		logger.Error().Msg(formFileError.Error())
		return badRequestErrorWithMessage(formFileError.Error())
	}
	studentDTO := request.Student{
		Name:                   name,
		Registration:           registration,
		ProfilePicture:         filePath,
		Institution:            institution,
		Course:                 course,
		InternshipLocationName: internshipLocationName,
		InternshipAddress:      internshipAddress,
		InternshipLocation:     internshipLocation,
		TotalWorkload:          totalWorkload,
	}
	_student, validationError := studentDTO.ToDomain()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return unprocessableEntityErrorWithMessage(validationError.String())
	}
	validationError = _student.SetID(id)
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return unprocessableEntityErrorWithMessage(validationError.String())
	}
	err := this.services.Update(_student)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.NoContent(http.StatusNoContent)
}

// Delete Student
// @ID Student.Delete
// @Summary Deletar um estudante.
// @Description Remove um estudante do sistema.
// @Security	bearerAuth
// @Tags Estudantes
// @Produce json
// @Param id path string true "ID do estudante"
// @Success 204 {object} nil "Estudante removido com sucesso."
// @Failure 404 {object} response.ErrorMessage "Estudante não encontrado."
// @Failure 500 {object} response.ErrorMessage "Erro inesperado. Por favor, entre em contato com o suporte."
// @Router /students/{id} [delete]
func (h *studentHandlers) Delete(context RichContext) error {
	id := context.Param("id")
	studentID, conversionError := uuid.Parse(id)
	if conversionError != nil {
		return response.ErrorBuilder().NewBadRequestFromCoreError()
	}
	err := h.services.Delete(studentID)
	if err != nil {
		return response.ErrorBuilder().NewFromDomain(err)
	}

	return context.NoContent(http.StatusNoContent)
}

// List Students
// @ID Student.List
// @Summary Listar todos os estudantes.
// @Description Recupera todos os estudantes registrados no sistema.
// @Security	bearerAuth
// @Tags Estudantes
// @Produce json
// @Success 200 {array} response.Student "Lista de estudantes."
// @Failure 500 {object} response.ErrorMessage "Erro inesperado. Por favor, entre em contato com o suporte."
// @Router /students [get]
func (h *studentHandlers) List(context RichContext) error {
	result, err := h.services.List()
	if err != nil {
		return response.ErrorBuilder().NewFromDomain(err)
	}
	return context.JSON(http.StatusOK, response.StudentBuilder().BuildFromDomainList(result))
}

// Get Student by ID
// @ID Student.Get
// @Summary Obter um estudante por ID.
// @Description Recupera os dados de um estudante específico pelo seu ID.
// @Security	bearerAuth
// @Tags Estudantes
// @Produce json
// @Param id path string true "ID do estudante"
// @Success 200 {object} response.Student "Estudante encontrado."
// @Failure 404 {object} response.ErrorMessage "Estudante não encontrado."
// @Failure 500 {object} response.ErrorMessage "Erro inesperado. Por favor, entre em contato com o suporte."
// @Router /students/{id} [get]
func (h *studentHandlers) Get(context RichContext) error {
	id := context.Param("id")
	studentID, conversionError := uuid.Parse(id)
	if conversionError != nil {
		return response.ErrorBuilder().NewBadRequestFromCoreError()
	}
	result, err := h.services.Get(studentID)
	if err != nil {
		return response.ErrorBuilder().NewFromDomain(err)
	}

	return context.JSON(http.StatusOK, response.StudentBuilder().BuildFromDomain(result))
}
