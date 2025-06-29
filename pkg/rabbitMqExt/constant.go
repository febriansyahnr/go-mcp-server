package rabbitMqExt

const (
	// exchange
	SlackExchange                  = "snap-core.slack"
	TransferRouteExchange          = "snap-core.transfer"
	TransferStatusExchange         = "snap-core.transfer-status"
	TransferReconcileExchange      = "snap-core.transfer-reconcile"
	CheckTransferStatusExchange    = "snap-core.check-status"
	CheckQrisIssuingStatusExchange = "snap-core.qris-issuing.check-status"
	InquiryAccountCallbackExchange = "snap-core.inquiry-account.callback"

	// routing key
	SlackPostWebhookRoutingKey       = "snap-core.slack.post-webhook"
	TransferRouteRoutingKey          = "snap-core.transfer.routing"
	TransferStatusRoutingKey         = "snap-core.transfer.status"
	TransferReconcileRoutingKey      = "snap-core.transfer.reconcile"
	CheckTransferStatusKey           = "snap-core.transfer.check-status"
	CheckQrisIssuingStatusKey        = "snap-core.qris-issuing.check-status"
	SnapQrisIssuingRoutingKey        = "snap-core.qris-issuing.payment"
	InquiryAccountCallbackRoutingKey = "snap-core.inquiry-account.callback"

	// queue name
	SlackPostWebhookQueueName       = "q.snap-core.slack.post-webhook"
	TransferRouteQueueName          = "q.snap-core.transfer.routing"
	TransferStatusQueueName         = "q.snap-core.transfer.status"
	TransferReconcileQueueName      = "q.snap-core.transfer.reconcile"
	CheckTransferStatusQueueName    = "q.snap-core.transfer.check-status"
	CheckQrisIssuingStatusQueueName = "q.snap-core.qris-issuing.check-status"
	SnapQrisIssuingPaymentQueueName = "q.snap-core.qris-issuing.payment"
	EwalletPaymentQueueName         = "q.snap.ewallet.payment"
	InquiryAccountCallbackQueueName = "q.snap-core.inquiry-account.callback"

	// DLQ
	SlackPostWebhookDLQExchange        = "dle.snap-core.slack.post-webhook"
	SlackPostWebhookDLQQueueName       = "dlq.snap-core.slack.post-webhook"
	TransferRouteDLQExchange           = "dle.snap-core.transfer.routing"
	TransferRouteDLQQueueName          = "dlq.snap-core.transfer.routing"
	TransferStatusDLQExchange          = "dle.snap-core.transfer.status"
	TransferStatusDLQQueueName         = "dlq.snap-core.transfer.status"
	CheckTransferStatusDLQExchange     = "dle.snap-core.transfer.check-status"
	CheckQrisIssuingStatusDLQExchange  = "dle.snap-core.qris-issuing.check-status"
	CheckTransferStatusDLQQueueName    = "dlq.snap-core.transfer.check-status"
	CheckQrisIssuingStatusDLQQueueName = "dlq.snap-core.qris-issuing.check-status"
	QrisIssuingPaymentDLExchangeName   = "dle.snap-core.qris-issuing.payment"
	QrisIssuingPaymentDLQueueName      = "dlq.snap-core.qris-issuing.payment"
	TransferReconcileDLQExchange       = "dle.snap-core.transfer.reconcile"
	TransferReconcileDLQQueueName      = "dlq.snap-core.transfer.reconcile"
	EwalletPaymentDLQExchange          = "dle.snap-core"
	EwalletPaymentDLQQueueName         = "dlq.snap.ewallet.payment"
	InquiryAccountCallbackDLQExchange  = "dle.snap-core.inquiry-account.callback"
	InquiryAccountCallbackDLQQueueName = "dlq.snap-core.inquiry-account.callback"
)

const (
	HeaderTraceId = "trace-id"
)

type contextRabbit string

const (
	CTX_SCHEDULED_DELAY contextRabbit = "delayTime"

	DOCUMENT_TYPE_VA_PAYMENT                    = "va.payment"
	DOCUMENT_TYPE_VA_PAYMENT_WALLET             = "va.payment.wallet"
	DOCUMENT_TYPE_EWALLET_PAYMENT               = "ewallet.payment"
	DOCUMENT_TYPE_VA_STATUS_SCHEDULED           = "va.status.scheduled"
	DOCUMENT_TYPE_QRIS_REGISTRATION_CALLBACK    = "qris.registration-callback"
	DOCUMENT_TYPE_QRIS_PAYMENT_CALLBACK         = "qris.payment-callback"
	DOCUMENT_TYPE_QRIS_ISSUING_PAYMENT_CALLBACK = "qris-issuing.payment-callback"
	DOCUMENT_TYPE_SLACK_POST_WEBHOOK            = "slack.post-webhook"
	DOCUMENT_TYPE_QRIS_STATUS_SCHEDULED         = "qris.status.scheduled"
	DOCUMENT_TYPE_TRANSFER_ROUTING              = "transfer.routing"
	DOCUMENT_TYPE_UPDATE_TRANSFER_STATUS        = "transfer.status"
	DOCUMENT_TYPE_CHECK_TRANSFER_STATUS         = "transfer.check-status"
	DOCUMENT_TYPE_CHECK_QRIS_ISSUING_STATUS     = "qris.issuing.check-status"
	DOCUMENT_TYPE_TRANSFER_RECONCILE            = "transfer.reconcile"
	DOCUMENT_TYPE_INQUIRY_ACCOUNT_CALLBACK      = "inquiry-account.callback"

	// exchange
	DEFAULT_EXCHANGE_NAME                    = "snap-core"
	DELAYED_SCHEDULER_EXCHANGE_NAME          = "snap-delayed-scheduler"
	QRIS_REGISTRATION_CALLBACK_EXCHANGE_NAME = "snap-core-qris-registration-callback"

	// routing key
	VA_PAYMENT_ROUTING_KEY                 = "snap.va.payment"
	VA_UPDATE_STATUS_SCHEDULED_ROUTING_KEY = "snap.va.status.scheduled"
	QRIS_REGISTRATION_CALLBACK_ROUTING_KEY = "snap.qris.registration-callback"
	EWALLET_PAYMENT_ROUTING_KEY            = "snap.ewallet.payment"
)

const (
	EXCHANGE_TYPE_TOPIC   = "topic"
	EXCHANGE_TYPE_DELAYED = "x-delayed-message"
)

const ReplyToQueueName = "amq.rabbitmq.reply-to"
