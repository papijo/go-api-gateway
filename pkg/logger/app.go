package logger

import (
	"io"
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/rs/zerolog"
)

var (
	Logger      zerolog.Logger
	ErrorLogger zerolog.Logger
	DebugLogger zerolog.Logger
)

func init() {

	requestLogFile := &lumberjack.Logger{
		Filename:   "logs/requests/api-gateway.log",
		MaxSize:    10,
		MaxAge:     7,
		MaxBackups: 10,
		Compress:   true,
	}

	errorLogFile := &lumberjack.Logger{
		Filename:   "logs/errors/api-gateway-errors.log",
		MaxSize:    10,
		MaxAge:     7,
		MaxBackups: 10,
		Compress:   true,
	}

	requestMultiWriter := io.MultiWriter(os.Stdout, requestLogFile)
	errorMultiWriter := io.MultiWriter(os.Stdout, errorLogFile)
	consoleWriter := io.Writer(os.Stdout)

	zerolog.TimestampFieldName = "time"
	zerolog.TimeFieldFormat = "2006-01-02T15:04:05Z07:00"

	Logger = zerolog.New(requestMultiWriter).With().Timestamp().Logger()
	ErrorLogger = zerolog.New(errorMultiWriter).With().Timestamp().Logger()
	DebugLogger = zerolog.New(consoleWriter).With().Timestamp().Logger()
}
