package utils

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"tsBootstrup/internal/cmd"
)

func CreateReadMe(text string) {
	err := os.WriteFile("README.md", []byte(text), 0o644)
	cmd.Confirm(err, "create README")
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func hasNonASCII(s string) bool {
	for _, r := range s {
		if r > unicode.MaxASCII {
			return true
		}
	}
	return false
}

func CheckPathForNPM() error {
	path, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get working directory: %w", err)
	}

	parts := strings.Split(path, "/")
	currentDir := parts[len(parts)-1]
	if hasNonASCII(currentDir) {
		return fmt.Errorf("invalid path part: %q", currentDir)
	}

	return nil
}

func ThisFile() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Base(exePath), nil
}
