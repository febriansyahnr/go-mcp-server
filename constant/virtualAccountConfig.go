package constant

import "errors"

const (
	// Integration Type
	VA_CONFIG_INTEGRATION_TYPE_FACILITATOR = "FACILITATOR"
	VA_CONFIG_INTEGRATION_TYPE_AGGREGATOR  = "AGGREGATOR"
	// Integration Method
	VA_CONFIG_INTEGRATION_METHOD_SERVER = "SERVER"
	VA_CONFIG_INTEGRATION_METHOD_CLIENT = "CLIENT"
)

var (
	ErrInvalidVAConfigType              = errors.New("invalid virtual account config type")
	ErrInvalidVAConfigIntegrationType   = errors.New("invalid virtual account config integration type")
	ErrInvalidVAConfigIntegrationMethod = errors.New("invalid virtual account config integration method")
	ErrVAConfigAlreadyExist             = errors.New("virtual account config already exist")
)
