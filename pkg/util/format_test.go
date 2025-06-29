package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetOrDefault(t *testing.T) {
	testCases := []struct {
		desc     string
		value    *int
		defValue int
		expected int
	}{
		{
			desc:     "should return value when not nil",
			value:    intPtr(42),
			defValue: 10,
			expected: 42,
		},
		{
			desc:     "should return default when nil",
			value:    nil,
			defValue: 10,
			expected: 10,
		},
		{
			desc:     "should return zero value when provided",
			value:    intPtr(0),
			defValue: 10,
			expected: 0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := GetOrDefault(tc.value, tc.defValue)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestGetStringOrDefault(t *testing.T) {
	testCases := []struct {
		desc     string
		value    string
		defValue string
		expected string
	}{
		{
			desc:     "should return value when not empty",
			value:    "hello",
			defValue: "default",
			expected: "hello",
		},
		{
			desc:     "should return default when empty",
			value:    "",
			defValue: "default",
			expected: "default",
		},
		{
			desc:     "should return space when provided",
			value:    " ",
			defValue: "default",
			expected: " ",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := GetStringOrDefault(tc.value, tc.defValue)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestGetSliceOrDefault(t *testing.T) {
	testCases := []struct {
		desc     string
		value    *[]string
		defValue []string
		expected []string
	}{
		{
			desc:     "should return value when not nil",
			value:    &[]string{"a", "b"},
			defValue: []string{"default"},
			expected: []string{"a", "b"},
		},
		{
			desc:     "should return default when nil",
			value:    nil,
			defValue: []string{"default"},
			expected: []string{"default"},
		},
		{
			desc:     "should return default when slice is nil",
			value:    func() *[]string { var s []string; return &s }(),
			defValue: []string{"default"},
			expected: []string{"default"},
		},
		{
			desc:     "should return empty slice when provided",
			value:    &[]string{},
			defValue: []string{"default"},
			expected: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := GetSliceOrDefault(tc.value, tc.defValue)
			assert.Equal(t, tc.expected, *result)
		})
	}
}

func TestRightPadWithSpace(t *testing.T) {
	testCases := []struct {
		desc     string
		input    string
		length   int
		expected string
	}{
		{
			desc:     "should pad short string with spaces",
			input:    "123",
			length:   5,
			expected: "123  ",
		},
		{
			desc:     "should return string as is when equal length",
			input:    "12345",
			length:   5,
			expected: "12345",
		},
		{
			desc:     "should truncate when string is longer",
			input:    "123456789",
			length:   5,
			expected: "12345",
		},
		{
			desc:     "should handle empty string",
			input:    "",
			length:   3,
			expected: "   ",
		},
		{
			desc:     "should handle zero length",
			input:    "test",
			length:   0,
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := RightPadWithSpace(tc.input, tc.length)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestGenerateNextPANDigit(t *testing.T) {
	testCases := []struct {
		desc     string
		pan      string
		expected string
	}{
		{
			desc:     "should generate valid PAN with Luhn digit",
			pan:      "42424242424242424",
			expected: "424242424242424242",
		},
		{
			desc:     "should handle empty PAN",
			pan:      "",
			expected: "",
		},
		{
			desc:     "should handle short PAN",
			pan:      "123",
			expected: "1230",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := GenerateNextPANDigit(tc.pan)
			assert.Equal(t, tc.expected, result)
		})
	}
}

func TestGenerateNextLuhnDigit(t *testing.T) {
	testCases := []struct {
		desc        string
		prefix      string
		expectedNum int
		expectedOk  bool
	}{
		{
			desc:        "should generate correct Luhn digit",
			prefix:      "42424242424242424",
			expectedNum: 2,
			expectedOk:  true,
		},
		{
			desc:        "should generate correct digit for even checksum",
			prefix:      "12345",
			expectedNum: 5,
			expectedOk:  true,
		},
		{
			desc:        "should handle short prefix",
			prefix:      "123",
			expectedNum: 0,
			expectedOk:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			num, ok := GenerateNextLuhnDigit(tc.prefix)
			assert.Equal(t, tc.expectedNum, num)
			assert.Equal(t, tc.expectedOk, ok)
		})
	}
}

func TestCalculateLuhnChecksum(t *testing.T) {
	testCases := []struct {
		desc           string
		cardNumber     string
		expectedSum    int
		expectedValid  bool
	}{
		{
			desc:           "should calculate checksum for valid card number",
			cardNumber:     "4242424242424242",
			expectedSum:    0,
			expectedValid:  true,
		},
		{
			desc:           "should calculate checksum for invalid card number",
			cardNumber:     "4242424242424243",
			expectedSum:    1,
			expectedValid:  true,
		},
		{
			desc:           "should return false for non-numeric input",
			cardNumber:     "424242424242424a",
			expectedSum:    0,
			expectedValid:  false,
		},
		{
			desc:           "should handle single digit",
			cardNumber:     "4",
			expectedSum:    4,
			expectedValid:  true,
		},
		{
			desc:           "should handle empty string",
			cardNumber:     "",
			expectedSum:    0,
			expectedValid:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			sum, valid := CalculateLuhnChecksum(tc.cardNumber)
			assert.Equal(t, tc.expectedSum, sum)
			assert.Equal(t, tc.expectedValid, valid)
		})
	}
}

func TestIsLuhnValid(t *testing.T) {
	testCases := []struct {
		desc       string
		cardNumber string
		expected   bool
	}{
		{
			desc:       "should return true for valid card number",
			cardNumber: "4242424242424242",
			expected:   true,
		},
		{
			desc:       "should return false for invalid card number",
			cardNumber: "4242424242424243",
			expected:   false,
		},
		{
			desc:       "should return false for non-numeric input",
			cardNumber: "424242424242424a",
			expected:   false,
		},
		{
			desc:       "should return false for single invalid digit",
			cardNumber: "4",
			expected:   false,
		},
		{
			desc:       "should return true for empty string (checksum 0)",
			cardNumber: "",
			expected:   true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			result := IsLuhnValid(tc.cardNumber)
			assert.Equal(t, tc.expected, result)
		})
	}
}

// Helper function to create int pointer
func intPtr(i int) *int {
	return &i
}