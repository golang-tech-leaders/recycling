package logger

import (
	"os"
	"recycling/internal/config"
	"strings"

	"github.com/rs/zerolog"
)

type Logger struct {
	zerolog.Logger
}

func New(conf config.LogConfig) *Logger {
	var level zerolog.Level
	switch strings.ToLower(conf.LogLevel) {
	case "debug":
		level = zerolog.DebugLevel
	case "error":
		level = zerolog.ErrorLevel
	case "panic":
		level = zerolog.PanicLevel
	case "fatal":
		level = zerolog.FatalLevel
	case "warning":
		level = zerolog.WarnLevel
	case "info":
		level = zerolog.InfoLevel
	default:
		level = zerolog.ErrorLevel
	}

	zerolog.SetGlobalLevel(level)

	return &Logger{zerolog.New(os.Stdout).With().Timestamp().Logger()}
}
