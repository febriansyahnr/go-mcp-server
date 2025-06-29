package util

import (
	"testing"
)

// TestToIDR tests the ToIDR function with various scenarios
func TestToIDR(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{
			name:     "normal large number",
			input:    5000000000.00,
			expected: "IDR 5,000,000,000",
		},
		{
			name:     "number with thousands separator",
			input:    1234567.00,
			expected: "IDR 1,234,567",
		},
		{
			name:     "very large number",
			input:    123456789012345.00,
			expected: "IDR 123,456,789,012,345",
		},
		{
			name:     "small number",
			input:    123.00,
			expected: "IDR 123",
		},
		{
			name:     "zero value",
			input:    0.00,
			expected: "IDR 0",
		},
		{
			name:     "negative number",
			input:    -1234567.00,
			expected: "IDR -1,234,567",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToIDR(tt.input)
			if result != tt.expected {
				t.Errorf("ToIDR(%f) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}

// TestToIDRWithSign tests the ToIDRWithSign function with various scenarios
func TestToIDRWithSign(t *testing.T) {
	tests := []struct {
		name     string
		input    float64
		expected string
	}{
		{
			name:     "positive large number",
			input:    5000000000.00,
			expected: "+IDR 5,000,000,000",
		},
		{
			name:     "negative large number",
			input:    -5000000000.00,
			expected: "-IDR 5,000,000,000",
		},
		{
			name:     "positive small number",
			input:    123.00,
			expected: "+IDR 123",
		},
		{
			name:     "negative small number",
			input:    -123.00,
			expected: "-IDR 123",
		},
		{
			name:     "zero value",
			input:    0.00,
			expected: "+IDR 0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToIDRWithSign(tt.input)
			if result != tt.expected {
				t.Errorf("ToIDRWithSign(%f) = %s; want %s", tt.input, result, tt.expected)
			}
		})
	}
}
func TestTrimWhitespace(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "leading and trailing spaces",
			input:    "  hello world  ",
			expected: "helloworld",
		},
		{
			name:     "multiple spaces between words",
			input:    "hello    world",
			expected: "helloworld",
		},
		{
			name:     "tabs and newlines",
			input:    "\thello\n\tworld\n",
			expected: "helloworld",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "only whitespace",
			input:    "   \t\n\r  ",
			expected: "",
		},
		{
			name:     "special characters with spaces",
			input:    " @#$%  ^&*() ",
			expected: "@#$%^&*()",
		},
		{
			name:     "numbers with spaces",
			input:    " 123 456  789 ",
			expected: "123456789",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := TrimWhitespace(tt.input)
			if result != tt.expected {
				t.Errorf("TrimWhitespace(%q) = %q; want %q", tt.input, result, tt.expected)
			}
		})
	}
}
