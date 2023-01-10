package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const ConsoleLoggerType = "console"

func GetLogger(logType, logLevel string) zerolog.Logger {
	if logType == ConsoleLoggerType {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	}

	level, err := zerolog.ParseLevel(logLevel)
	if err != nil {
		level = zerolog.WarnLevel
	}

	return log.Logger.Level(level).With().Logger()
}
