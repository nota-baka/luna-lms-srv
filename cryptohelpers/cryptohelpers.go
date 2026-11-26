package cryptohelpers

import (
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/chacha20poly1305"
)

var aead cipher.AEAD

// TODO: make configurable
const bcryptCost int = 12

func Setup(encKey string) error {
	if aead != nil {
		return errors.New("cryptohelpers is already setup")
	}
	encKeyBytes, err := base64.StdEncoding.DecodeString(encKey)
	if err != nil {
		return fmt.Errorf("failed to convert encryption key string to bytes: %w", err)
	}
	aead, err = chacha20poly1305.NewX(encKeyBytes)
	if err != nil {
		return fmt.Errorf("failed to create AEAD from key: %w", err)
	}
	return nil
}

func VerifyPassword(password string, hash []byte) bool {
	return bcrypt.CompareHashAndPassword(hash, []byte(password)) == nil
}

func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcryptCost)
}

func GenerateToken() ([]byte, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return nil, fmt.Errorf("failed to generate token: %w", err)
	}
	return b, nil
}

func HashToken(token []byte) []byte {
	sum := sha256.Sum256(token)
	return sum[:]
}

func VerifyToken(token []byte, hash []byte) bool {
	sum := HashToken(token)
	return subtle.ConstantTimeCompare(sum, hash) == 1
}

// TODO: add encryption and decryption
