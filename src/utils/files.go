package utils

import (
	"os"
	"tsBootstrup/src/cmd"
)

func CreateReadMe(text string) {
	err := os.WriteFile("README.md", []byte(text), 0644)
	cmd.Confirm(err, "create README")
}
