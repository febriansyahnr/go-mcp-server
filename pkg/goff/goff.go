package goff

import (
	"context"
	"log"
	"time"

	consulExt "github.com/paper-indonesia/pg-mcp-server/pkg/consulExt"
	"github.com/paper-indonesia/pg-mcp-server/pkg/logger"
	ffclient "github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/notifier"
	"github.com/thomaspoignant/go-feature-flag/notifier/slacknotifier"
)

type Config struct {
	PollingInterval             time.Duration
	Logger                      *log.Logger
	Context                     context.Context
	Environment                 string
	FileFormat                  string
	StartWithRetrieverError     bool // Default: false
	Offline                     bool // Default: false
	EvaluationContextEnrichment map[string]interface{}

	ConsulAddr              string
	ConsulConfigPath        string
	SlackWebhookURL         string
	Token                   string
	ExporterSlackWebhookURL string
}

func NewGoff(config Config, logger logger.ILogger) error {
	consulRetriever, err := consulExt.New(config.ConsulAddr, config.ConsulConfigPath, config.Token)
	if err != nil {
		return err
	}

	logsNotifier := NewLogsNotifier(logger)
	slackDataExporter := NewSlackDataExporter(config.ExporterSlackWebhookURL)

	return ffclient.Init(ffclient.Config{
		PollingInterval: config.PollingInterval,
		Logger:          config.Logger,
		Context:         config.Context,
		Environment:     config.Environment,
		Retriever:       consulRetriever,
		Notifiers: []notifier.Notifier{
			&slacknotifier.Notifier{
				SlackWebhookURL: config.SlackWebhookURL,
			},
			logsNotifier,
		},
		FileFormat: "YAML", // Default: YAML
		DataExporter: ffclient.DataExporter{
			FlushInterval:    10 * time.Second,
			MaxEventInMemory: 1000,
			Exporter:         slackDataExporter,
		},
		StartWithRetrieverError:     config.StartWithRetrieverError,
		Offline:                     config.Offline,
		EvaluationContextEnrichment: config.EvaluationContextEnrichment,
	})
}
