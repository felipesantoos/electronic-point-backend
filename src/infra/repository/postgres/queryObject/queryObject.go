package queryObject

import (
	"eletronic_point/src/infra"
)

var logger = infra.Logger().With().Str("port", "postgres").Logger()
