package app

import (
	"tripstory/config"

	"github.com/rs/zerolog/log"
)

func RunServer() {
	cfg := config.NewConfig()
	_, err := cfg.ConnectionPostgress()

	if err != nil {
		log.Fatal().Msgf("Error connecting to database: %v", err)
		return
	}
}
