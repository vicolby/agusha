package logging

import (
	"github.com/gimmefear/dswv3/internal/domain"
	"go.uber.org/zap"
)

type ZapLogger struct {
	logger *zap.Logger
}

func NewZapLogger() domain.Logger {
	logger, _ := zap.NewProduction()
	return &ZapLogger{logger: logger}
}

func (z *ZapLogger) Info(msg string, keysAndValues ...interface{}) {
	z.logger.Sugar().Infow(msg, keysAndValues...)
}

func (z *ZapLogger) Error(msg string, keysAndValues ...interface{}) {
	z.logger.Sugar().Errorw(msg, keysAndValues...)
}
