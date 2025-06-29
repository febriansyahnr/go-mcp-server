package config

type FeatureFlagConfig struct {
	PollingInterval         int    `mapstructure:"POLLING_INTERVAL"`
	FileFormat              string `mapstructure:"FILE_FORMAT"`
	StartWithRetrieverError bool   `mapstructure:"START_WITH_RETRIEVER_ERROR"`
	Offline                 bool   `mapstructure:"OFFLINE"`
	SlackWebhookURL         string `mapstructure:"SLACK_WEBHOOK_URL"`
	ExporterSlackWebhookURL string `mapstructure:"EXPORTER_SLACK_WEBHOOK_URL"`
}
