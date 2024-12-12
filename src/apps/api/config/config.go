package config

import (
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

func Load() *Config {
	log.Info().Msg("Setting up application...")

	if err := godotenv.Load(".env"); err != nil {
		log.Info().Msgf(".env could not be read: %s", err.Error())
	}

	var r reader
	config, err := r.config()
	if err != nil {
		log.Fatal().Msgf("Failed to read environment file: %s", err.Error())
	}

	return config
}
