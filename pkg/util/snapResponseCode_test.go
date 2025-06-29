package util

import (
	"testing"

	"github.com/paper-indonesia/pg-mcp-server/constant"
	"github.com/stretchr/testify/require"
)

func TestGenerateResponseCode(t *testing.T) {
	testCases := []struct {
		name         string
		code         string
		service      string
		expectedCode string
	}{
		{
			name:         "Success",
			code:         constant.SNAP_SUCCESS,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "200" + constant.SNAP_SERVICE_INQUIRY + "00",
		},
		{
			name:         "Inprogress",
			code:         constant.SNAP_INPROGRESS,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "202" + constant.SNAP_SERVICE_INQUIRY + "00",
		},
		{
			name:         "Bad Request",
			code:         constant.SNAP_BAD_REQUEST,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "400" + constant.SNAP_SERVICE_INQUIRY + "00",
		},
		{
			name:         "Invalid Field Format",
			code:         constant.SNAP_INVALID_FIELD,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "400" + constant.SNAP_SERVICE_INQUIRY + "01",
		},
		{
			name:         "Missing Mandatory Field",
			code:         constant.SNAP_INVALID_MANDATORY,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "400" + constant.SNAP_SERVICE_INQUIRY + "02",
		},
		{
			name:         "Unauthorized",
			code:         constant.SNAP_UNAUTHORIZED,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "401" + constant.SNAP_SERVICE_INQUIRY + "00",
		},
		{
			name:         "Invalid Token (B2B)",
			code:         constant.SNAP_INVALID_TOKEN_B2B,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "401" + constant.SNAP_SERVICE_INQUIRY + "01",
		},
		{
			name:         "Invalid Customer Token",
			code:         constant.SNAP_INVALID_CUSTOMER_TOKEN,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "401" + constant.SNAP_SERVICE_INQUIRY + "02",
		},
		{
			name:         "Token Not Found (B2B)",
			code:         constant.SNAP_TOKEN_NOT_FOUND,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "401" + constant.SNAP_SERVICE_INQUIRY + "03",
		},
		{
			name:         "Customer Token Not Found",
			code:         constant.SNAP_CUSTOMER_TOKEN_NOT_FOUND,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "401" + constant.SNAP_SERVICE_INQUIRY + "04",
		},
		{
			name:         "Access Token Invalid",
			code:         constant.SNAP_ACCESS_TOKEN_INVALID,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "401" + constant.SNAP_SERVICE_INQUIRY + "01",
		},
		{
			name:         "Unauthorized Signature",
			code:         constant.SNAP_INVALID_SIGNATURE,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "401" + constant.SNAP_SERVICE_INQUIRY + "00",
		},
		{
			name:         "Transaction Expired",
			code:         constant.SNAP_TRANSACTION_EXPIRED,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "403" + constant.SNAP_SERVICE_INQUIRY + "00",
		},
		{
			name:         "Feature Not Allowed",
			code:         constant.SNAP_FEATURE_NOT_ALLOWED,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "403" + constant.SNAP_SERVICE_INQUIRY + "01",
		},
		{
			name:         "Exceeds Transaction Amount Limit",
			code:         constant.SNAP_EXCEEDS_TRANSACTION_AMOUNT_LIMIT,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "403" + constant.SNAP_SERVICE_INQUIRY + "02",
		},
		{
			name:         "Suspected Fraud",
			code:         constant.SNAP_SUSPECTED_FRAUD,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "403" + constant.SNAP_SERVICE_INQUIRY + "03",
		},
		{
			name:         "Activity Count Limit Exceeded",
			code:         constant.SNAP_ACTIVITY_LIMIT_EXCEEDED,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "403" + constant.SNAP_SERVICE_INQUIRY + "04",
		},
		{
			name:         "Do Not Honor",
			code:         constant.SNAP_DO_NOT_HONOR,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "403" + constant.SNAP_SERVICE_INQUIRY + "05",
		},
		{
			name:         "Feature Not Allowed At This Time.",
			code:         constant.SNAP_FEATURE_NOT_ALLOWED_THIS_TIME,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "403" + constant.SNAP_SERVICE_INQUIRY + "06",
		},

		{
			name:         "Invalid Transaction Status",
			code:         constant.SNAP_INVALID_TRANSACTION_STATUS,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "404" + constant.SNAP_SERVICE_INQUIRY + "00",
		},
		{
			name:         "Transaction Not Found",
			code:         constant.SNAP_TRANSACTION_NOT_FOUND,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "404" + constant.SNAP_SERVICE_INQUIRY + "01",
		},
		{
			name:         "Invalid Routing",
			code:         constant.SNAP_INVALID_ROUTING,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "404" + constant.SNAP_SERVICE_INQUIRY + "02",
		},
		{
			name:         "Bank Not Supported By Switch",
			code:         constant.SNAP_BANK_NOT_SUPPORTED,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "404" + constant.SNAP_SERVICE_INQUIRY + "03",
		},
		{
			name:         "Transaction Cancelled",
			code:         constant.SNAP_TRANSACTION_CANCELLED,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "404" + constant.SNAP_SERVICE_INQUIRY + "04",
		},
		{
			name:         "Invalid VA",
			code:         constant.SNAP_INVALID_VA,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "404" + constant.SNAP_SERVICE_INQUIRY + "12",
		},
		{
			name:         "Invalid Amount",
			code:         constant.SNAP_INVALID_AMOUNT,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "404" + constant.SNAP_SERVICE_INQUIRY + "13",
		},
		{
			name:         "Bill has been paid",
			code:         constant.SNAP_INVALID_ALREADY_PAID,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "404" + constant.SNAP_SERVICE_INQUIRY + "14",
		},
		{
			name:         "Bill expired",
			code:         constant.SNAP_INVALID_BILL_EXPIRED,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "404" + constant.SNAP_SERVICE_INQUIRY + "19",
		},
		{
			name:         "Conflict",
			code:         constant.SNAP_CONFLICT,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "409" + constant.SNAP_SERVICE_INQUIRY + "00",
		},
		{
			name:         "Too Many Request",
			code:         constant.SNAP_TO_MANY_REQUEST,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "429" + constant.SNAP_SERVICE_INQUIRY + "00",
		},

		{
			name:         "General Error",
			code:         constant.SNAP_GENERAL_ERROR,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "500" + constant.SNAP_SERVICE_INQUIRY + "00",
		},
		{
			name:         "Internal Server Error",
			code:         constant.SNAP_INTERNAL_SERVER_ERROR,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "500" + constant.SNAP_SERVICE_INQUIRY + "01",
		},
		{
			name:         "External Server Error",
			code:         constant.SNAP_EXTERNAL_SERVER_ERROR,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "500" + constant.SNAP_SERVICE_INQUIRY + "02",
		},

		{
			name:         "Timeout",
			code:         constant.SNAP_TIMEOUT,
			service:      constant.SNAP_SERVICE_INQUIRY,
			expectedCode: "504" + constant.SNAP_SERVICE_INQUIRY + "00",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			code, _ := GenerateResponseCode(tc.code, tc.service)
			require.Equal(t, tc.expectedCode, code)
		})
	}
}
