package goff

import (
	"context"
	"fmt"

	"github.com/paper-indonesia/pg-mcp-server/pkg/logger"
	"github.com/thomaspoignant/go-feature-flag/notifier"
	"go.uber.org/zap"
)

type LoggerNotifier struct {
	logger.ILogger
}

func NewLogsNotifier(logger logger.ILogger) notifier.Notifier {
	return &LoggerNotifier{logger}
}

func (l *LoggerNotifier) Notify(diff notifier.DiffCache) error {
	for key := range diff.Deleted {
		l.Info(context.Background(), fmt.Sprintf("flag %v removed\n", key), zap.Any("flag", key))
	}

	for key, value := range diff.Added {
		l.Info(context.Background(), fmt.Sprintf("flag %v added\n", key), zap.Any("flag", key), zap.Any("value", value))
	}

	for key, flagDiff := range diff.Updated {
		if flagDiff.After.IsDisable() != flagDiff.Before.IsDisable() {
			if flagDiff.After.IsDisable() {
				// Flag is disabled
				l.Info(context.Background(), fmt.Sprintf("flag %v is turned OFF\n", key), zap.Any("flag", key))
				continue
			}
			// Flag is enabled
			l.Info(context.Background(), fmt.Sprintf("flag %v is turned ON\n", key), zap.Any("flag", key))
			continue
		}
		// key has changed in cache
		l.Info(context.Background(), fmt.Sprintf("flag %v updated\n", key), zap.Any("flag", key), zap.Any("before", flagDiff.Before), zap.Any("after", flagDiff.After))
	}

	return nil
}
