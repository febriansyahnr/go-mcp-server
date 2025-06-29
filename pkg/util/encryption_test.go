package util

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAES256GCM_Encrypt(t *testing.T) {
	tests := []struct {
		name        string
		secret      string
		nonce       string
		input       string
		wantErr     bool
		errContains string
	}{
		{
			name:    "successful encryption",
			secret:  "0123456789abcdef0123456789abcdef",
			nonce:   "123456789012",
			input:   "test data",
			wantErr: false,
		},
		{
			name:        "invalid key size",
			secret:      "short-key",
			nonce:       "123456789012",
			input:       "test data",
			wantErr:     true,
			errContains: "invalid key size",
		},
		{
			name:    "empty string encryption",
			secret:  "0123456789abcdef0123456789abcdef",
			nonce:   "123456789012",
			input:   "",
			wantErr: false,
		},
		{
			name:    "long string encryption",
			secret:  "0123456789abcdef0123456789abcdef",
			nonce:   "123456789012",
			input:   strings.Repeat("a", 1000),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aes := NewAES256GCM(tt.secret, tt.nonce)
			result, err := aes.Encrypt(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
				return
			}

			assert.NoError(t, err)
			assert.NotEmpty(t, result)
			assert.Len(t, result, len(tt.input)*2+32) // hex encoding doubles length, plus GCM tag
			assert.Regexp(t, "^[0-9a-f]+$", result)   // hex encoded result
		})
	}
}
func TestAES256GCM_Decrypt(t *testing.T) {
	tests := []struct {
		name        string
		secret      string
		nonce       string
		input       string
		wantOutput  string
		wantErr     bool
		errContains string
	}{
		{
			name:        "authentication failed",
			secret:      "0123456789abcdef0123456789abcdef",
			nonce:       "123456789012",
			input:       "deadbeefdeadbeefdeadbeefdeadbeef",
			wantErr:     true,
			errContains: "authentication failed",
		},
		{
			name:        "empty ciphertext",
			secret:      "0123456789abcdef0123456789abcdef",
			nonce:       "123456789012",
			input:       "",
			wantErr:     true,
			errContains: "authentication failed",
		},
		{
			name:        "modified ciphertext",
			secret:      "0123456789abcdef0123456789abcdef",
			nonce:       "123456789012",
			input:       strings.Repeat("a", 64),
			wantErr:     true,
			errContains: "authentication failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aes := NewAES256GCM(tt.secret, tt.nonce)
			_, err := aes.Decrypt(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
				if tt.errContains != "" {
					assert.Contains(t, err.Error(), tt.errContains)
				}
				return
			}

			assert.NoError(t, err)
		})
	}
}
