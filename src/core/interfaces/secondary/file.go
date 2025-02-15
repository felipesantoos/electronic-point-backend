package secondary

import (
	"eletronic_point/src/core/domain/errors"
	"os"
)

type FilePort interface {
	Get(name string) (*os.File, errors.Error)
}
