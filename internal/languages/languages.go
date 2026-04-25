package languages

import (
	"embed"
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"strings"
)

// Встраиваем все json-файлы из текущей папки в exe.
//
//go:embed *.json
var files embed.FS

var languageData map[string]string

func AutoInit() error {
	lang := detectSystemLanguage()
	return Init(lang)
}

func detectSystemLanguage() string {
	switch runtime.GOOS {

	case "windows":
		// Пока безопасный fallback.
		// Позже можно добавить WinAPI GetUserDefaultUILanguage.
		return "en"

	default:
		// Linux / macOS / BSD
		// Обычно LANG=ru_RU.UTF-8
		value := os.Getenv("LANG")

		if value == "" {
			value = os.Getenv("LC_ALL")
		}

		if value == "" {
			value = os.Getenv("LC_MESSAGES")
		}

		value = strings.ToLower(value)

		if strings.HasPrefix(value, "ru") {
			return "ru"
		}

		if strings.HasPrefix(value, "uk") {
			return "ru"
		}

		if strings.HasPrefix(value, "be") {
			return "ru"
		}

		return "en"
	}
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
	value, ok := languageData[key]
	if !ok {
		return "[missing key: " + key + "]"
	}

	return value
}
