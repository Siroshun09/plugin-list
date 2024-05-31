package handler

import (
	"github.com/lmittmann/tint"
	"github.com/mattn/go-colorable"
	"log/slog"
	"time"
)

func InitConsoleLogger() {
	slog.SetDefault(slog.New(tint.NewHandler(
		colorable.NewColorableStdout(),
		&tint.Options{
			TimeFormat: time.TimeOnly,
		},
	)))
}
