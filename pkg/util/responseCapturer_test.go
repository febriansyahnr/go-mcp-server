package util

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResponseCapturerWriteHeader(t *testing.T) {
	tests := []struct {
		name           string
		statusToWrite  int
		secondWrite    int
		expectedStatus int
	}{
		{
			name:           "WriteHeader sets status code",
			statusToWrite:  http.StatusTeapot,
			secondWrite:    http.StatusOK,
			expectedStatus: http.StatusTeapot,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			capturer := NewResponseCapturer(rr)

			capturer.WriteHeader(tt.statusToWrite)
			require.True(t, capturer.WroteHeader, "WroteHeader should be true after WriteHeader")
			require.Equal(t, tt.expectedStatus, capturer.StatusCode, "StatusCode should match after WriteHeader")

			// WriteHeader should not overwrite if called again
			capturer.WriteHeader(tt.secondWrite)
			require.Equal(t, tt.expectedStatus, capturer.StatusCode, "StatusCode should not change after second WriteHeader")
		})
	}
}

func TestResponseCapturerWrite(t *testing.T) {
	tests := []struct {
		name           string
		firstWrite     []byte
		secondWrite    []byte
		expectedBody   []byte
		expectedStatus int
	}{
		{
			name:           "Write appends body and sets status",
			firstWrite:     []byte("hello"),
			secondWrite:    []byte(" world"),
			expectedBody:   []byte("hello world"),
			expectedStatus: http.StatusOK,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rr := httptest.NewRecorder()
			capturer := NewResponseCapturer(rr)

			n, err := capturer.Write(tt.firstWrite)
			require.NoError(t, err)
			require.Equal(t, len(tt.firstWrite), n)
			require.True(t, capturer.WroteHeader, "WroteHeader should be true after Write")
			require.Equal(t, tt.expectedStatus, capturer.StatusCode, "StatusCode should be set to 200 after Write")

			// Test multiple writes
			capturer.Write(tt.secondWrite)
			require.Equal(t, tt.expectedBody, capturer.Body, "Body should be appended correctly")
		})
	}
}
