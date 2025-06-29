package logger

import (
	"context"

	"github.com/bluele/slack"

	"go.uber.org/zap"
)

type ILogger interface {
	Debug(ctx context.Context, msg string, fields ...zap.Field)
	Info(ctx context.Context, msg string, fields ...zap.Field)
	Error(ctx context.Context, msg string, fields ...zap.Field)
	Warn(ctx context.Context, msg string, fields ...zap.Field)
	Panic(ctx context.Context, msg string, fields ...zap.Field)
	CustomSlackNotification(attachment slack.Attachment)
	CustomSlackAlertFinopsNotification(attachment slack.Attachment)
	Sync() error

	GetLogger() *zap.Logger
	CleanupSlackLogger()
}
