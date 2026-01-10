package handlers

import (
	"eletronic_point/src/apps/api/handlers/checkers"
	"eletronic_point/src/apps/api/handlers/dto/request"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/apps/api/handlers/formData"
	"eletronic_point/src/apps/api/handlers/params"
	"eletronic_point/src/core/domain/role"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/services/filters"
	"eletronic_point/src/utils"
	"fmt"
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
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param name formData string true "Nome do estudante" default(Nome 1)
// @Param birth_date formData string true "Data de nascimento" default(2000-01-01)
// @Param cpf formData string true "CPF do estudante" default(73595867041)
// @Param email formData string true "Email do estudante" default(email@example.com)
// @Param phone formData string true "Telefone do estudante" default(82999999999)
// @Param registration formData string true "Matrícula do estudante" default(0000000001)
// @Param profile_picture formData file false "Foto de perfil do estudante (arquivo de imagem)"
// @Param campus_id formData string true "ID do campus do estudante" default(6de43e83-6bdf-4637-83d0-bcb8611082be)
// @Param course_id formData string true "ID do curso do estudante" default(9e29482b-408e-49df-84fa-1543af68e036)
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
	personID := ctx.ProfileID()
	if formDataError := ctx.Request().ParseMultipartForm(10 << 20); formDataError != nil {
		logger.Error().Msg(formDataError.Error())
		return badRequestErrorWithMessage(formDataError.Error())
	}
	name := ctx.FormValue(formData.StudentName)
	birthDate := ctx.FormValue(formData.StudentBirthDate)
	cpf := ctx.FormValue(formData.StudentCPF)
	email := ctx.FormValue(formData.StudentEmail)
	phone := ctx.FormValue(formData.StudentPhone)
	registration := ctx.FormValue(formData.StudentRegistration)
	campusID, err := getUUIDFormDataValue(ctx, formData.StudentCampusID)
	if err != nil {
		logger.Error().Msg(err.String())
		return unprocessableEntityErrorWithMessage(err.String())
	}
	courseID, err := getUUIDFormDataValue(ctx, formData.StudentCourseID)
	if err != nil {
		logger.Error().Msg(err.String())
		return unprocessableEntityErrorWithMessage(err.String())
	}
	totalWorkload, conversionError := strconv.Atoi(ctx.FormValue(formData.StudentTotalWorkload))
	if conversionError != nil {
		logger.Error().Msg(conversionError.Error())
		return badRequestErrorWithMessage(conversionError.Error())
	}
	var fileName *string
	file, header, formFileError := ctx.Request().FormFile(formData.StudentProfilePicture)
	if formFileError == nil {
		defer file.Close()
		name := fmt.Sprintf("%s%s", uuid.NewString(), utils.ExtractFileExtension(header.Filename))
		path := fmt.Sprintf("%s/%s", os.Getenv("FILE_STORAGE_FOLDER"), name)
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
		fileName = &name
	} else if formFileError != http.ErrMissingFile {
		logger.Error().Msg(formFileError.Error())
		return badRequestErrorWithMessage(formFileError.Error())
	}
	studentDTO := request.Student{
		Name:           name,
		BirthDate:      birthDate,
		CPF:            cpf,
		Email:          email,
		Phone:          phone,
		Registration:   registration,
		ProfilePicture: fileName,
		CampusID:       *campusID,
		CourseID:       *courseID,
		TotalWorkload:  totalWorkload,
	}
	_student, validationError := studentDTO.ToDomain()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return unprocessableEntityErrorWithMessage(validationError.String())
	}
	err = _student.SetResponsibleTeacherID(*personID)
	if err != nil {
		logger.Error().Msg(err.String())
		return unprocessableEntityErrorWithMessage(err.String())
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
// @Security BearerAuth
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "ID do estudante" default(02e62826-bf41-4944-adb2-051b6a30a131)
// @Param name formData string true "Nome do estudante" default(Nome 1)
// @Param birth_date formData string true "Data de nascimento" default(2000-01-01)
// @Param cpf formData string true "CPF do estudante" default(73595867041)
// @Param email formData string true "Email do estudante" default(email@example.com)
// @Param phone formData string true "Telefone do estudante" default(82999999999)
// @Param registration formData string true "Matrícula do estudante" default(0000000001)
// @Param profile_picture formData file false "Foto de perfil do estudante (arquivo de imagem)"
// @Param campus_id formData string true "ID do campus do estudante" default(6de43e83-6bdf-4637-83d0-bcb8611082be)
// @Param course_id formData string true "ID do curso do estudante" default(9e29482b-408e-49df-84fa-1543af68e036)
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
		logger.Error().Msg(conversionError.Error())
		return badRequestErrorWithMessage(conversionError.Error())
	}
	if formDataError := ctx.Request().ParseMultipartForm(10 << 20); formDataError != nil {
		logger.Error().Msg(formDataError.Error())
		return badRequestErrorWithMessage(formDataError.Error())
	}
	name := ctx.FormValue(formData.StudentName)
	birthDate := ctx.FormValue(formData.StudentBirthDate)
	cpf := ctx.FormValue(formData.StudentCPF)
	email := ctx.FormValue(formData.StudentEmail)
	phone := ctx.FormValue(formData.StudentPhone)
	registration := ctx.FormValue(formData.StudentRegistration)
	campusID, err := getUUIDFormDataValue(ctx, formData.StudentCampusID)
	if err != nil {
		logger.Error().Msg(err.String())
		return unprocessableEntityErrorWithMessage(err.String())
	}
	courseID, err := getUUIDFormDataValue(ctx, formData.StudentCourseID)
	if err != nil {
		logger.Error().Msg(err.String())
		return unprocessableEntityErrorWithMessage(err.String())
	}
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
		Name:           name,
		BirthDate:      birthDate,
		CPF:            cpf,
		Email:          email,
		Phone:          phone,
		Registration:   registration,
		ProfilePicture: filePath,
		CampusID:       *campusID,
		CourseID:       *courseID,
		TotalWorkload:  totalWorkload,
	}
	_student, validationError := studentDTO.ToDomain()
	if validationError != nil {
		logger.Error().Msg(validationError.String())
		return unprocessableEntityErrorWithMessage(validationError.String())
	}
	_student.SetID(&id)
	err = this.services.Update(_student)
	if err != nil {
		return responseFromError(err)
	}
	return successNoContent(ctx)
}

