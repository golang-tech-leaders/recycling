package logger

import (
	"os"
	"recycling/internal/config"
	"strings"

	"github.com/rs/zerolog"
)

type Logger struct {
	logger zerolog.Logger
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

	customLogger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	return &Logger{logger: customLogger}
}

func (l *Logger) Info(message string) {
	l.logger.Info().Msg(message)
}

func (l *Logger) Debug(message string) {
	l.logger.Debug().Msg(message)
}

func (l *Logger) Fatal(message interface{}) {
	switch message.(type) {
	case nil:
		l.logger.Fatal().Msg("")
	case string:
		l.logger.Fatal().Msg(message.(string))
	case error:
		l.logger.Fatal().Msg(message.(error).Error())
	}
}
