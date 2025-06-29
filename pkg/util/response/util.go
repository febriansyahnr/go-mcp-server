package response

import (
	"net/http"
	"strconv"

	errors "github.com/paper-indonesia/pg-mcp-server/pkg/error"
)

func GetHTTPStatusCode(errorCode string) (string, int) {
	switch errorCode {
	case errors.ErrCodeInvalidAPIKey,
		errors.ErrCodeInvalidCredential:
		return HttpStatusErrorUnauthorized, http.StatusUnauthorized
	case errors.ErrCodeAPIValidation,
		errors.ErrCodeMaxAmountLimit,
		errors.ErrCodeInvalidAccountDetails,
		errors.ErrCodeInvalidPaymentMethod,
		errors.ErrCodePaymentExpired,
		errors.ErrCodeInvalidRequest:
		return HttpStatusErrorRequest, http.StatusBadRequest
	case errors.ErrCodeRequestForbidden,
		errors.ErrCodeFeatureNotActivated,
		errors.ErrCodeFeatureNotSupported,
		errors.ErrCodeChannelNotActivated,
		errors.ErrCodePaymentRejectedByChannel,
		errors.ErrCodeResourceNotComplete:
		return HttpStatusErrorForbidden, http.StatusForbidden
	case errors.ErrCodeNotFound,
		errors.ErrCodeResourceNotFound:
		return HttpStatusErrorNotFound, http.StatusNotFound
	case errors.ErrCodeDuplicate,
		errors.ErrCodeIdempotency,
		errors.ErrCodePaymentCancelled:
		return HttpStatusErrorDuplicatedCheck, http.StatusConflict
	case errors.ErrCodeFrequencyAboveLimit:
		return HttpStatusErrorLimiter, http.StatusTooManyRequests
	case errors.ErrCodeBadGateway:
		return HttpStatusErrorBadGateway, http.StatusBadGateway
	case errors.ErrCodeServiceUnavailable,
		errors.ErrCodeChannelUnavailable,
		errors.ErrCodePartnerChannel:
		return HttpStatusErrorThirdParty, http.StatusServiceUnavailable
	case errors.ErrCodeGatewayTimeout:
		return HttpStatusErrorThirdParty, http.StatusGatewayTimeout
	case errors.ErrCodeDatabase,
		errors.ErrCodeInternal:
		return HttpStatusErrorInternal, http.StatusInternalServerError
	default:
		return HttpStatusErrorInternal, http.StatusInternalServerError
	}
}

func extractStatusFromResponseCode(errorCode string, responseCode string) (string, int) {
	code, status := GetHTTPStatusCode(errorCode)

	if len(responseCode) > 3 {
		errCode := responseCode[:3]
		// parse string to int
		if statusCode, err := strconv.Atoi(errCode); err == nil {
			status = statusCode
		}
	}

	return code, status
}
