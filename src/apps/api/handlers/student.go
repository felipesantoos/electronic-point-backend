package handlers

import (
	"backend_template/src/apps/api/handlers/dto/request"
	"backend_template/src/apps/api/handlers/dto/response"
	"backend_template/src/core/interfaces/usecases"
	"net/http"

	"github.com/google/uuid"
)

type StudentHandler interface {
	Create(RichContext) error
	Update(RichContext) error
	Delete(RichContext) error
	List(RichContext) error
	Get(RichContext) error
}

type studentHandler struct {
	usecase usecases.StudentUseCase
}

func NewStudentHandler(usecase usecases.StudentUseCase) StudentHandler {
	return &studentHandler{usecase}
}

// Create Student
// @ID Student.Create
// @Summary Criar um novo estudante.
// @Description Cria um novo estudante no sistema com os dados fornecidos.
// @Security	bearerAuth
// @Tags Estudantes
// @Accept json
// @Produce json
// @Param student body request.CreateStudent true "Dados do estudante"
// @Success 201 {object} response.Student "Estudante criado com sucesso."
// @Failure 400 {object} response.ErrorMessage "Dados inválidos fornecidos."
// @Failure 500 {object} response.ErrorMessage "Erro inesperado. Por favor, entre em contato com o suporte."
// @Router /students [post]
func (h *studentHandler) Create(context RichContext) error {
	var studentRequest request.CreateStudent
	if err := context.Bind(&studentRequest); err != nil {
		return response.ErrorBuilder().NewBadRequestFromCoreError()
	}
	student, err := studentRequest.ToDomain()
	if err != nil {
		return response.ErrorBuilder().NewFromDomain(err)
	}
	id, err := h.usecase.Create(student)
	if err != nil {
		return response.ErrorBuilder().NewFromDomain(err)
	}
	return context.JSON(http.StatusCreated, response.IDBuilder().FromUUID(*id))
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
// @Param student body request.CreateStudent true "Dados do estudante"
// @Success 200 {object} response.Student "Estudante atualizado com sucesso."
// @Failure 400 {object} response.ErrorMessage "Dados inválidos fornecidos."
// @Failure 404 {object} response.ErrorMessage "Estudante não encontrado."
// @Failure 500 {object} response.ErrorMessage "Erro inesperado. Por favor, entre em contato com o suporte."
// @Router /students/{id} [put]
func (h *studentHandler) Update(context RichContext) error {
	id := context.Param("id")
	studentID, conversionError := uuid.Parse(id)
	if conversionError != nil {
		return response.ErrorBuilder().NewBadRequestFromCoreError()
	}
	var studentRequest request.CreateStudent
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
	err = h.usecase.Update(student)
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
func (h *studentHandler) Delete(context RichContext) error {
	id := context.Param("id")
	studentID, conversionError := uuid.Parse(id)
	if conversionError != nil {
		return response.ErrorBuilder().NewBadRequestFromCoreError()
	}
	err := h.usecase.Delete(studentID)
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
func (h *studentHandler) List(context RichContext) error {
	result, err := h.usecase.List()
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
func (h *studentHandler) Get(context RichContext) error {
	id := context.Param("id")
	studentID, conversionError := uuid.Parse(id)
	if conversionError != nil {
		return response.ErrorBuilder().NewBadRequestFromCoreError()
	}
	result, err := h.usecase.Get(studentID)
	if err != nil {
		return response.ErrorBuilder().NewFromDomain(err)
	}

	return context.JSON(http.StatusOK, response.StudentBuilder().BuildFromDomain(result))
}
