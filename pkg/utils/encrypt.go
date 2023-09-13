package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
)

const (
	AesEncryptKey = "tcr-access-control-secret-012345" // 32 bytes length
)

func EncryptAES(key string, text string) (string, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("NewCipher: %w", err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", fmt.Errorf("NewGCM: %w", err)
	}
	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("ReadFull: %w", err)
	}

	e := hex.EncodeToString(gcm.Seal(nonce, nonce, []byte(text), nil))
	return e, nil
}

func DecryptAES(key string, e string) (string, error) {
	ciphertext, err := hex.DecodeString(e)
	if err != nil {
		return "", fmt.Errorf("DecodeString: %w", err)
	}
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", fmt.Errorf("NewCipher: %w", err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return "", fmt.Errorf("NewCipher: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("text length should >= %v", nonceSize)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	decrypted, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}

	return string(decrypted), nil
}
