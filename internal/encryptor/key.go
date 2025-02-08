package encryptor

import (
	"crypto/rand"
	"fmt"
	"os"
	"path/filepath"
)

func SaveKey(key []byte) error {
	secretDir := "data/secret"

	if err := os.MkdirAll(secretDir, os.ModePerm); err != nil {
		return err
	}

	keyPath := filepath.Join(secretDir, "key.txt")
	err := os.WriteFile(keyPath, key, 0600)
	if err == nil {
		fmt.Println("✅ Encryption key successfully saved at:", keyPath)
	}
	return err
}

func LoadKey() ([]byte, error) {
	keyPath := "data/secret/key.txt"
	return os.ReadFile(keyPath)
}

// GenerateKey creates a 32-byte AES key (AES-256)
func GenerateKey() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	return key, err
}
