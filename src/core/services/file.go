package services

import (
	"eletronic_point/src/core/domain/errors"
	"eletronic_point/src/core/interfaces/primary"
	"eletronic_point/src/core/interfaces/secondary"
	"os"
)

type fileServices struct {
	repository secondary.FilePort
}

func NewFileService(repository secondary.FilePort) primary.FilePort {
	return &fileServices{repository}
}

func (this fileServices) Get(name string) (*os.File, errors.Error) {
	return this.repository.Get(name)
}
