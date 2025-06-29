package config

type Secret struct {
	MySQLSecret        MysqlSecret    `mapstructure:"DATABASE"`
	RedisSecret        RedisSecret    `mapstructure:"REDIS"`
	RabbitMQSecret     RabbitMQSecret `mapstructure:"RABBITMQ"`
	SecuritySecret     SecuritySecret `mapstructure:"SECURITY"`
	NewRelicLicenseKey string         `mapstructure:"NEW_RELIC_LICENSE_KEY"`
	InternalServiceKey string         `mapstructure:"INTERNAL_SERVICE_KEY"`
	StatsdHost         string         `mapstructure:"STATSD_HOST"`
	StatsdPort         string         `mapstructure:"STATSD_PORT"`
	CrmSecret          CrmSecret      `mapstructure:"CRM"`
	ConsulSecret       ConsulSecret   `mapstructure:"CONSUL"`
	VaultSecret        VaultSecret    `mapstructure:"VAULT"`
}

type DatabaseSecret struct {
	Database string `mapstructure:"DB_NAME"`
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
}

type MysqlSecret struct {
	SnapCore      DatabaseSecret `mapstructure:"SNAP_CORE"`
	BackendPortal DatabaseSecret `mapstructure:"BACKEND_PORTAL"`
}

type RedisSecret struct {
	Password string `mapstructure:"PASSWORD"`
}

type RabbitMQSecret struct {
	Username string `mapstructure:"USERNAME"`
	Password string `mapstructure:"PASSWORD"`
}

type SecuritySecret struct {
	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

type CrmSecret struct {
	ApiKey string `mapstructure:"API_KEY"`
}

type ConsulSecret struct {
	Token string `mapstructure:"TOKEN"`
}

type VaultSecret struct {
	Token string `mapstructure:"TOKEN"`
}

type WalletBackendSecret struct {
	Token string `mapstructure:"TOKEN"`
}
