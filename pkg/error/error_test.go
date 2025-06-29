package errors

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNew(t *testing.T) {
	testCases := []struct {
		desc    string
		errType string
		err     error
		want    string
	}{
		{
			desc:    "should create error with type prefix",
			errType: "TEST_ERROR",
			err:     fmt.Errorf("original error"),
			want:    "TEST_ERROR | original error",
		},
		{
			desc:    "should create error with empty type",
			errType: "",
			err:     fmt.Errorf("original error"),
			want:    " | original error",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := New(tc.errType, tc.err)
			assert.Equal(t, tc.want, result.Error())
		})
	}
}

func TestExtractError(t *testing.T) {
	testCases := []struct {
		desc             string
		input            error
		expectedStrError string
		expectedError    error
	}{
		{
			desc:             "should extract error type and message from formatted error",
			input:            New("test", fmt.Errorf("test error")),
			expectedStrError: "test",
			expectedError:    fmt.Errorf("test error"),
		},
		{
			desc:             "should handle plain error without type",
			input:            fmt.Errorf("test error"),
			expectedStrError: "",
			expectedError:    fmt.Errorf("test error"),
		},
		{
			desc:             "should extract from complex nested error",
			input:            New("OUTER_ERROR", New("INNER_ERROR", fmt.Errorf("inner message"))),
			expectedStrError: "INNER_ERROR",
			expectedError:    fmt.Errorf("inner message"),
		},
		{
			desc:             "should handle multiple pipe separators",
			input:            fmt.Errorf("TYPE1 | TYPE2 | final message"),
			expectedStrError: "TYPE2",
			expectedError:    fmt.Errorf("final message"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			strError, err := ExtractError(tc.input)
			require.Equal(t, tc.expectedStrError, strError)
			require.Equal(t, tc.expectedError, err)
		})
	}
}

func TestGenerateFeatureNotSupportedError(t *testing.T) {
	testCases := []struct {
		desc string
		err  error
		want string
	}{
		{
			desc: "should generate feature not supported error",
			err:  fmt.Errorf("some feature"),
			want: "FEATURE_NOT_SUPPORTED | some feature",
		},
		{
			desc: "should handle nil error",
			err:  fmt.Errorf("nil error case"),
			want: "FEATURE_NOT_SUPPORTED | nil error case",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := GenerateFeatureNotSupportedError(tc.err)
			assert.Equal(t, tc.want, result.Error())
		})
	}
}

func TestIsErr(t *testing.T) {
	testCases := []struct {
		desc    string
		errType string
		err     error
		want    bool
	}{
		{
			desc:    "should return true when error type matches",
			errType: "TEST_ERROR",
			err:     New("TEST_ERROR", fmt.Errorf("test message")),
			want:    true,
		},
		{
			desc:    "should return false when error type does not match",
			errType: "OTHER_ERROR",
			err:     New("TEST_ERROR", fmt.Errorf("test message")),
			want:    false,
		},
		{
			desc:    "should return false for plain error without type",
			errType: "TEST_ERROR",
			err:     fmt.Errorf("plain error"),
			want:    false,
		},
		{
			desc:    "should handle nested errors correctly",
			errType: "INNER_ERROR",
			err:     New("OUTER_ERROR", New("INNER_ERROR", fmt.Errorf("nested message"))),
			want:    true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := IsErr(tc.errType, tc.err)
			assert.Equal(t, tc.want, result)
		})
	}
}
