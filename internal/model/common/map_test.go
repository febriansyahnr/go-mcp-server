package commonModel

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTMapString_Json(t *testing.T) {
	tests := []struct {
		name     string
		input    *TMapString
		expected string
	}{
		{
			name:     "nil map",
			input:    nil,
			expected: "{}",
		},
		{
			name:     "empty map",
			input:    &TMapString{},
			expected: "{}",
		},
		{
			name: "single key-value pair",
			input: &TMapString{
				"key1": "value1",
			},
			expected: `{"key1":"value1"}`,
		},
		{
			name: "multiple key-value pairs",
			input: &TMapString{
				"name":  "John",
				"email": "john@example.com",
				"city":  "Jakarta",
			},
			expected: `{"city":"Jakarta","email":"john@example.com","name":"John"}`,
		},
		{
			name: "empty string values",
			input: &TMapString{
				"empty": "",
				"key":   "value",
			},
			expected: `{"empty":"","key":"value"}`,
		},
		{
			name: "special characters in values",
			input: &TMapString{
				"special": "value with spaces & symbols!",
				"unicode": "测试",
			},
			expected: `{"special":"value with spaces \u0026 symbols!","unicode":"测试"}`,
		},
		{
			name: "special characters in keys",
			input: &TMapString{
				"key with spaces": "value1",
				"key-with-dash":   "value2",
				"key_with_under":  "value3",
			},
			expected: `{"key with spaces":"value1","key-with-dash":"value2","key_with_under":"value3"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.Json()

			// For non-deterministic JSON ordering (maps), parse and compare
			if tt.name == "multiple key-value pairs" || tt.name == "special characters in keys" {
				var expectedMap, resultMap map[string]string
				err1 := json.Unmarshal([]byte(tt.expected), &expectedMap)
				err2 := json.Unmarshal(result, &resultMap)

				assert.NoError(t, err1)
				assert.NoError(t, err2)
				assert.Equal(t, expectedMap, resultMap)
			} else {
				assert.Equal(t, tt.expected, string(result))
			}

			// Ensure result is valid JSON
			var jsonCheck map[string]string
			err := json.Unmarshal(result, &jsonCheck)
			assert.NoError(t, err, "Result should be valid JSON")
		})
	}
}

func TestTMapString_ToMapAny(t *testing.T) {
	tests := []struct {
		name     string
		input    *TMapString
		expected TMapAny
	}{
		{
			name:     "nil map",
			input:    nil,
			expected: nil,
		},
		{
			name:     "empty map",
			input:    &TMapString{},
			expected: TMapAny{},
		},
		{
			name: "single key-value pair",
			input: &TMapString{
				"key1": "value1",
			},
			expected: TMapAny{
				"key1": "value1",
			},
		},
		{
			name: "multiple key-value pairs",
			input: &TMapString{
				"name":  "John",
				"email": "john@example.com",
				"city":  "Jakarta",
			},
			expected: TMapAny{
				"name":  "John",
				"email": "john@example.com",
				"city":  "Jakarta",
			},
		},
		{
			name: "empty string values",
			input: &TMapString{
				"empty": "",
				"key":   "value",
			},
			expected: TMapAny{
				"empty": "",
				"key":   "value",
			},
		},
		{
			name: "special characters",
			input: &TMapString{
				"special": "value with spaces & symbols!",
				"unicode": "测试",
			},
			expected: TMapAny{
				"special": "value with spaces & symbols!",
				"unicode": "测试",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.input.ToMapAny()
			assert.Equal(t, tt.expected, result)

			// Additional validation for non-nil cases
			if tt.input != nil && tt.expected != nil {
				// Ensure all original keys are present
				for key, value := range *tt.input {
					assert.Contains(t, result, key)
					assert.Equal(t, value, result[key])
				}

				// Ensure no extra keys are added
				assert.Equal(t, len(*tt.input), len(result))
			}
		})
	}
}

func TestTMapString_ToMapAny_TypeAssertion(t *testing.T) {
	input := &TMapString{
		"string_value": "test",
		"number_like":  "123",
		"boolean_like": "true",
	}

	result := input.ToMapAny()

	// Verify all values are still strings (not converted to other types)
	for key, value := range result {
		assert.IsType(t, "", value, "Value for key %s should remain as string", key)
	}
}

func TestTMapString_ToMapAny_Independence(t *testing.T) {
	original := &TMapString{
		"key1": "value1",
		"key2": "value2",
	}

	result := original.ToMapAny()

	// Modify the result
	result["key1"] = "modified"
	result["new_key"] = "new_value"

	// Original should be unchanged
	assert.Equal(t, "value1", (*original)["key1"])
	assert.NotContains(t, *original, "new_key")
}

func TestTMapString_Json_ErrorHandling(t *testing.T) {
	// Create a map that would normally marshal successfully
	input := &TMapString{
		"normal": "value",
	}

	result := input.Json()

	// Should return valid JSON
	var jsonCheck map[string]string
	err := json.Unmarshal(result, &jsonCheck)
	assert.NoError(t, err)
	assert.Equal(t, "value", jsonCheck["normal"])
}
