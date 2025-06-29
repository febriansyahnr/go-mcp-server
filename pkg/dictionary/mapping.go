package dictionary

import errors "github.com/paper-indonesia/pg-mcp-server/pkg/error"

func GetTranslationCode(code string) string {
	switch code {
	case errors.ErrCodeInvalidCredential:
		return TranslationErrInvalidCredentials
	case errors.ErrCodeInvalidAPIKey:
		return TranslationErrInvalidAPIKey
	case errors.ErrCodeAPIValidation:
		return TranslationErrAPIValidation
	case errors.ErrCodeRequestForbidden:
		return TranslationErrRequestForbidden
	case errors.ErrCodeFeatureNotActivated:
		return TranslationErrFeatureNotActivated
	case errors.ErrCodeFeatureNotSupported:
		return TranslationErrFeatureNotSupported
	case errors.ErrCodeNotFound:
		return TranslationErrNotFound
	case errors.ErrCodeResourceNotFound:
		return TranslationErrResourceNotFound
	case errors.ErrCodeResourceNotComplete:
		return TranslationErrResourceNotComplete
	case errors.ErrCodeDuplicate:
		return TranslationErrDuplicate
	case errors.ErrCodeIdempotency:
		return TranslationErrIdempotency
	case errors.ErrCodeFrequencyAboveLimit:
		return TranslationErrFrequencyAboveLimit
	case errors.ErrCodeDatabase:
		return TranslationErrDatabase
	case errors.ErrCodeInternal:
		return TranslationErrInternal
	case errors.ErrCodeBadGateway:
		return TranslationErrBadGateway
	case errors.ErrCodeServiceUnavailable:
		return TranslationErrServiceUnavailable
	case errors.ErrCodeGatewayTimeout:
		return TranslationErrGatewayTimeout
	case errors.ErrCodeChannelNotActivated:
		return TranslationErrChannelNotActivated
	case errors.ErrCodeChannelUnavailable:
		return TranslationErrChannelUnavailable
	case errors.ErrCodeMaxAmountLimit:
		return TranslationErrMaxAmountLimit
	case errors.ErrCodePartnerChannel:
		return TranslationErrPartnerChannel
	case errors.ErrCodeInvalidAccountDetails:
		return TranslationErrInvalidAccountDetails
	case errors.ErrCodeInvalidPaymentMethod:
		return TranslationErrInvalidPaymentMethod
	case errors.ErrCodePaymentRejectedByChannel:
		return TranslationErrPaymentRejectedByChannel
	case errors.ErrCodePaymentExpired:
		return TranslationErrPaymentExpired
	case errors.ErrCodePaymentCancelled:
		return TranslationErrPaymentCancelled
	case errors.ErrCodeInsufficientBalance:
		return TranslationErrInsufficientBalance
	case errors.ErrCodeConflict:
		return TranslationErrConflict
	case errors.ErrCodeAlreadySuccess:
		return TranslationErrAlreadySuccess
	default:
		return TranslationErrInternal
	}
}
