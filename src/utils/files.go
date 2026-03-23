package utils

import (
	"os"
	"tsBootstrup/src/cmd"
)

func CreateReadMe(text string) {
	err := os.WriteFile("README.md", []byte(text), 0644)
	cmd.Confirm(err, "create README")
}

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
