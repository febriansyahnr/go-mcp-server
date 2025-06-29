package errors

import (
	"fmt"
	"strings"
)

func New(errType string, err error) error {
	err = fmt.Errorf("%s | %w", errType, err)
	return err
}

func ExtractError(err error) (string, error) {
	extErr := strings.Split(err.Error(), " | ")
	if len(extErr) >= 2 {
		return extErr[len(extErr)-2], fmt.Errorf(extErr[len(extErr)-1])
	}
	return "", err
}

func GenerateFeatureNotSupportedError(err error) error {
	return New(ErrCodeFeatureNotSupported, err)
}

func IsErr(errType string, err error) bool {
	codeErr, err := ExtractError(err)
	return codeErr == errType
}
