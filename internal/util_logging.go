package internal

import (
	"log/slog"
	"os"
)

type Logger struct {
	Code int `json:"code"`
	Message string `json:"message"`
	From string `json:"from"`
	Layer string `json:"layer"`
	TypeLog string `json:"type_log"`
}

func NewLoggerError(code int, message, from, layer, typeLog string) *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Error(
		"ERROR",
		"code:", code,
		"message:", message,
		"from:", from,
		"layer:", layer,
		"type:", typeLog,
	)

	return logger
}

func NewLoggerInfo(code int, message, from, layer, typeLog string) *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info(
		"INFO",
		"code:", code,
		"message:", message,
		"from:", from,
		"layer:", layer,
		"type:", typeLog,
	)

	return logger
}

func NewLoggerWarning(code int, message, from, layer, typeLog string) *slog.Logger {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Warn(
		"WARNING",
		"code:", code,
		"message:", message,
		"from:", from,
		"layer:", layer,
		"type:", typeLog,
	)

	return logger
}