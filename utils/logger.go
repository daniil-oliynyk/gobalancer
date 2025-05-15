package utils

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.DateTime}).
	Level(zerolog.TraceLevel).
	With().
	Timestamp().
	Caller().
	Logger()

func GetLogger() *zerolog.Logger {
	return &logger
}
