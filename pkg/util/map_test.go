package util

import (
	"reflect"
	"testing"

	"github.com/paper-indonesia/pg-mcp-server/constant"
	"github.com/stretchr/testify/require"
)

func TestGetValueFromMap(t *testing.T) {
	testCases := []struct {
		input    map[string]any
		key      string
		expected string
	}{
		{
			input:    map[string]any{"key": "value"},
			key:      "key",
			expected: "value",
		},
		{
			input:    map[string]any{"key": "value"},
			key:      "not-found-key",
			expected: "default",
		},
		{
			input:    nil,
			key:      "not-found-key",
			expected: "default",
		},
	}

	for _, tc := range testCases {
		actual := GetValueFromMap(tc.input, tc.key, "default")
		if actual != tc.expected {
			t.Errorf("GetValueFromMap(%v, %s) = %s, want %s", tc.input, tc.key, actual, tc.expected)
		}
	}
}
func TestConvert(t *testing.T) {
	testCases := []struct {
		name     string
		input    interface{}
		expected interface{}
	}{
		{
			name:     "Convert string to string",
			input:    "test",
			expected: "test",
		},
		{
			name:     "Convert int to int",
			input:    42,
			expected: 42,
		},
		{
			name:     "Convert float64 to float64",
			input:    3.14,
			expected: 3.14,
		},
		{
			name:     "Convert bool to bool",
			input:    true,
			expected: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := Convert[interface{}](tc.input)
			if result != tc.expected {
				t.Errorf("Convert(%v) = %v, want %v", tc.input, result, tc.expected)
			}
		})
	}

	// Test with specific types
	t.Run("Convert to string", func(t *testing.T) {
		result := Convert[string]("hello")
		if result != "hello" {
			t.Errorf("Convert[string](%q) = %q, want %q", "hello", result, "hello")
		}
	})

	t.Run("Convert to int", func(t *testing.T) {
		result := Convert[int](42)
		if result != 42 {
			t.Errorf("Convert[int](%d) = %d, want %d", 42, result, 42)
		}
	})

	t.Run("Convert incompatible type to int", func(t *testing.T) {
		result := Convert[int]("not an int")
		if result != 0 {
			t.Errorf("Convert[int](%q) = %d, want %d", "not an int", result, 0)
		}
	})
}
func TestExtractValueFromMap(t *testing.T) {
	testCases := []struct {
		name         string
		input        constant.TMapAny
		key          string
		defaultValue interface{}
		expected     interface{}
	}{
		{
			name:         "Extract string value",
			input:        constant.TMapAny{"key": "value"},
			key:          "key",
			defaultValue: "",
			expected:     "value",
		},
		{
			name:         "Extract int value",
			input:        constant.TMapAny{"key": 42},
			key:          "key",
			defaultValue: 0,
			expected:     42,
		},
		{
			name:         "Extract float value",
			input:        constant.TMapAny{"key": 3.14},
			key:          "key",
			defaultValue: 0.0,
			expected:     3.14,
		},
		{
			name:         "Extract bool value",
			input:        constant.TMapAny{"key": true},
			key:          "key",
			defaultValue: false,
			expected:     true,
		},
		{
			name:         "Key not found, return default value",
			input:        constant.TMapAny{"other": "value"},
			key:          "key",
			defaultValue: "default",
			expected:     "default",
		},
		{
			name:         "Empty map, return default value",
			input:        constant.TMapAny{},
			key:          "key",
			defaultValue: 100,
			expected:     100,
		},
		{
			name:         "Nil map, return default value",
			input:        nil,
			key:          "key",
			defaultValue: []string{"default"},
			expected:     []string{"default"},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := ExtractValueFromMap(tc.input, tc.key, tc.defaultValue)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("ExtractValueFromMap(%v, %q, %v) = %v, want %v", tc.input, tc.key, tc.defaultValue, result, tc.expected)
			}
		})
	}
}

func TestExtractValueFromMapWithDifferentTypes(t *testing.T) {
	m := constant.TMapAny{
		"string": "hello",
		"int":    42,
		"float":  3.14,
		"bool":   true,
		"slice":  []int{1, 2, 3},
	}

	t.Run("Extract string", func(t *testing.T) {
		result := ExtractValueFromMap[string](m, "string", "default")
		if result != "hello" {
			t.Errorf("ExtractValueFromMap[string]() = %v, want %v", result, "hello")
		}
	})
}
func TestCreateMapFromJsonStr(t *testing.T) {
	tests := []struct {
		name      string
		jsonStr   string
		wantError bool
	}{
		{
			name:      "Valid JSON object",
			jsonStr:   `{"name": "John", "age": 30, "city": "New York"}`,
			wantError: false,
		},
		{
			name:      "Empty JSON object",
			jsonStr:   `{}`,
			wantError: false,
		},
		{
			name:      "JSON with nested objects",
			jsonStr:   `{"user": {"name": "John", "age": 30}, "active": true}`,
			wantError: false,
		},
		{
			name:      "Invalid JSON",
			jsonStr:   `{"name": "John", "age": 30,}`,
			wantError: true,
		},
		{
			name:      "Malformed JSON",
			jsonStr:   `not a json string`,
			wantError: true,
		},
		{
			name:      "JSON with array",
			jsonStr:   `{"numbers": [1,2,3], "letters": ["a","b","c"]}`,
			wantError: false,
		},
		{
			name:      "JSON with null values",
			jsonStr:   `{"name": null, "age": 30}`,
			wantError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := CreateMapFromJsonStr(tt.jsonStr)
			if tt.wantError {
				require.NotEmpty(t, result["error"])
			}
		})
	}
}

func TestChangeNestedMapValue(t *testing.T) {
	testCases := []struct {
		desc     string
		wantErr  bool
		inputMap constant.TMapAny
		expected constant.TMapAny
		path     []string
		newValue any
	}{
		{
			desc:    "success change value of inteface",
			wantErr: false,
			inputMap: constant.TMapAny{
				"object": constant.TMapAny{
					"name": "food",
				},
			},
			path: []string{"object"},
			newValue: constant.TMapAny{
				"name": "drink",
			},
			expected: constant.TMapAny{
				"object": constant.TMapAny{
					"name": "drink",
				},
			},
		},
		{
			desc:    "success change value of nested interface",
			wantErr: false,
			inputMap: constant.TMapAny{
				"object": constant.TMapAny{
					"class": constant.TMapAny{
						"name": "food",
					},
				},
			},
			path: []string{"object"},
			newValue: constant.TMapAny{
				"class": constant.TMapAny{
					"name": "drink",
				},
			},
			expected: constant.TMapAny{
				"object": constant.TMapAny{
					"class": constant.TMapAny{
						"name": "drink",
					},
				},
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := ChangeNestedMapValue(tc.inputMap, tc.path, tc.newValue)
			if tc.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expected, tc.inputMap)
			}
		})
	}
}
