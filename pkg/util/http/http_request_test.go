package httputil_test

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	httputil "github.com/paper-indonesia/pg-mcp-server/pkg/util/http"
)

func TestRequestHitAPI(t *testing.T) {
	tests := []struct {
		name               string
		method             string
		uri                string
		data               interface{}
		header             map[string]string
		mockResponseStatus int
		mockResponseBody   string
		wantCode           int
		wantErr            bool
	}{
		{
			name:               "Successful GET Request",
			method:             "GET",
			uri:                "/success",
			data:               nil,
			header:             map[string]string{"Authorization": "Bearer token123", "Custom-Header": "Value123"},
			mockResponseStatus: http.StatusOK,
			mockResponseBody:   `{"status":"ok"}`,
			wantCode:           http.StatusOK,
			wantErr:            false,
		},
		{
			name:               "Successful POST Request",
			method:             "POST",
			uri:                "/create",
			data:               map[string]string{"Authorization": "Bearer token123", "Custom-Header": "Value123"},
			header:             nil,
			mockResponseStatus: http.StatusCreated,
			mockResponseBody:   `{"status":"created"}`,
			wantCode:           http.StatusCreated,
			wantErr:            false,
		},
		{
			name:               "Network Failure",
			method:             "GET",
			uri:                "/network-failure",
			data:               nil,
			header:             nil,
			mockResponseStatus: 0,
			mockResponseBody:   "",
			wantCode:           0,
			wantErr:            true,
		},
		{
			name:               "Error Unmarshalling JSON",
			method:             "GET",
			uri:                "/json-error",
			data:               nil,
			header:             nil,
			mockResponseStatus: http.StatusBadRequest,
			mockResponseBody:   "{invalid-json}",
			wantCode:           http.StatusBadRequest,
			wantErr:            true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(tt.mockResponseStatus)
				w.Write([]byte(tt.mockResponseBody))
			}))
			defer mockServer.Close()

			ctx := context.Background()
			ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			clientHttp := &http.Client{}

			res, code, err := httputil.RequestHitAPI(ctx, clientHttp, tt.method, mockServer.URL+tt.uri, tt.data, tt.header)

			if (err != nil) != tt.wantErr {
				t.Errorf("%s: RequestHitAPI() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			}
			if code != tt.wantCode {
				t.Errorf("%s: RequestHitAPI() code = %v, want %v", tt.name, code, tt.wantCode)
			}
			if tt.mockResponseBody != "" && !strings.Contains(string(res), tt.mockResponseBody) {
				t.Errorf("%s: RequestHitAPI() response = %v, want contains %v", tt.name, string(res), tt.mockResponseBody)
			}
		})
	}
}

func TestNewHttpRequest(t *testing.T) {
	testCases := []struct {
		name          string
		method        string
		url           string
		header        map[string]string
		bodyReq       interface{}
		expectedError bool
	}{
		{
			name:          "Nil bodyReq",
			method:        "GET",
			url:           "http://example.com",
			header:        map[string]string{},
			bodyReq:       nil,
			expectedError: false,
		},
		{
			name:          "*bytes.Buffer bodyReq",
			method:        "POST",
			url:           "http://example.com",
			header:        map[string]string{},
			bodyReq:       bytes.NewBufferString("test data"),
			expectedError: false,
		},
		{
			name:          "Other types of bodyReq",
			method:        "POST",
			url:           "http://example.com",
			header:        map[string]string{},
			bodyReq:       map[string]interface{}{"key": "value"},
			expectedError: false,
		},
		// Add more test cases here
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := httputil.NewHttpRequest(context.Background(), tc.method, tc.url, tc.header, tc.bodyReq)
			if tc.expectedError && err == nil {
				t.Error("Expected an error, but got nil")
			}
			if !tc.expectedError && err != nil {
				t.Errorf("Expected no error, but got: %v", err)
			}
		})
	}
}
