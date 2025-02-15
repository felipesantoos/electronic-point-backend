package handlers

import (
	"eletronic_point/src/apps/api/handlers/checkers"
	"eletronic_point/src/apps/api/handlers/params"
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/interfaces/primary"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type FileHandlers interface {
	Get(RichContext) error
}

type fileHandlers struct {
	services primary.FilePort
}

func NewFileHandlers(services primary.FilePort) FileHandlers {
	return &fileHandlers{services}
}

// Get
// @ID File.Get
// @Summary Recuperar arquivo salvo no servidor
// @Description Esta rota permite que arquivos salvos no servidor sejam recuperados.
// @Tags Arquivos
// @Security BearerAuth
// @Param name path string true "Nome do arquivo"
// @Success 200 {object} nil "Requisição realizada com sucesso."
// @Failure 400 {object} response.ErrorMessage "Requisição mal formulada."
// @Failure 401 {object} response.ErrorMessage "Usuário não autorizado."
// @Failure 403 {object} response.ErrorMessage "Acesso negado."
// @Failure 404 {object} response.ErrorMessage "Recurso não encontrado."
// @Failure 422 {object} response.ErrorMessage "Ocorreu um erro de validação de dados. Verifique os valores, tipos e formatos de dados enviados."
// @Failure 500 {object} response.ErrorMessage "Ocorreu um erro inesperado. Por favor, contate o suporte."
// @Failure 503 {object} response.ErrorMessage "A base de dados está temporariamente indisponível."
// @Router /files/{name} [get]
func (this fileHandlers) Get(context RichContext) error {
	name := context.Param(params.Name)
	if checkers.IsEmpty(name) {
		err := errors.NewFromString("you must provide a valid filename")
		logger.Error().Msg(err.String())
		return responseFromError(err)
	}
	file, err := this.services.Get(name)
	if err != nil {
		logger.Error().Msg(err.String())
		return responseFromError(err)
	}
	defer file.Close()
	buffer := make([]byte, 512)
	_, readErr := file.Read(buffer)
	if readErr != nil {
		err = errors.NewInternal(readErr)
		logger.Error().Msg(err.String())
		return responseFromError(err)
	}
	mimeType := http.DetectContentType(buffer)
	_, seekErr := file.Seek(0, io.SeekStart)
	if seekErr != nil {
		err := errors.NewInternal(seekErr)
		logger.Error().Msg(err.String())
		return responseFromError(err)
	}
	context.Response().Header().Set(echo.HeaderContentType, mimeType)
	_, copyErr := io.Copy(context.Response(), file)
	if copyErr != nil {
		err := errors.NewInternal(copyErr)
		logger.Error().Msg(err.String())
		return responseFromError(err)
	}
	return nil
}
