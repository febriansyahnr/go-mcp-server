package httpRequestExt

import (
	"context"
	"crypto/tls"
	"net/http"

	cb "github.com/paper-indonesia/pdk/go/circuitbreaker"
	pdkConst "github.com/paper-indonesia/pdk/v2/constant"
	pdkLogger "github.com/paper-indonesia/pdk/v2/logger"
	"github.com/paper-indonesia/pg-mcp-server/constant"
	httputil "github.com/paper-indonesia/pg-mcp-server/pkg/util/http"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

type IHTTPRequest interface {
	GET(ctx context.Context, uri string, header map[string]string) ([]byte, int, error)
	POST(ctx context.Context, uri string, data interface{}, header map[string]string) ([]byte, int, error)
}

type HTTPRequest struct {
	cbClient   *cb.CircuitBreaker
	logger     pdkLogger.ILogger
	httpClient *http.Client
}
type HTTPRequestConfig func(client *HTTPRequest)

func New(config ...HTTPRequestConfig) IHTTPRequest {
	client := &HTTPRequest{}
	for _, conf := range config {
		conf(client)
	}

	if client.httpClient == nil {
		client.setupDefaultClient()
	}

	return client
}

func WithLogger(logger pdkLogger.ILogger) HTTPRequestConfig {
	return func(client *HTTPRequest) {
		client.logger = logger
	}
}

func WithHttpClient(httpClient *http.Client) HTTPRequestConfig {
	return func(client *HTTPRequest) {
		client.httpClient = httpClient
	}
}

// POST implements IHTTPRequest.
func (c *HTTPRequest) POST(ctx context.Context, uri string, data interface{}, header map[string]string) ([]byte, int, error) {
	if c.cbClient != nil {
		ctx = context.WithValue(ctx, constant.CtxCircuitBreakerKey, c.cbClient)
	}

	return httputil.RequestHitAPI(ctx, c.httpClient, "POST", uri, data, header)
}

// GET implements IHTTPRequest.
func (c *HTTPRequest) GET(ctx context.Context, uri string, header map[string]string) ([]byte, int, error) {
	if c.cbClient != nil {
		ctx = context.WithValue(ctx, constant.CtxCircuitBreakerKey, c.cbClient)
	}

	return httputil.RequestHitAPI(ctx, c.httpClient, "GET", uri, nil, header)
}

func getTraceId(ctx context.Context) string {
	if ctx.Value(pdkConst.CtxTraceIdKey) != nil {
		return ctx.Value(pdkConst.CtxTraceIdKey).(string)
	}

	return ""
}

func (c *HTTPRequest) setupDefaultClient() {
	customTransport := &http.Transport{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
			RootCAs:            nil,
		},
	}

	c.httpClient = &http.Client{
		Timeout:   constant.DEFAULT_HTTP_TIMEOUT,
		Transport: otelhttp.NewTransport(customTransport),
	}
}
