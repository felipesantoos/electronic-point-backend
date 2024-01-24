package infra

import (
	"backend_template/src/utils"

	"github.com/rs/zerolog"
)

func Logger() zerolog.Logger {
	return utils.Logger().With().Str("layer", "infra").Logger()
}
