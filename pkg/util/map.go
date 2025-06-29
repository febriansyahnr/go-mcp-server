package util

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/paper-indonesia/pg-mcp-server/constant"
)

// GetValueFromMap get value from map, if not found, return default value
func GetValueFromMap(m map[string]any, key string, defaultValue string) string {
	if m == nil {
		return defaultValue
	}
	if val, ok := m[key]; ok {
		return fmt.Sprintf("%v", val)
	}
	return defaultValue
}

// ExtractValueFromMap extract value from map, if not found, return default value
func ExtractValueFromMap[T any](obj constant.TMapAny, key string, defaultValue T) T {
	if obj == nil {
		return defaultValue
	}

	if val, ok := obj[key]; ok {
		return Convert[T](val)
	}

	return defaultValue
}

// Convert convert interface to type T
func Convert[T any](data interface{}) T {
	if res, ok := data.(T); ok {
		return res
	}
	var res T
	return res
}

// CreateMapFromJsonStr creates a map from a given json string. If the json string
func CreateMapFromJsonStr(jsonStr string) map[string]any {
	result := make(map[string]any)
	if err := json.Unmarshal([]byte(jsonStr), &result); err == nil {
		return result
	} else {
		result["error"] = err.Error()
	}

	return result
}

// ChangeNestedMapValue updates the value in a nested map at the specified path.
// to the existing value type at the final path key.
func ChangeNestedMapValue(m constant.TMapAny, path []string, newValue interface{}) error {
	current := reflect.ValueOf(m)
	currentMap := m // Keep a copy of the map to modify

	for i, key := range path {
		if current.Kind() != reflect.Map {
			return fmt.Errorf("path element %d key: %s is not a map, but %v", i, key, current.Kind())
		}

		value := current.MapIndex(reflect.ValueOf(key))

		if !value.IsValid() {
			return fmt.Errorf("path element %q not found", key)
		}

		if i == len(path)-1 {
			if reflect.TypeOf(newValue) == reflect.TypeOf(value.Interface()) {
				currentMap[key] = newValue // Directly assign the new value
				return nil
			} else {
				fmt.Println(reflect.TypeOf(newValue))
				fmt.Println(value.Interface())
				return fmt.Errorf("cannot assign %T to %T", newValue, value.Interface())
			}
		}

		current = value
		currentMap = currentMap[key].(constant.TMapAny) // move down the map
	}
	return nil
}

func MapToJsonString(m constant.TMapAny) string {
	b, err := json.Marshal(m)
	if err != nil {
		return ""
	}
	return string(b)
}
