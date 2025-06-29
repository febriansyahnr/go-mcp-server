package validatorExt

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestVAStruct struct {
	VirtualAccountNo string `validate:"required,va_number"`
}

func TestValidateVANumber(t *testing.T) {
	validator := New()

	testCases := []struct {
		name        string
		vaNumber    string
		shouldPass  bool
		description string
	}{
		{
			name:        "Valid - only digits",
			vaNumber:    "1234567890",
			shouldPass:  true,
			description: "Should pass with only digits",
		},
		{
			name:        "Invalid - digits with spaces in middle",
			vaNumber:    "123 456 789",
			shouldPass:  false,
			description: "Should fail with spaces in middle",
		},
		{
			name:        "Valid - leading space",
			vaNumber:    " 1234567890",
			shouldPass:  true,
			description: "Should pass with leading space",
		},
		{
			name:        "Valid - trailing space",
			vaNumber:    "1234567890 ",
			shouldPass:  true,
			description: "Should pass with trailing space",
		},
		{
			name:        "Invalid - contains letters",
			vaNumber:    "123abc456",
			shouldPass:  false,
			description: "Should fail with letters",
		},
		{
			name:        "Invalid - contains special characters",
			vaNumber:    "123-456-789",
			shouldPass:  false,
			description: "Should fail with special characters",
		},
		{
			name:        "Invalid - empty string",
			vaNumber:    "",
			shouldPass:  false,
			description: "Should fail with empty string",
		},
		{
			name:        "Invalid - only spaces",
			vaNumber:    "   ",
			shouldPass:  false,
			description: "Should fail with only spaces",
		},
		{
			name:        "Valid - leading and trailing spaces",
			vaNumber:    "  123456  ",
			shouldPass:  true,
			description: "Should pass with leading and trailing spaces",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			testStruct := TestVAStruct{
				VirtualAccountNo: tc.vaNumber,
			}

			err := validator.Struct(testStruct)

			if tc.shouldPass {
				assert.NoError(t, err, tc.description)
			} else {
				assert.Error(t, err, tc.description)
			}
		})
	}
}