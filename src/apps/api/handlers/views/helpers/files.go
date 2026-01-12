package helpers

import (
	"eletronic_point/src/infra"
	"eletronic_point/src/utils"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

var logger = infra.Logger()

// SaveUploadedFile saves an uploaded file to the storage folder and returns the filename
func SaveUploadedFile(c echo.Context, fieldName string) (*string, error) {
	file, header, err := c.Request().FormFile(fieldName)
	if err == http.ErrMissingFile {
		logger.Debug().Msgf("No file provided for field: %s", fieldName)
		return nil, nil
	}
	if err != nil {
		logger.Error().Msgf("Error getting file from form: %s", err.Error())
		return nil, fmt.Errorf("erro ao obter arquivo do formulário: %w", err)
	}
	defer file.Close()

	fileName := fmt.Sprintf("%s%s", uuid.NewString(), utils.ExtractFileExtension(header.Filename))
	storageFolder := os.Getenv("FILE_STORAGE_FOLDER")
	if storageFolder == "" {
		storageFolder = "uploads" // Default fallback
	}

	// Ensure the storage folder exists
	if err := os.MkdirAll(storageFolder, 0755); err != nil {
		logger.Error().Msgf("Error creating storage folder: %s", err.Error())
		return nil, fmt.Errorf("erro ao criar diretório de armazenamento: %w", err)
	}
	
	path := filepath.Join(storageFolder, fileName)
	logger.Info().Msgf("Saving uploaded file to: %s (original name: %s)", path, header.Filename)
	
	out, err := os.Create(path)
	if err != nil {
		logger.Error().Msgf("Error creating file: %s", err.Error())
		return nil, fmt.Errorf("erro ao criar arquivo: %w", err)
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		logger.Error().Msgf("Error copying file content: %s", err.Error())
		return nil, fmt.Errorf("erro ao copiar conteúdo do arquivo: %w", err)
	}

	logger.Info().Msgf("File saved successfully: %s", fileName)
	return &fileName, nil
}
