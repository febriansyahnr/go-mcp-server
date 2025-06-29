package constant

type tableType string

const (
	// CtxNewRelicTxnKey is the context key for newrelic app
	CtxNewRelicTxnKey constantKey = "newrelic_txn"
	// CtxSQLTableNameKey is the context key for sql table name
	CtxSQLTableNameKey tableType = "table_name"
	// CtxRabbitMQStartTime is the context key for rabbitmq start time
	CtxRabbitMQStartTime string = "start_time"
	CtxFeatureFrom       string = "feature_from"
	Controller                  = "controller"
	Service                     = "service"
	Repository                  = "repository"
	// auto recon
	ProcessorName       = "SNAP_CORE_PROCESSOR"
	TransactionTypeVA   = "VA"
	TransactionTypeQRIS = "QRIS"

	// other processor
	CreditCardCoreProcessor = "CREDIT_CARD_CORE_PROCESSOR"
	XBCoreProcessor         = "XB_CORE_PROCESSOR"
	BillCoreProcessor       = "BILL_CORE_PROCESSOR"
	FlipPGProcessor         = "FLIP_PG_PROCESSOR"
)
