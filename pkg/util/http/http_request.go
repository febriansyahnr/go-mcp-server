package httputil

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net"
	"net/http"
	"strings"
	"sync"
	"syscall"

	pdkConst "github.com/paper-indonesia/pdk/v2/constant"
	"github.com/paper-indonesia/pg-mcp-server/constant"
	"go.opentelemetry.io/otel"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

var (
	httpClient *http.Client
	tracer     = otel.Tracer("httpRequest")
	once       sync.Once
	msgTimeout = `{"responseCode": "504xx00", "responseMessage": "%s"}`
)

func RequestHitAPI(
	ctx context.Context,
	httpClient *http.Client,
	method string,
	uri string,
	data interface{},
	header map[string]string,
) (
	res []byte,
	code int,
	err error,
) {
	var (
		tags     = map[string]interface{}{}
		request  *http.Request
		response *http.Response
	)

	ctx, span := tracer.Start(ctx, "util/http/RequestHitAPI",
		trace.WithAttributes(attribute.KeyValue{
			Key:   "uri",
			Value: attribute.StringValue(uri),
		}),
		trace.WithAttributes(attribute.KeyValue{
			Key:   "method",
			Value: attribute.StringValue(method),
		}))

	defer span.End()

	defer func() {
		// record error
		span.SetAttributes(attribute.KeyValue{
			Key:   "responseCode",
			Value: attribute.IntValue(code),
		})
		if err != nil {
			span.SetStatus(codes.Error, err.Error())
			span.RecordError(err)
		}
	}()

	contentType := strings.ToLower(header[constant.HeaderContentType])
	switch {
	case strings.Contains(contentType, constant.MIMEApplicationFormData):
		buf, ok := data.(*bytes.Buffer)
		if !ok {
			return
		}
		request, err = http.NewRequest(method, uri, buf)
		if err != nil {
			return res, code, err
		}
		request.Header.Set(constant.HeaderContentType, header[constant.HeaderContentType])
	case strings.Contains(contentType, constant.MIMEApplicationForm):
		// Accept *bytes.Buffer, *strings.Reader, or io.Reader for form-encoded data
		var bodyReader io.Reader
		switch v := data.(type) {
		case *bytes.Buffer:
			bodyReader = v
		case *strings.Reader:
			bodyReader = v
		case io.Reader:
			bodyReader = v
		case string:
			bodyReader = strings.NewReader(v)
		default:
			return
		}
		request, err = http.NewRequest(method, uri, bodyReader)
		if err != nil {
			return res, code, err
		}
		request.Header.Set(constant.HeaderContentType, header[constant.HeaderContentType])
	default:
		request, err = AssertTypeRequest(data, method, uri)
		if err != nil {
			return res, code, err
		}
		if header[constant.HeaderContentType] == "" {
			request.Header.Set(constant.HeaderContentType, constant.MIMEApplicationJSON)
		}
	}

	for k, v := range header {
		request.Header[k] = []string{v}
	}

	response, err = httpClient.Do(request)

	if err != nil {
		if err, ok := err.(net.Error); ok && err.Timeout() {
			return []byte(`{"responseCode": "504xx00", "responseMessage": "Timeout, action cancelled"}`), 504, err
		}

		if errors.Is(err, syscall.ECONNRESET) {
			return []byte(fmt.Sprintf(msgTimeout, "ECONNRESET, connection reset by peer")), 504, err
		}

		if res == nil {
			res = []byte(fmt.Sprintf(msgTimeout, "PARTNER_ERROR "+err.Error()))
		}

		return res, code, err
	}

	defer response.Body.Close()

	code = response.StatusCode

	res, err = io.ReadAll(response.Body)
	if err != nil {
		return res, code, err
	}

	traceID, ok := ctx.Value(pdkConst.CtxTraceIdKey).(string)
	if !ok {
		traceID = ""
	}

	requestID, ok := ctx.Value(pdkConst.CtxRequestIdKey).(string)
	if !ok {
		requestID = ""
	}

	var responseMap constant.TMapAny
	_ = json.Unmarshal(res, &responseMap)
	slog.Info("http log:",
		slog.String("req_method", request.Method),
		slog.String("req_url", request.URL.String()),
		slog.Any("req_header", request.Header),
		slog.Any("req_body", data),
		slog.Any("response", responseMap),
		slog.String("trace_id", traceID),
		slog.String("request_id", requestID),
	)

	tags["http_response"] = res
	tags["http_status_code"] = code

	if isHttpError := code != http.StatusOK && code != http.StatusCreated; isHttpError {
		var errRes map[string]interface{}
		err := json.Unmarshal(res, &errRes)
		if err != nil && code == http.StatusGatewayTimeout {
			return []byte(`{"responseCode": "504xx00", "responseMessage": "Timeout, action cancelled"}`), 504, err
		}
		return res, code, err
	}

	return res, code, err
}

func AssertTypeRequest(data interface{}, method string, uri string) (request *http.Request, err error) {
	if data == nil {
		request, err = http.NewRequest(method, uri, nil)
		return
	}

	paramReq, _ := json.Marshal(data)

	// check if data is actually a json formatted
	err = json.Unmarshal(paramReq, &map[string]interface{}{})
	if err != nil {
		// if the data is failed to marshal, it's probably already a byte, e.g pgp armored data
		request, err = http.NewRequest(method, uri, bytes.NewBuffer(data.([]byte)))
		return
	}
	request, err = http.NewRequest(method, uri, bytes.NewBuffer(paramReq))
	return
}

func NewHttpRequest(
	ctx context.Context,
	method string,
	url string,
	header map[string]string,
	bodyReq interface{},
) (*http.Request, error) {
	ctx, span := tracer.Start(ctx, "util/http/NewHttpRequest")
	defer span.End()

	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, err
	}

	// Set or modify the request body
	if bodyReq != nil {
		// For handling form data
		if formData, ok := bodyReq.(*bytes.Buffer); ok {
			request.Body = io.NopCloser(formData)
		} else {
			payloadBody, _ := json.Marshal(bodyReq)
			request.Body = io.NopCloser(bytes.NewBuffer(payloadBody))
		}
	}

	for k, v := range header {
		request.Header.Add(k, v)
	}

	return request, nil
}