// Delete
// @ID Student.Delete
// @Summary Deletar um estudante.
// @Description Remove um estudante do sistema.
// @Tags Estudantes
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID do estudante" default(02e62826-bf41-4944-adb2-051b6a30a131)
// @Success 204 {object} nil "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /students/{id} [delete]
func (this *studentHandlers) Delete(ctx RichContext) error {
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
// @ID Student.List
// @Summary Listar todos os estudantes.
// @Description Recupera todos os estudantes registrados no sistema.
// @Tags Estudantes
// @Security BearerAuth
// @Produce json
// @Param institutionID query string false "ID da instituição"
// @Param campusID query string false "ID do campus"
// @Success 200 {array} response.StudentList "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /students [get]
func (this *studentHandlers) List(ctx RichContext) error {
	_filters := filters.StudentFilters{}
	if ctx.RoleName() == role.TEACHER_ROLE_CODE {
		_filters.TeacherID = ctx.ProfileID()
	} else if ctx.RoleName() != role.ADMIN_ROLE_CODE {
		return forbiddenError
	}
	if !checkers.IsEmpty(ctx.QueryParam(params.InstitutionID)) {
		value, conversionError := getUUIDQueryParamValue(ctx, params.InstitutionID)
		if conversionError != nil {
			logger.Error().Msg(conversionError.String())
			return responseFromError(conversionError)
		}
		_filters.InstitutionID = value
	}
	if !checkers.IsEmpty(ctx.QueryParam(params.CampusID)) {
		value, conversionError := getUUIDQueryParamValue(ctx, params.CampusID)
		if conversionError != nil {
			logger.Error().Msg(conversionError.String())
			return responseFromError(conversionError)
		}
		_filters.CampusID = value
	}
	result, err := this.services.List(_filters)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.StudentListBuilder().BuildFromDomainList(result))
}

// Get
// @ID Student.Get
// @Summary Obter um estudante por ID.
// @Description Recupera os dados de um estudante específico pelo seu ID.
// @Tags Estudantes
// @Security BearerAuth
// @Produce json
// @Param id path string true "ID do estudante" default(02e62826-bf41-4944-adb2-051b6a30a131)
// @Success 200 {array} response.Student "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /students/{id} [get]
func (this *studentHandlers) Get(ctx RichContext) error {
	id, conversionError := uuid.Parse(ctx.Param(params.ID))
	if conversionError != nil {
		logger.Error().Msg(conversionError.Error())
		return badRequestErrorWithMessage(conversionError.Error())
	}
	_filters := filters.StudentFilters{}
	if ctx.RoleName() == role.TEACHER_ROLE_CODE {
		_filters.TeacherID = ctx.ProfileID()
	} else if ctx.RoleName() != role.ADMIN_ROLE_CODE {
		return forbiddenError
	}
	result, err := this.services.Get(id, _filters)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusOK, response.StudentBuilder().BuildFromDomain(result))
}
