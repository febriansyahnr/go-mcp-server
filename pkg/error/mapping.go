package errors

func GetErrorType(errCode string) string {
	switch errCode {
	case ErrCodeAPIValidation,
		ErrCodeRequestForbidden,
		ErrCodeNotFound,
		ErrCodeResourceNotFound,
		ErrCodeResourceNotComplete,
		ErrCodeDuplicate,
		ErrCodeFrequencyAboveLimit,
		ErrCodeDatabase,
		ErrCodeInternal,
		ErrCodeBadGateway,
		ErrCodeServiceUnavailable,
		ErrCodeGatewayTimeout,
		ErrCodeInvalidPaymentMethod,
		ErrCodeInvalidCredential,
		ErrCodeInvalidAPIKey,
		ErrCodePaymentExpired,
		ErrCodePaymentCancelled,
		ErrCodeInvalidAccountDetails,
		ErrCodeConflict,
		ErrCodeAlreadySuccess,
		ErrCodeTransactionInProgress:
		return ErrTypeAPI
	case ErrCodeFeatureNotActivated, ErrCodeFeatureNotSupported:
		return ErrTypeInvalidRequest
	case ErrCodeInvalidRequest:
		return ErrTypeInvalidRequest
	case ErrCodeIdempotency:
		return ErrTypeIdempotency
	case ErrCodeChannelNotActivated:
		return ErrTypeInvalidRequest
	case ErrCodeChannelUnavailable,
		ErrCodePartnerChannel,
		ErrCodePaymentRejectedByChannel:
		return ErrTypePSP
	case ErrCodeMaxAmountLimit:
		return "?" //  Not yet implemented
	default:
		return ErrTypeUnknown
	}
}
