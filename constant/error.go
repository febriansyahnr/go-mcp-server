package constant

import "errors"

var (
	ErrTrxNotFound                               = errors.New("virtual account not found")
	ErrRecordNotFound                            = errors.New("record not found")
	ErrTrxExpired                                = errors.New("virtual account has expired")
	ErrUpdateRefID                               = errors.New("failed to update ref id")
	ErrPaidAmountGreaterThanTotalAmount          = errors.New("paid amount greater than total amount")
	ErrAlreadyPaid                               = errors.New("already paid")
	ErrVANumberInUse                             = errors.New("va number still in use")
	ErrInvalidAmount                             = errors.New("invalid amount")
	ErrInvalidVAAndRefnumber                     = errors.New("invalid va and ref number")
	ErrServiceNotAvailable                       = errors.New("service not available")
	ErrKeyAuthRequired                           = errors.New("key auth required")
	ErrInvalidKey                                = errors.New("invalid key")
	ErrUnauthorized                              = errors.New("unauthorized")
	ErrInconsistentRequest                       = errors.New("inconsistent request")
	ErrXIdentifier                               = errors.New("x-identifier header is missing")
	ErrCannotDeleteVANumber                      = errors.New("virtual account number must be active and have pending status")
	ErrTotalAmountGreaterThanIncomingLimitAmount = errors.New("total amount greater than incoming limit amount")
	// This error is used for flagging error in circuit breaker. So it will increase failed count
	ErrCircuitBreaker = errors.New("flag error circuit breaker")
	// Error transaction already in progress
	ErrTransactionInProgress     = errors.New("transaction already in progress")
	ErrTransactionAlreadySuccess = errors.New("transaction already success")
	ErrUnknownClient             = errors.New("unknown client")
	ErrInternalServer            = errors.New("internal server error")
	ErrConflict                  = errors.New("conflict")
	ErrInvalidMandatoryField     = errors.New("Invalid Mandatory Field")
	ErrInvalidFieldFormat        = errors.New("Invalid Field Format")
)
