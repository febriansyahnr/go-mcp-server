package util

import "strings"

func Contains[T comparable](slice []T, item T) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}

func ContainsPrefix(slice []string, path string) bool {
	for _, item := range slice {
		if strings.Contains(path, item) {
			return true
		}
	}
	return false
}

func CheckPrefix(prefixes []string, str string) bool {
	for _, item := range prefixes {
		if strings.HasPrefix(str, strings.ToLower(item)) {
			return true
		}
	}
	return false
}
