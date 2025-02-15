package file

import (
	"eletronic_point/src/infra"
)

var logger = infra.Logger().With().Str("port", "file").Logger()
