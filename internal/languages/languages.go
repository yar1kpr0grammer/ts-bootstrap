package languages

import (
	"embed"
	"encoding/json"
	"fmt"
)

//go:embed *.json
var files embed.FS

var languageData map[string]string

func AutoInit() error {
	return Init(detectSystemLanguage())
}

func Init(lang string) error {
	filename := lang + ".json"
	data, err := files.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("language file not found: %s", filename)
	}
	if err := json.Unmarshal(data, &languageData); err != nil {
		return fmt.Errorf("invalid json in %s: %w", filename, err)
	}
	return nil
}

func Get(key string) string {
	v, ok := languageData[key]
	if !ok {
		return "[missing key: " + key + "]"
	}
	return v
}
