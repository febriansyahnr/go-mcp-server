package secret_reader

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoadSecretReaders(t *testing.T) {
	acquirers := []string{"permata", "aspi"}

	secretAcquirers := LoadSecrets(acquirers)
	permataSecret, ok := secretAcquirers["permata"]
	if !ok {
		t.Error("secret acquirer permata not found")
	}
	require.NotEmpty(t, permataSecret)
}
