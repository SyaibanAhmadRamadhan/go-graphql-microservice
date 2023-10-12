package infra

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func LogInit() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	log.Logger = log.With().Caller().Logger()

	log.Info().Msg("Logger initialization successfully")
}
