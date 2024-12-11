package infra

import (
	"eletronic_point/src/utils"

	"github.com/rs/zerolog"
)

func Logger() zerolog.Logger {
	return utils.Logger().With().Str("layer", "infra").Logger()
}
