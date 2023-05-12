package infra

import (
	"backend_template/src/core"

	"github.com/rs/zerolog"
)

func Logger() zerolog.Logger {
	return core.CoreLogger().With().Str("layer", "infrastructure").Logger()
}
