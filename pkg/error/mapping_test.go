package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetErrorType(t *testing.T) {
	testCases := []struct {
		desc     string
		errCode  string
		expected string
	}{
		{
			desc:     "should return API error type for API validation error",
			errCode:  ErrCodeAPIValidation,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for request forbidden error",
			errCode:  ErrCodeRequestForbidden,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for not found error",
			errCode:  ErrCodeNotFound,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for resource not found error",
			errCode:  ErrCodeResourceNotFound,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for resource not complete error",
			errCode:  ErrCodeResourceNotComplete,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for duplicate error",
			errCode:  ErrCodeDuplicate,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for frequency above limit error",
			errCode:  ErrCodeFrequencyAboveLimit,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for database error",
			errCode:  ErrCodeDatabase,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for internal error",
			errCode:  ErrCodeInternal,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for bad gateway error",
			errCode:  ErrCodeBadGateway,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for service unavailable error",
			errCode:  ErrCodeServiceUnavailable,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for gateway timeout error",
			errCode:  ErrCodeGatewayTimeout,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for invalid payment method error",
			errCode:  ErrCodeInvalidPaymentMethod,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for invalid credential error",
			errCode:  ErrCodeInvalidCredential,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for invalid API key error",
			errCode:  ErrCodeInvalidAPIKey,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for payment expired error",
			errCode:  ErrCodePaymentExpired,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for payment cancelled error",
			errCode:  ErrCodePaymentCancelled,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for invalid account details error",
			errCode:  ErrCodeInvalidAccountDetails,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for conflict error",
			errCode:  ErrCodeConflict,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for already success error",
			errCode:  ErrCodeAlreadySuccess,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return API error type for transaction in progress error",
			errCode:  ErrCodeTransactionInProgress,
			expected: ErrTypeAPI,
		},
		{
			desc:     "should return invalid request error type for feature not activated error",
			errCode:  ErrCodeFeatureNotActivated,
			expected: ErrTypeInvalidRequest,
		},
		{
			desc:     "should return invalid request error type for feature not supported error",
			errCode:  ErrCodeFeatureNotSupported,
			expected: ErrTypeInvalidRequest,
		},
		{
			desc:     "should return invalid request error type for invalid request error",
			errCode:  ErrCodeInvalidRequest,
			expected: ErrTypeInvalidRequest,
		},
		{
			desc:     "should return invalid request error type for channel not activated error",
			errCode:  ErrCodeChannelNotActivated,
			expected: ErrTypeInvalidRequest,
		},
		{
			desc:     "should return idempotency error type for idempotency error",
			errCode:  ErrCodeIdempotency,
			expected: ErrTypeIdempotency,
		},
		{
			desc:     "should return PSP error type for channel unavailable error",
			errCode:  ErrCodeChannelUnavailable,
			expected: ErrTypePSP,
		},
		{
			desc:     "should return PSP error type for partner channel error",
			errCode:  ErrCodePartnerChannel,
			expected: ErrTypePSP,
		},
		{
			desc:     "should return PSP error type for payment rejected by channel error",
			errCode:  ErrCodePaymentRejectedByChannel,
			expected: ErrTypePSP,
		},
		{
			desc:     "should return question mark for max amount limit error",
			errCode:  ErrCodeMaxAmountLimit,
			expected: "?",
		},
		{
			desc:     "should return unknown error type for unrecognized error code",
			errCode:  "UNKNOWN_ERROR_CODE",
			expected: ErrTypeUnknown,
		},
		{
			desc:     "should return unknown error type for empty error code",
			errCode:  "",
			expected: ErrTypeUnknown,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := GetErrorType(tc.errCode)
			assert.Equal(t, tc.expected, result)
		})
	}
}