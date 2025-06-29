package goff

import (
	"context"
	"log"

	"github.com/thomaspoignant/go-feature-flag/exporter"
)

type ISlackDataExport interface {
	Export(ctx context.Context, log *log.Logger, events []exporter.FeatureEvent) error
	IsBulk() bool
}

type SlackDataExporter struct {
	WebhookURL string
}

func NewSlackDataExporter(webhookURL string) ISlackDataExport {
	return &SlackDataExporter{
		WebhookURL: webhookURL,
	}
}

func (s *SlackDataExporter) Export(ctx context.Context, log *log.Logger, events []exporter.FeatureEvent) error {
	for _, event := range events {
		// TODO: Send event to slack
		log.Printf("Kind: %s\n", event.Kind)
		log.Printf("ContextKind: %s\n", event.ContextKind)
		log.Printf("UserKey: %s\n", event.UserKey)
		log.Printf("CreationDate: %d\n", event.CreationDate)
		log.Printf("Key: %s\n", event.Key)
		log.Printf("Variation: %s\n", event.Variation)
		log.Printf("Value: %v\n", event.Value)
		log.Printf("Default: %v\n", event.Default)
		log.Printf("Source: %s\n", event.Source)
	}

	return nil
}

func (s *SlackDataExporter) IsBulk() bool {
	return false
}
