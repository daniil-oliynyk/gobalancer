package main

import (
	"errors"
	"flag"
	"os"
	"time"

	Config "github.com/daniil-oliynyk/gobalancer/config"

	"github.com/rs/zerolog"
)

func main() {

	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Caller().
		Logger()

	configFilePath := flag.String("config", "./config.yaml", "use absolute path for config file")
	flag.Parse()

	if *configFilePath == "" {
		logger.Error().Msg("Provide a file path for the config file")
		return
	}

	if _, err := os.Stat(*configFilePath); errors.Is(err, os.ErrNotExist) {
		logger.Error().Msgf("Config file does not exist at path %s", *configFilePath)
	}

	logger.Info().Msg("Config file opened")

	config, err := Config.ReadConfigFile(*configFilePath)
	if err != nil {
		logger.Err(err).Msgf("Failed to read config file at %s", *configFilePath)
	}

	logger.Info().Msg("Config file read, now validating")

	err = config.ValidateConfig()
	if err != nil {
		logger.Err(err)
	}

}
