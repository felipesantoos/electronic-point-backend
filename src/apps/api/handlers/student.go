package handlers

import (
	"eletronic_point/src/apps/api/handlers/dto/request"
	"eletronic_point/src/apps/api/handlers/dto/response"
	"eletronic_point/src/core/interfaces/primary"
	"net/http"

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
// @Description Cria um novo estudante no sistema com os dados fornecidos.
// @Tags Estudantes
// @Accept json
// @Produce json
// @Param student body request.Student true "Dados do estudante"
// @Success 201 {object} response.ID "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 409 {object} response.ErrorMessage "A solicitação não pôde ser concluída devido a um conflito com o estado atual do recurso de destino."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Vefique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /students [post]
func (this *studentHandlers) Create(ctx RichContext) error {
	var studentDTO request.Student
	if err := ctx.Bind(&studentDTO); err != nil {
		logger.Error().Msg(err.Error())
		return badRequestErrorWithMessage(err.Error())
	}
	student, err := studentDTO.ToDomain()
	if err != nil {
		logger.Error().Msg(err.String())
		return unprocessableEntityErrorWithMessage(err.String())
	}
	id, err := this.services.Create(student)
	if err != nil {
		return responseFromError(err)
	}
	return ctx.JSON(http.StatusCreated, response.IDBuilder().FromUUID(*id))
}

// Update Student
// @ID Student.Update
// @Summary Atualizar informações de um estudante.
// @Description Atualiza os dados de um estudante existente no sistema.
// @Security	bearerAuth
// @Tags Estudantes
// @Accept json
// @Produce json
// @Param id path string true "ID do estudante"
// @Param student body request.Student true "Dados do estudante"
// @Success 200 {object} response.Student "Estudante atualizado com sucesso."
// @Failure 400 {object} response.ErrorMessage "Dados inválidos fornecidos."
// @Failure 404 {object} response.ErrorMessage "Estudante não encontrado."
// @Failure 500 {object} response.ErrorMessage "Erro inesperado. Por favor, entre em contato com o suporte."
// @Router /students/{id} [put]
func (h *studentHandlers) Update(context RichContext) error {
	id := context.Param("id")
	studentID, conversionError := uuid.Parse(id)
	if conversionError != nil {
		return response.ErrorBuilder().NewBadRequestFromCoreError()
	}
	var studentRequest request.Student
	if err := context.Bind(&studentRequest); err != nil {
		return response.ErrorBuilder().NewBadRequestFromCoreError()
	}
	student, validationError := studentRequest.ToDomain()
	if validationError != nil {
		return response.ErrorBuilder().NewFromDomain(validationError)
	}
	err := student.SetID(studentID)
	if err != nil {
		return response.ErrorBuilder().NewFromDomain(err)
	}
	err = h.services.Update(student)
	if err != nil {
		return response.ErrorBuilder().NewFromDomain(err)
	}
	return context.NoContent(http.StatusNoContent)
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
