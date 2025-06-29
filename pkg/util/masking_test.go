package util

import "testing"

func TestMaskingCreditCardNumber(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{
			input:    "123456789",
			expected: "123456789",
		},
		{
			input:    "1234567890123456",
			expected: "123456******3456",
		},
	}
	for _, tc := range testCases {
		actual := MaskCreditCardNumber(tc.input)
		if actual != tc.expected {
			t.Errorf("MaskCreditCardNumber(%s) = %s, want %s", tc.input, actual, tc.expected)
		}
	}
}
