package constant

import (
	"encoding/json"
	"time"
)

type (
	constantKey   string
	boolKey       bool
	TMapStr       map[string]string
	TMapAny       map[string]any
	TMapBool      map[string]bool
	TArrAny       []any
	TConfigStatus string
	TReconCode    string
)

func (m *TMapAny) Json() []byte {
	if m == nil {
		return []byte("{}")
	}
	if jsonByte, err := json.Marshal(m); err == nil {
		return jsonByte
	}
	return []byte("{}")
}

const (
	// ctxTraceIdKey is the context key for trace id
	CtxRequestIdKey      constantKey = "request_id"
	CtxTraceIdKey        constantKey = "trace_id"
	CrmIdentifierCtxKey  constantKey = "crm-identifier"
	CtxCircuitBreakerKey constantKey = "circuit-breaker"

	// CtxAcceptLanguage is the context key for dictionary language
	CtxAcceptLanguage string = "accept_language"
	// CtxResponseMessage is the context key for dictionary message
	CtxResponseMessage string = "response_message"

	CtxRabbitMqReplyTo constantKey = "reply_to"

	HeaderAcceptLanguage      = "Accept-Language"
	HeaderAuthorization       = "Authorization"
	HeaderContentType         = "Content-Type"
	HeaderXCRMKey             = "X-CRM-Key"
	HeaderXInternalServiceKey = "X-Internal-Service-Key"
	HeaderXClientKey          = "X-CLIENT-KEY"
	HeaderXTimestamp          = "X-TIMESTAMP"
	HeaderXSignature          = "X-SIGNATURE"
	HeaderXPartnerID          = "X-PARTNER-ID"
	HeaderChannelID           = "CHANNEL-ID"
	HeaderXExternalID         = "X-EXTERNAL-ID"
	HeaderXIdempotencyCheck   = "X-IDEMPOTENCY-CHECK"

	EnvironmentDevelopment = "development"
	EnvironmentLocal       = "local"
	EnvironmentStaging     = "staging"
	EnvironmentProduction  = "production"

	MIMEApplicationJSON     = "application/json"
	MIMEApplicationForm     = "application/x-www-form-urlencoded"
	MIMEApplicationFormData = "multipart/form-data"
)

const (
	DefaultCurrencyIDR string = "IDR"
)

const (
	EntryServiceTransfer string = "TRANSFER"
)

const (
	CtxValueInquiryStatusVa       string = "inquiry_status_va"
	CtxValueCreateVa              string = "create_va"
	CtxValueUpdateVa              string = "update_va"
	CtxValueDeleteVa              string = "delete_va"
	CtxValueInquiryVa             string = "inquiry_va"
	CtxValueCreateQris            string = "create_qris"
	AdditionalInfoFailedReasonKey string = "failedReason"
)

var IgnoreLoggingPath = []string{
	"/health-check",
	"/ping",
}

const (
	TransactionStatusSuccess = "SUCCESS"
	TransactionStatusFailed  = "FAILED"
	TransactionStatusPending = "PENDING"
)

const (
	ReconStatusInvalid        = "INVALID"
	ReconStatusValid          = "VALID"
	ReconStatusTrue           = "TRUE"
	ReconStatusReview         = "REVIEW"
	ReconCodeInvalidAmount    = TReconCode("INVALID_AMOUNT")
	ReconCodeInvalidReference = TReconCode("INVALID_REFERENCE")
	ReconCodeInvalidStatus    = TReconCode("INVALID_STATUS")
	ReconCodeInvalidDate      = TReconCode("INVALID_DATE")
	ReconCodeOk               = TReconCode("OK")

	PayoutAutoReconcileDurationInHour = 3
)

const (
	SubCompanyWallet = "WALLET"
)

const (
	VAStatusDeleted string = "deleted"
)

const (
	TIMER_IDEMPOTENCY_KEY     = "IDEMPOTEN-KEY"
	DEFAULT_LIMITER_THRESHOLD = 10
	DEFAULT_HTTP_TIMEOUT      = 55 * time.Second
	LIMITER_KEY               = "rate-limit:%s:%s"
	TOKEN_KEY_VA              = "b2btoken:%s:va"
	TOKEN_KEY_QRIS            = "b2btoken:%s:qris"
	TOKEN_KEY_DEFAULT         = "b2btoken:%s"
	DEFAULT_TOKEN_EXPIRATION  = 14 * time.Minute
	RECALL_TRANSFER_KEY       = "recall-transfer:%s:%s" // channelService, partnerReferenceNo
)

var AcquirerWithPrefix = []string{
	ACQUIRER_BNI,
}

const (
	MAX_EXTERNAL_ID_LENGTH               = 32
	MAX_QRIS_PARTNER_REFERENCE_NO_LENGTH = 20
)

const (
	CONFIG_STATUS_ACTIVE   TConfigStatus = "ACTIVE"
	CONFIG_STATUS_INACTIVE TConfigStatus = "INACTIVE"
)

// Metadata Keys
const (
	MetadataAdditionalInfoKey = "additionalInfo"
	MetadataNotifyInfoKey     = "notifyInfo"
)
