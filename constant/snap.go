package constant

const (
	SnapDateFormatLayout = "2006-01-02T15:04:05-07:00"
	// va type
	VATypeCloseDynamic = "CLOSED_DYNAMIC"
	VATypeCloseStatic  = "CLOSED_STATIC"
	VATypeOpenStatic   = "OPEN_STATIC"
	// response code

	// 200XX00
	SNAP_SUCCESS = "success" // 200XX00
	// 202XX00
	SNAP_INPROGRESS = "inprogress" // 202XX00
	// 400XX00
	SNAP_BAD_REQUEST       = "bad_request"       // 400XX00
	SNAP_INVALID_FIELD     = "invalid_field"     // 400XX01
	SNAP_INVALID_MANDATORY = "invalid_mandatory" // 400XX02
	// 401XX00
	SNAP_UNAUTHORIZED             = "unauthorized"             // 401XX00
	SNAP_INVALID_SIGNATURE        = "invalid_signature"        // 401XX01
	SNAP_ACCESS_TOKEN_INVALID     = "access_token_invalid"     // 401XX02
	SNAP_INVALID_TOKEN_B2B        = "invalid_token_b2b"        // 401XX01
	SNAP_INVALID_CUSTOMER_TOKEN   = "invalid_customer_token"   // 401XX02
	SNAP_TOKEN_NOT_FOUND          = "token_not_found"          // 401XX03
	SNAP_CUSTOMER_TOKEN_NOT_FOUND = "customer_token_not_found" // 401XX04
	// 403XX00
	SNAP_TRANSACTION_EXPIRED              = "transaction_expired"              // 403XX00
	SNAP_FEATURE_NOT_ALLOWED              = "feature_not_allowed"              // 403XX01
	SNAP_EXCEEDS_TRANSACTION_AMOUNT_LIMIT = "exceeds_transaction_amount_limit" // 403XX02
	SNAP_SUSPECTED_FRAUD                  = "suspected_fraud"                  // 403XX03
	SNAP_ACTIVITY_LIMIT_EXCEEDED          = "activity_limit_exceeded"          // 403XX04
	SNAP_DO_NOT_HONOR                     = "do_not_honor"                     // 403XX05
	SNAP_FEATURE_NOT_ALLOWED_THIS_TIME    = "feature_not_allowed_this_time"    // 403XX06
	SNAP_DORMANT_ACCOUNT                  = "dormant_account"                  // 409XX09
	SNAP_TRANSACTION_NOT_PERMITTED        = "transaction_not_permitted"        // 403XX15
	SNAP_INSUFFICIENT_FUND                = "insufficient_fund"                // 403xx16
	SNAP_INACTIVE_ACCOUNT                 = "inactive_account"                 // 403xx17
	SNAP_SET_LIMIT_NOT_ALLOWED            = "set_limit_not_allowed"
	SNAP_ACCOUNT_LIMIT_EXCEEDED           = "account_limit_exceeded"
	SNAP_SUSPEND_TRANSACTION              = "suspend_transaction"
	// 404XX00
	SNAP_INVALID_TRANSACTION_STATUS = "invalid_transaction_status" // 404XX00
	SNAP_TRANSACTION_NOT_FOUND      = "transaction_not_found"      // 404XX01
	SNAP_INVALID_ROUTING            = "invalid_routing"            // 404XX02
	SNAP_BANK_NOT_SUPPORTED         = "bank_not_supported"         // 404XX03
	SNAP_TRANSACTION_CANCELLED      = "transaction_cancelled"      // 404XX04
	SNAP_INVALID_ACCOUNT            = "invalid_account"
	SNAP_INVALID_VA                 = "invalid_va"                   // 404XX12
	SNAP_INVALID_AMOUNT             = "invalid_amount"               // 404XX13
	SNAP_INVALID_ALREADY_PAID       = "invalid_already_paid"         // 404XX14
	SNAP_INVALID_BILL_EXPIRED       = "invalid_bill_expired"         // 404XX19
	SNAP_INCONSISTENT_REQUEST       = "invalid_inconsistent_request" // 404XX18
	SNAP_INVALID_QR                 = "invalid_qr"                   // 404XX12
	// 409XX00
	SNAP_CONFLICT                       = "conflict"                       // 409XX00
	SNAP_DUPLICATE_PARTNER_REFERENCE_NO = "duplicate_partner_reference_no" // 409xx01
	// 429XX00
	SNAP_TO_MANY_REQUEST = "to_many_request" // 00
	// 500XX00
	SNAP_GENERAL_ERROR         = "general_error"         // 500XX00
	SNAP_INTERNAL_SERVER_ERROR = "internal_server_error" // 500XX01
	SNAP_EXTERNAL_SERVER_ERROR = "external_server_error" // 500XX02
	// 504XX00
	SNAP_TIMEOUT = "timeout" // 00

	// snap service

	SNAP_SERVICE_ACCOUNT_INQUIRY_INTERNAL   = "15"
	SNAP_SERVICE_ACCOUNT_INQUIRY_EXTERNAL   = "16"
	SNAP_SERVICE_INTRABANK_TRANSFER         = "17"
	SNAP_SERVICE_INTERBANK_TRANSFER         = "18"
	SNAP_SERVICE_SKN_TRANSFER               = "23"
	SNAP_SERVICE_INQUIRY                    = "24"
	SNAP_SERVICE_PAYMENT                    = "25"
	SNAP_SERVICE_EWALLET_TO_BANK            = "43"
	SNAP_SERVICE_QRIS                       = "47"
	SNAP_SERVICE_QRIS_PAYMENT               = "52"
	SNAP_SERVICE_DEBIT_PAYMENT_HOST_TO_HOST = "54"
	SNAP_SERVICE_DEBIT_STATUS               = "55"
	SNAP_SERVICE_DEBIT_NOTIFY               = "56"
	SNAP_SERVICE_B2B                        = "73"
	SNAP_SERVICE_PAYMENT_NOTIFY             = "53"

	// VA paymentFlagStatus
	SNAP_VA_SUCCESS_FLAG_STATUS = "00"
	SNAP_VA_REJECT_FLAG_STATUS  = "01"
	VA_PAYMENT_FLAG_STATUS      = "paymentFlagStatus"
	VA_INQUIRY_FLAG_STATUS      = "inquiryFlagStatus"
	VA_VIRTUAL_ACCOUNT_NO       = "virtualAccountNo"
	VA_CUSTOMER_NO              = "customerNo"
	VA_REFERENCE_NO             = "referenceNo"
	VA_TRX_DATE_TIME            = "trxDateTime"
	VA_PARTNER_SERVICE_ID       = "partnerServiceId"
	VA_INQUIRY_REQUEST_ID       = "inquiryRequestId"
	VA_PAYMENT_REQUEST_ID       = "paymentRequestId"
	VA_FLAG_REASON              = "flagReason"
	VA_TOTAL_AMOUNT             = "totalAmount"
	VA_PAID_AMOUNT              = "paidAmount"

	// VA trxType
	VA_TRX_TYPE         = "trxType"
	VA_TRX_TYPE_PAYMENT = "PAYMENT"
	VA_TRX_TYPE_INQUIRY = "INQUIRY"

	// acquirer
	ACQUIRER_PERMATA   = "PERMATA"
	ACQUIRER_ASPI_BANK = "ASPI"
	ACQUIRER_MOCK      = "WIREMOCK"

	ACQUIRER_BRI      = "BRI"
	ACQUIRER_BRI_QRIS = "BRI_QRIS"

	ACQUIRER_MANDIRI         = "MANDIRI"
	ACQUIRER_CIMB            = "CIMB"
	ACQUIRER_BSI             = "BSI"
	ACQUIRER_BNC             = "BNC"
	ACQUIRER_SIMULATION      = "SIMULATION"
	ACQUIRER_C2AMANDIRI      = "C2AMANDIRI"
	ACQUIRER_C2ABRI          = "C2ABRI"
	ACQUIRER_DBS             = "DBS"
	ACQUIRER_BCA             = "BCA"
	ACQUIRER_BNI             = "BNI"
	ACQUIRER_BNI_VA          = "BNI_VA"
	ACQUIRER_MANDIRI_CENTRAL = "MANDIRI_CENTRAL"
	ACQUIRER_OCBC            = "OCBC"
	ACQUIRER_UOB             = "UOB"
	ACQUIRER_DANAMON         = "DANAMON"
	ACQUIRER_SMBC            = "SMBC"
	ACQUIRER_ALTO            = "ALTO"

	// wallet provider
	ACQUIRER_DANA      = "DANA"
	ACQUIRER_SHOPEEPAY = "SHOPEEPAY"
	ACQUIRER_FLIP      = "FLIP"

	SERVICE_TRANSFER       = "TRANSFER"
	SERVICE_QRIS           = "QRIS"
	SERVICE_VA             = "VA"
	SERVICE_GENERAL        = "GENERAL"
	SERVICE_WALLET_BACKEND = "WALLET_BACKEND"
	SERVICE_DIRECT_DEBIT   = "DIRECT_DEBIT"

	// message
	TRANSFER_MESSAGE_INPROGRESS = "Transaction In Progress"
)
