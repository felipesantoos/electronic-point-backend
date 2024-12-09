package queryObject

import (
	"backend_template/src/infra"
)

var logger = infra.Logger().With().Str("port", "postgres").Logger()
