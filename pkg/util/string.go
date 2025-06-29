package util

import (
	"encoding/json"
	"fmt"
	"html"
	"regexp"
	"strings"

	"github.com/paper-indonesia/pg-mcp-server/constant"
)

var numericRegex = regexp.MustCompile(constant.NUMERIC_REGEX)

func escapestr(sql string) string {
	r := strings.NewReplacer(string(byte(0)), "0", "\\", "\\\\", "\r", "\\r", "\n", "\\n", "\\0", "\\\\0", "\x1a", "\\Z", `"`, `\"`, `'`, `''`)
	return r.Replace(sql)
}

func unescapestr(sql string) string {
	r := strings.NewReplacer("\\\\", "\\", "\\r", "\r", "\\n", "\n", "\\\\0", "\\0", "\\Z", "\x1a", `\"`, `"`, `'`, `'`)
	return r.Replace(sql)
}

func SanitizeStr(str string) string {
	str = html.EscapeString(str)
	return escapestr(str)
}

func ToIDR(f float64) string {
	// Convert float to string with no decimal part
	s := fmt.Sprintf("%.0f", f)
	// Create a builder to construct the formatted string
	var result strings.Builder
	result.WriteString("IDR ")
	// Add characters to the result string with commas
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result.WriteString(",")
		}
		result.WriteRune(c)
	}
	return result.String()
}

func ToIDRWithSign(f float64) string {
	// Determine the sign
	sign := "+"
	if f < 0 {
		sign = "-"
		f = -f // Convert to positive for formatting
	}
	// Format the number as IDR
	formatted := ToIDR(f)
	// Add the sign back
	return sign + formatted
}

func ToJsonString(data any) string {
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(jsonBytes)
}

func IsPatternMatch(pattern, str string) bool {
	// Compile the regular expression
	re, err := regexp.Compile(pattern)
	if err != nil {
		return false
	}

	// Compare the string with the pattern
	if re.MatchString(str) {
		return true
	}

	return false
}

func TrimLength(str string, maxLength int) string {
	// trim, max length maxLength
	strLen := len(str) - maxLength
	if strLen <= 0 {
		strLen = 0
	}
	return str[strLen:]
}

func InArray[T comparable](str T, arr []T) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func TrimWhitespace(str string) string {
	trimedString := strings.ReplaceAll(strings.TrimSpace(str), " ", "")
	trimedString = strings.ReplaceAll(trimedString, "\t", "")
	trimedString = strings.ReplaceAll(trimedString, "\n", "")
	return trimedString
}

func RemoveDuplicates(arr []string) []string {
	seen := make(map[string]bool)
	var result []string
	for _, value := range arr {
		if !seen[value] {
			seen[value] = true
			result = append(result, value)
		}
	}
	return result
}

func GetLastString(s string, length int) string {
	if len(s) <= length {
		return s
	}
	return s[len(s)-length:]
}

func IsNumericString(s string) bool {
	if s == "" {
		return true
	}

	return numericRegex.MatchString(s)
}
