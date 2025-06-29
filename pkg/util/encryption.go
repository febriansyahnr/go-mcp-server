package util

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

// AES 256 GCM
type AES256GCM struct {
	secret string
	nonce  string
}

func NewAES256GCM(secret, nonce string) *AES256GCM {
	return &AES256GCM{
		secret: secret,
		nonce:  nonce,
	}
}

func (a *AES256GCM) Encrypt(data string) (string, error) {
	block, err := aes.NewCipher([]byte(a.secret))
	if err != nil {
		return "", err
	}
	nonce := []byte(a.nonce)
	dataBytes := []byte(data)
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	ciphertext := aesgcm.Seal(nil, nonce, dataBytes, nil)

	return fmt.Sprintf("%x", ciphertext), nil
}

func (a *AES256GCM) Decrypt(data string) (string, error) {
	block, err := aes.NewCipher([]byte(a.secret))
	if err != nil {
		return "", err
	}
	nonce := []byte(a.nonce)
	ciphertext, _ := hex.DecodeString(data)
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plaintext), nil
}
