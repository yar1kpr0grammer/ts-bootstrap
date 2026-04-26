//go:build !windows

package languages

import (
	"os"
	"strings"
)

func detectSystemLanguage() string {
	value := os.Getenv("LANG")
	if value == "" {
		value = os.Getenv("LC_ALL")
	}
	if value == "" {
		value = os.Getenv("LC_MESSAGES")
	}
	return localeToLang(strings.ToLower(value))
}

// localeToLang одинакова на всех платформах, но линкуется только нужная.
func localeToLang(locale string) string {
	// LANG может быть "ru_RU.UTF-8" или BCP-47 "ru-RU"
	// Нормализуем: заменяем '_' на '-' и берём первый сегмент
	locale = strings.NewReplacer("_", "-").Replace(locale)
	lang := strings.SplitN(locale, "-", 2)[0]

	switch lang {
	case "ru", "uk", "be":
		return "ru"
	default:
		return "en"
	}
}
