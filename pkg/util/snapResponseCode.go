package util

import "github.com/paper-indonesia/pg-mcp-server/constant"

func GenerateResponseCode(snapCode, service string) (string, string) {
	var code, message string

	switch snapCode {
	case constant.SNAP_SUCCESS:
		code = "200" + service + "00"
		message = "Successful"
	case constant.SNAP_INPROGRESS:
		code = "202" + service + "00"
		message = "Request In Progress"
	case constant.SNAP_BAD_REQUEST,
		constant.SNAP_INVALID_FIELD,
		constant.SNAP_INVALID_MANDATORY:

		code, message = generate400Response(snapCode, service)
	case constant.SNAP_UNAUTHORIZED,
		constant.SNAP_INVALID_TOKEN_B2B,
		constant.SNAP_INVALID_CUSTOMER_TOKEN,
		constant.SNAP_TOKEN_NOT_FOUND,
		constant.SNAP_CUSTOMER_TOKEN_NOT_FOUND,
		constant.SNAP_ACCESS_TOKEN_INVALID,
		constant.SNAP_INVALID_SIGNATURE:

		code, message = generate401Response(snapCode, service)
	case constant.SNAP_TRANSACTION_EXPIRED,
		constant.SNAP_FEATURE_NOT_ALLOWED,
		constant.SNAP_EXCEEDS_TRANSACTION_AMOUNT_LIMIT,
		constant.SNAP_SUSPECTED_FRAUD,
		constant.SNAP_ACTIVITY_LIMIT_EXCEEDED,
		constant.SNAP_DO_NOT_HONOR,
		constant.SNAP_FEATURE_NOT_ALLOWED_THIS_TIME,
		constant.SNAP_TRANSACTION_NOT_PERMITTED,
		constant.SNAP_SUSPEND_TRANSACTION,
		constant.SNAP_INACTIVE_ACCOUNT:

		code, message = generate403Response(snapCode, service)
	case constant.SNAP_INVALID_TRANSACTION_STATUS,
		constant.SNAP_TRANSACTION_NOT_FOUND,
		constant.SNAP_INVALID_ROUTING,
		constant.SNAP_BANK_NOT_SUPPORTED,
		constant.SNAP_TRANSACTION_CANCELLED,
		constant.SNAP_INVALID_VA,
		constant.SNAP_INVALID_AMOUNT,
		constant.SNAP_INVALID_ALREADY_PAID,
		constant.SNAP_INVALID_BILL_EXPIRED,
		constant.SNAP_INVALID_QR,
		constant.SNAP_INCONSISTENT_REQUEST:

		code, message = generate404Response(snapCode, service)
	case constant.SNAP_TO_MANY_REQUEST:
		code = "429" + service + "00"
		message = "Too Many Request"
	case constant.SNAP_CONFLICT:
		code = "409" + service + "00"
		message = "Conflict"
	case constant.SNAP_DUPLICATE_PARTNER_REFERENCE_NO:
		code = "409" + service + "01"
		message = "Duplicate Partner Reference no"
	case constant.SNAP_GENERAL_ERROR,
		constant.SNAP_INTERNAL_SERVER_ERROR,
		constant.SNAP_EXTERNAL_SERVER_ERROR:

		code, message = generate500Response(snapCode, service)
	case constant.SNAP_TIMEOUT:
		code = "504" + service + "00"
		message = "Timeout"
	}

	return code, message
}

func generate400Response(snapCode, service string) (string, string) {
	var code, message string
	switch snapCode {
	case constant.SNAP_BAD_REQUEST:
		code = "400" + service + "00"
		message = "Bad Request"
	case constant.SNAP_INVALID_FIELD:
		code = "400" + service + "01"
		message = "Invalid Field Format"
	case constant.SNAP_INVALID_MANDATORY:
		code = "400" + service + "02"
		message = "Missing Mandatory Field"
	}
	return code, message
}

