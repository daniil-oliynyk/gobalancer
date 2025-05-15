package main

import (
	"errors"
	"flag"
	"os"
	"time"

	Config "github.com/daniil-oliynyk/gobalancer/config"
	"github.com/valyala/fasthttp"

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
		return
	}

	logger.Info().Msg("Config file opened")

	config, err := Config.ReadConfigFile(*configFilePath)
	if err != nil {
		logger.Error().Err(err).Msgf("Failed to read config file at %s", *configFilePath)
		return
	}

	logger.Info().Msg("Config file read, now validating")

	err = config.ValidateConfig()
	if err != nil {
		logger.Error().Err(err).Msg("Error while validating")
		return
	}

	logger.Info().Msg("Config file validated")
	logger.Debug().Msgf("%#v\n", config)

	// setup backends
	serverPool := NewServerPool(config)
	logger.Info().Msg("Backends created and added to server pool")

	handler := NewLoadBalancerHandler(config, serverPool)

	// start loadbalancer
	loadBalancerServer := fasthttp.Server{
		Logger:       &logger,
		Handler:      handler.Serve(),
		ReadTimeout:  config.Server.ReadTimeout,
		WriteTimeout: config.Server.WriteTimeout,
		IdleTimeout:  config.Server.IdleTimeout,
		Name:         "gobalancer",
	}

	err = loadBalancerServer.ListenAndServe(config.Host + ":" + config.Port)
	if err != nil {
		logger.Error().Err(err).Msg("ListenAndServeError")
		return
	}

}
