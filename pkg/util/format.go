package util

import (
	"strconv"
	"strings"
)

func GetOrDefault[T any](value *T, defaultValue T) T {
	if value == nil {
		return defaultValue
	}
	return *value
}

func GetStringOrDefault(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}

func GetSliceOrDefault[T any](value *[]T, defaultValue []T) *[]T {
	if value == nil || *value == nil {
		return &defaultValue
	}
	return value
}

// RightPadWithSpace pads a string with spaces on the right to ensure it reaches the specified length.
// If the input string is already longer than or equal to the specified length, it returns the original string.
// Example: RightPadWithSpace("123", 5) returns "123  "
func RightPadWithSpace(s string, length int) string {
	if len(s) >= length {
		return s[:length]
	}

	return s + strings.Repeat(" ", length-len(s))
}

// GenerateNextPANDigit takes a PAN and generates the next digit based on the Luhn algorithm.
// It increments the last digit by 1 and ensures the result is valid according to the Luhn algorithm.
// Example: GenerateNextPANDigit("936008860370000001") might return "936008860370000010"
func GenerateNextPANDigit(pan string) string {
	if len(pan) == 0 {
		return pan
	}

	nextDigit, ok := GenerateNextLuhnDigit(pan)
	if !ok {
		return ""
	}

	return pan + strconv.Itoa(nextDigit)
}

func GenerateNextLuhnDigit(cardNumberPrefix string) (int, bool) {
	tempCardNumber := cardNumberPrefix + "0"

	checksum, valid := CalculateLuhnChecksum(tempCardNumber)
	if !valid {
		return 0, false
	}

	if checksum == 0 {
		return 0, true
	}
	return 10 - checksum, true
}

// CalculateLuhnChecksum calculates the Luhn checksum for a given string of digits.
// It returns the checksum (0-9) and true if the input is valid, or false if not.
func CalculateLuhnChecksum(cardNumber string) (int, bool) {
	sum := 0
	numDigits := len(cardNumber)
	parity := numDigits % 2 // 0-indexed, so if length is even, first digit is at odd position (0),
	// if length is odd, first digit is at even position (0)

	for i, r := range cardNumber {
		digit, err := strconv.Atoi(string(r))
		if err != nil {
			// Not a digit
			return 0, false
		}

		if i%2 == parity { // Double every second digit from the right
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum % 10, true
}

// IsLuhnValid checks if a given number string is valid according to the Luhn algorithm.
func IsLuhnValid(cardNumber string) bool {
	checksum, valid := CalculateLuhnChecksum(cardNumber)
	return valid && checksum == 0
}