func generate401Response(snapCode, service string) (string, string) {
	var code, message string

	switch snapCode {
	case constant.SNAP_UNAUTHORIZED:
		code = "401" + service + "00"
		message = "Unauthorized."
	case constant.SNAP_INVALID_TOKEN_B2B:
		code = "401" + service + "01"
		message = "Invalid Token (B2B)"
	case constant.SNAP_INVALID_CUSTOMER_TOKEN:
		code = "401" + service + "02"
		message = "Invalid Customer Token"
	case constant.SNAP_TOKEN_NOT_FOUND:
		code = "401" + service + "03"
		message = "Token Not Found (B2B)"
	case constant.SNAP_CUSTOMER_TOKEN_NOT_FOUND:
		code = "401" + service + "04"
		message = "	Customer Token Not Found"
	case constant.SNAP_ACCESS_TOKEN_INVALID:
		code = "401" + service + "01"
		message = "Access Token Invalid"
	case constant.SNAP_INVALID_SIGNATURE:
		code = "401" + service + "00"
		message = "Unauthorized. [Signature]"
	}

	return code, message
}

func generate403Response(snapCode, service string) (string, string) {
	var code, message string
	switch snapCode {
	case constant.SNAP_TRANSACTION_EXPIRED:
		code = "403" + service + "00"
		message = "Transaction Expired"
	case constant.SNAP_FEATURE_NOT_ALLOWED:
		code = "403" + service + "01"
		message = "Feature Not Allowed"
	case constant.SNAP_EXCEEDS_TRANSACTION_AMOUNT_LIMIT:
		code = "403" + service + "02"
		message = "Exceeds Transaction Amount Limit"
	case constant.SNAP_SUSPECTED_FRAUD:
		code = "403" + service + "03"
		message = "Suspected Fraud"
	case constant.SNAP_ACTIVITY_LIMIT_EXCEEDED:
		code = "403" + service + "04"
		message = "Activity Count Limit Exceeded"
	case constant.SNAP_DO_NOT_HONOR:
		code = "403" + service + "05"
		message = "Do Not Honor"
	case constant.SNAP_FEATURE_NOT_ALLOWED_THIS_TIME:
		code = "403" + service + "06"
		message = "Feature Not Allowed At This Time."
	case constant.SNAP_TRANSACTION_NOT_PERMITTED:
		code = "403" + service + "15"
		message = "Transaction not Permitted"
	case constant.SNAP_SUSPEND_TRANSACTION:
		code = "403" + service + "16"
		message = "Suspend Transaction"
	case constant.SNAP_INACTIVE_ACCOUNT:
		code = "403" + service + "17"
		message = "Inactive Account"
	}
	return code, message
}

func generate404Response(snapCode, service string) (string, string) {
	var code, message string

	switch snapCode {
	case constant.SNAP_INVALID_TRANSACTION_STATUS:
		code = "404" + service + "00"
		message = "Invalid Transaction Status"
	case constant.SNAP_TRANSACTION_NOT_FOUND:
		code = "404" + service + "01"
		message = "Transaction Not Found"
	case constant.SNAP_INVALID_ROUTING:
		code = "404" + service + "02"
		message = "Invalid Routing"
	case constant.SNAP_BANK_NOT_SUPPORTED:
		code = "404" + service + "03"
		message = "Bank Not Supported By Switch"
	case constant.SNAP_TRANSACTION_CANCELLED:
		code = "404" + service + "04"
		message = "Transaction Cancelled"
	case constant.SNAP_INVALID_VA:
		code = "404" + service + "12"
		message = "Bill not found"
	case constant.SNAP_INVALID_AMOUNT:
		code = "404" + service + "13"
		message = "Invalid Amount"
	case constant.SNAP_INVALID_ALREADY_PAID:
		code = "404" + service + "14"
		message = "Paid Bill"
	case constant.SNAP_INCONSISTENT_REQUEST:
		code = "404" + service + "18"
		message = "Inconsistent Request"
	case constant.SNAP_INVALID_BILL_EXPIRED:
		code = "404" + service + "19"
		message = "Bill expired"
	}
	return code, message
}

func generate500Response(snapCode, service string) (string, string) {
	var code, message string
	switch snapCode {
	case constant.SNAP_GENERAL_ERROR:
		code = "500" + service + "00"
		message = "General Error"
	case constant.SNAP_INTERNAL_SERVER_ERROR:
		code = "500" + service + "01"
		message = "Internal Server Error"
	case constant.SNAP_EXTERNAL_SERVER_ERROR:
		code = "500" + service + "02"
		message = "External Server Error"
	}
	return code, message
}
