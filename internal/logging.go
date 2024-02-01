package internal

import (
	"log/slog"
	"os"
)

type Logger struct {
	Code int
	Message string
	From string
	Layer string
	TypeLog string
}

func NewLogger(code int, message, from, layer, typeLog string) *slog.Logger {
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