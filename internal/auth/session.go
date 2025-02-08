package auth

import (
	"os"
	"path/filepath"
)

var sessionFile = "data/session.txt"

func SaveSession(username string) error {
	sessionDir := filepath.Dir(sessionFile)
	if err := os.MkdirAll(sessionDir, os.ModePerm); err != nil {
		return err
	}
	return os.WriteFile(sessionFile, []byte(username), 0600)
}

func LoadSession() (string, error) {
	data, err := os.ReadFile(sessionFile)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ClearSession() error {
	if _, err := os.Stat(sessionFile); os.IsNotExist(err) {
		return nil
	}
	return os.Remove(sessionFile)
}
