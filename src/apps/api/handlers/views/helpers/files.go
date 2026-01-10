package helpers

import (
	"eletronic_point/src/utils"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

// SaveUploadedFile saves an uploaded file to the storage folder and returns the filename
func SaveUploadedFile(c echo.Context, fieldName string) (*string, error) {
	file, header, err := c.Request().FormFile(fieldName)
	if err == http.ErrMissingFile {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	fileName := fmt.Sprintf("%s%s", uuid.NewString(), utils.ExtractFileExtension(header.Filename))
	storageFolder := os.Getenv("FILE_STORAGE_FOLDER")
	if storageFolder == "" {
		storageFolder = "uploads" // Default fallback
	}
	
	path := fmt.Sprintf("%s/%s", storageFolder, fileName)
	out, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	defer out.Close()

	if _, err := io.Copy(out, file); err != nil {
		return nil, err
	}

	return &fileName, nil
}
