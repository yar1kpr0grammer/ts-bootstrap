package ts

import (
	"os"
	"path/filepath"
	"tsBootstrup/src/cmd"
)

func Init(outputMode cmd.Output) error {
	os.Remove("tsconfig.json")
	err := cmd.Run(outputMode, "npx", "tsc", "--init")
	cmd.Confirm(err, "npx tsc --init")
	return err
}

func UpdateConfig(path string, text string) error {
	err := os.WriteFile(path, []byte(text), 0644)
	cmd.Confirm(err, "update tsconfig.json")
	return err
}

func CreateIndexFile(text string) error {
	path := filepath.Join("src", "index.ts")
	os.MkdirAll("src", os.ModePerm)

	err := os.WriteFile(path, []byte(text), 0644)
	cmd.Confirm(err, "create src/index.ts")
	return err
}
