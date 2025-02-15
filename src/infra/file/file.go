package file

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/interfaces/secondary"
	"eletronic_point/src/core/messages"
	"fmt"
	"os"
	"strings"
)

type fileRepository struct{}

func NewFileRepository() secondary.FilePort {
	return &fileRepository{}
}

func (this fileRepository) Get(name string) (*os.File, errors.Error) {
	file, err := os.Open(fmt.Sprintf("%s/%s", os.Getenv("FILE_STORAGE_FOLDER"), name))
	if err != nil {
		logger.Error().Msg(err.Error())
		if strings.Contains(err.Error(), "no such file or directory") {
			return nil, errors.NewFromString(messages.FileNotFound)
		}
		return nil, errors.NewInternal(err)
	}
	return file, nil
}
