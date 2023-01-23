package logger

import (
	"context"

	"go.uber.org/zap"
)

type LoggerKeyType string

var logKey LoggerKeyType = "LOGGER"

func Inject(ctx context.Context, log *zap.Logger) context.Context {
	return context.WithValue(ctx, logKey, log)
}

func GetLoggerFromContext(ctx context.Context) *zap.Logger {
	c, _ := ctx.Value(logKey).(*zap.Logger)
	return c
}
