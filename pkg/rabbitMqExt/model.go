package rabbitMqExt

import (
	"bytes"
	"sync"
)

/*
single use: true -> dynamic
single use: false -> static (close, open)

amount: close -> amount already defined
amount: open -> di isi manual pas bayar
*/

type RabbitMqExchangeConfig struct {
	Exchange   string
	QueueName  string
	RoutingKey string
	Type       string

	DLQExchange  string
	DLQQueueName string
}

type once struct {
	directReplyTo *sync.Map
}

var rabbitMqExchangeConfig = map[string]RabbitMqExchangeConfig{
	DOCUMENT_TYPE_VA_PAYMENT: {
		DEFAULT_EXCHANGE_NAME,
		"q.snap.va.payment",
		"snap.va.payment",
		EXCHANGE_TYPE_TOPIC,
		"dle.snap-core",
		"dlq.snap.va.payment",
	},
	DOCUMENT_TYPE_VA_STATUS_SCHEDULED: {
		DELAYED_SCHEDULER_EXCHANGE_NAME,
		"q.snap.va.status.scheduled",
		"snap.va.status.scheduled",
		EXCHANGE_TYPE_DELAYED,
		"dle.snap-delayed-scheduler",
		"dlq.snap.va.status.scheduled",
	},
	DOCUMENT_TYPE_QRIS_REGISTRATION_CALLBACK: {
		QRIS_REGISTRATION_CALLBACK_EXCHANGE_NAME,
		"q.snap.qris.registration-callback",
		"snap.qris.registration-callback",
		EXCHANGE_TYPE_TOPIC,
		"dle.snap-core.qris.registration-callback",
		"dlq.snap.qris.registration-callback",
	},
	DOCUMENT_TYPE_SLACK_POST_WEBHOOK: {
		SlackExchange,
		SlackPostWebhookQueueName,
		SlackPostWebhookRoutingKey,
		EXCHANGE_TYPE_TOPIC,
		SlackPostWebhookDLQExchange,
		SlackPostWebhookDLQQueueName,
	},
	DOCUMENT_TYPE_QRIS_PAYMENT_CALLBACK: {
		DEFAULT_EXCHANGE_NAME,
		"q.snap.qris.payment",
		"snap.qris.payment",
		EXCHANGE_TYPE_TOPIC,
		"dle.snap-core",
		"dlq.snap.qris.payment",
	},
	DOCUMENT_TYPE_QRIS_ISSUING_PAYMENT_CALLBACK: {
		DEFAULT_EXCHANGE_NAME,
		SnapQrisIssuingPaymentQueueName,
		SnapQrisIssuingRoutingKey,
		EXCHANGE_TYPE_TOPIC,
		QrisIssuingPaymentDLExchangeName,
		QrisIssuingPaymentDLQueueName,
	},
	DOCUMENT_TYPE_QRIS_STATUS_SCHEDULED: {
		DELAYED_SCHEDULER_EXCHANGE_NAME,
		"q.snap.qris.status.scheduled",
		"snap.qris.status.scheduled",
		EXCHANGE_TYPE_DELAYED,
		"dle.snap-delayed-scheduler",
		"dlq.snap.qris.status.scheduled",
	},
	DOCUMENT_TYPE_TRANSFER_ROUTING: {
		TransferRouteExchange,
		TransferRouteQueueName,
		TransferRouteRoutingKey,
		EXCHANGE_TYPE_TOPIC,
		TransferRouteDLQExchange,
		TransferRouteDLQQueueName,
	},
	DOCUMENT_TYPE_UPDATE_TRANSFER_STATUS: {
		TransferStatusExchange,
		TransferStatusQueueName,
		TransferStatusRoutingKey,
		EXCHANGE_TYPE_TOPIC,
		TransferStatusDLQExchange,
		TransferStatusDLQQueueName,
	},
	DOCUMENT_TYPE_VA_PAYMENT_WALLET: {
		DEFAULT_EXCHANGE_NAME,
		"q.wallet.va.payment",
		"snap.va.payment",
		EXCHANGE_TYPE_TOPIC,
		"dle.snap-core",
		"dlq.wallet.va.payment",
	},
	DOCUMENT_TYPE_EWALLET_PAYMENT: {
		DEFAULT_EXCHANGE_NAME,
		EwalletPaymentQueueName,
		EWALLET_PAYMENT_ROUTING_KEY,
		EXCHANGE_TYPE_TOPIC,
		EwalletPaymentDLQExchange,
		EwalletPaymentDLQQueueName,
	},
	DOCUMENT_TYPE_CHECK_TRANSFER_STATUS: {
		CheckTransferStatusExchange,
		CheckTransferStatusQueueName,
		CheckTransferStatusKey,
		EXCHANGE_TYPE_DELAYED,
		CheckTransferStatusDLQExchange,
		CheckTransferStatusDLQQueueName,
	},
	DOCUMENT_TYPE_CHECK_QRIS_ISSUING_STATUS: {
		CheckQrisIssuingStatusExchange,
		CheckQrisIssuingStatusQueueName,
		CheckQrisIssuingStatusKey,
		EXCHANGE_TYPE_DELAYED,
		CheckQrisIssuingStatusDLQExchange,
		CheckQrisIssuingStatusDLQQueueName,
	},
	DOCUMENT_TYPE_TRANSFER_RECONCILE: {
		TransferReconcileExchange,
		TransferReconcileQueueName,
		TransferReconcileRoutingKey,
		EXCHANGE_TYPE_TOPIC,
		TransferReconcileDLQExchange,
		TransferReconcileDLQQueueName,
	},
	DOCUMENT_TYPE_INQUIRY_ACCOUNT_CALLBACK: {
		InquiryAccountCallbackExchange,
		InquiryAccountCallbackQueueName,
		InquiryAccountCallbackRoutingKey,
		EXCHANGE_TYPE_TOPIC,
		"",
		"",
	},
}

var pool = &sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func setRabbitMqConfig(document string) RabbitMqExchangeConfig {
	var result RabbitMqExchangeConfig

	check := rabbitMqExchangeConfig[document]
	if check.Exchange == "" {
		return result
	}

	return check
}
