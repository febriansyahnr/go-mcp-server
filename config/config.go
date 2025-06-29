package config

type Config struct {
	Environment             string     `mapstructure:"ENVIRONMENT"`
	ServiceName             string     `mapstructure:"SERVICE_NAME"`
	ServiceVersion          string     `mapstructure:"SERVICE_VERSION"`
	Host                    HostConfig `mapstructure:"HOST"`
	Port                    string     `mapstructure:"PORT"`
	DebugSlackWebHook       string     `mapstructure:"DEBUG_SLACK_WEB_HOOK"`
	AlertSlackWebhook       string     `mapstructure:"ALERT_SLACK_WEB_HOOK"`
	AlertFinopsSlackWebhook string     `mapstructure:"ALERT_FINOPS_SLACK_WEB_HOOK"`

	MySQLConfig        MySQLConfig        `mapstructure:"DATABASE"`
	RedisConfig        RedisConfig        `mapstructure:"REDIS"`
	RabbitMQConfig     RabbitMQConfig     `mapstructure:"RABBITMQ"`
	PaperCommunication PaperCommunication `mapstructure:"PAPER_COMMUNICATION"`
	GCSConfig          GCSConfig          `mapstructure:"GCS"`

	DictionaryConfig DictionaryConfig `mapstructure:"DICTIONARY"`

	GetOutboundIPTarget string            `mapstructure:"GET_OUTBOUND_IP_TARGET"`
	FeatureFlagConfig   FeatureFlagConfig `mapstructure:"FEATURE_FLAG"`
	Consul              ConsulConfig      `mapstructure:"CONSUL"`

	Slack SlackConfig `mapstructure:"SLACK"`

	OTLPConfig OTLPConfig  `mapstructure:"OTLP"`
	Vault      VaultConfig `mapstructure:"VAULT"`
}

type HostConfig struct {
	Local      string `mapstructure:"LOCAL"`
	Staging    string `mapstructure:"STAGING"`
	Production string `mapstructure:"PRODUCTION"`
}

type ConsulConfig struct {
	Host       string         `mapstructure:"HOST"`
	ConfigPath string         `mapstructure:"CONFIG_PATH"`
	Key        ConsulKeyValue `mapstructure:"KEY"`
}

type MySQLConfig struct {
	Dialect      string `mapstructure:"DIALECT"`
	Host         string `mapstructure:"HOST"`
	Port         string `mapstructure:"PORT"`
	MaxIdleConns int    `mapstructure:"MAX_IDLE_CONNS"`
	MaxOpenConns int    `mapstructure:"MAX_OPEN_CONNS"`
	MaxIdleTime  int    `mapstructure:"MAX_IDLE_TIME"`
	MaxLifeTime  int    `mapstructure:"MAX_LIFE_TIME"`
	SlaveHost    string `mapstructure:"SLAVE_HOST"`
	SlavePort    string `mapstructure:"SLAVE_PORT"`
}

type PaperCommunication struct {
	BaseURL     string `mapstructure:"BASE_URL"`
	EmailSender string `mapstructure:"EMAIL_SENDER"`
}

type RedisConfig struct {
	Host    string `mapstructure:"HOST"`
	Port    string `mapstructure:"PORT"`
	CacheDB int    `mapstructure:"CACHE_DB"`
}

type RabbitMQConfig struct {
	Host string `mapstructure:"HOST"`
	Port string `mapstructure:"PORT"`
}

type CbsConfig struct {
	Permata       CbConfig `mapstructure:"PERMATA"`
	WalletBackend CbConfig `mapstructure:"WALLET_BACKEND"`
}

type CbConfig struct {
	TIMEOUT_INTERVAL       int   `mapstructure:"TIMEOUT_INTERVAL"`
	MAX_FAILURES           int   `mapstructure:"MAX_FAILURES"`
	OPEN_TO_HALF_OPEN_WAIT int   `mapstructure:"OPEN_TO_HALF_OPEN_WAIT"`
	HALF_OPEN_MAX_SUCCESS  int   `mapstructure:"HALF_OPEN_MAX_SUCCESS"`
	HALF_OPEN_MAX_FAILURES int   `mapstructure:"HALF_OPEN_MAX_FAILURES"`
	RETRY_INTERVALS        []int `mapstructure:"RETRY_INTERVALS"`
}

type DictionaryConfig struct {
	Path string `mapstructure:"PATH"`
}

type ConsulKeyValue struct {
	BINManagement string `mapstructure:"BIN_MANAGEMENT"`
	CBConfig      string `mapstructure:"CB_CONFIG"`
}

type GCSConfig struct {
	ServiceBucketName              string `mapstructure:"SERVICE_BUCKET_NAME"`
	ReconciliationReportFolderName string `mapstructure:"RECONCILIATION_REPORT_FOLDER_NAME"`
}

type SlackConfig struct {
	TransferNotifWebhookUrl string `mapstructure:"TRANSFER_NOTIF_WEBHOOK_URL"`
}

type VaultConfig struct {
	Host   string `mapstructure:"HOST"`
	Engine string `mapstructure:"ENGINE"`
	Path   struct {
		PartnerSecretKeys string `mapstructure:"PARTNER_SECRET_KEYS"`
		Secrets           string `mapstructure:"SECRETS"`
	} `mapstructure:"PATH"`
}

type OTLPConfig struct {
	Host            string           `mapstructure:"HOST"`
	Insecure        bool             `mapstructure:"INSECURE"`
	TLSClientConfig *TLSClientConfig `mapstructure:"TLS_CLIENT_CONFIG"`
}

type TLSClientConfig struct {
	InsecureSkipVerify bool `mapstructure:"INSECURE_SKIP_VERIFY"`
}
