package httpRequestExt

import (
	"net/http"
	"time"

	cb "github.com/paper-indonesia/pdk/go/circuitbreaker"
)

const CircuitBreakerKey = "circuit-breaker:"

type CBFallbackFn func(*http.Request) *cb.CircuitBreakerResponse

func cbDefaultFallbackFunc(req *http.Request) *cb.CircuitBreakerResponse {
	// This is where you define your fallback logic. For example, return a static response or call an alternative service.
	// The following is a simple static response for demonstration purposes.
	return &cb.CircuitBreakerResponse{
		HttpStatus:   http.StatusServiceUnavailable,
		ResponseType: cb.Fallback,
		Data: map[string]interface{}{
			"message": "This is a fallback response due to circuit breaker open state.",
		},
	}
}

func WithCircuitBreaker(name string, config cb.Config, retryInterval []int, cacheClient cb.RedisClient, fallbackFn CBFallbackFn) HTTPRequestConfig {
	return func(client *HTTPRequest) {
		for _, ri := range retryInterval {
			config.RetryIntervals = append(config.RetryIntervals, time.Duration(ri)*time.Millisecond)
		}

		cbc := cb.NewCircuitBreaker(config, CircuitBreakerKey+name, cacheClient)
		if fallbackFn != nil {
			cbc.SetFallbackFunc(fallbackFn)
		} else {
			cbc.SetFallbackFunc(cbDefaultFallbackFunc)
		}
		client.cbClient = cbc
	}
}
