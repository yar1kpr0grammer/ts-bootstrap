//go:build windows

package languages

import (
	"strings"
	"syscall"
	"unsafe"
)

func detectSystemLanguage() string {
	kernel32 := syscall.NewLazyDLL("kernel32.dll")
	proc := kernel32.NewProc("GetUserDefaultLocaleName")

	// Максимальная длина locale name — 85 символов (LOCALE_NAME_MAX_LENGTH)
	buf := make([]uint16, 85)
	proc.Call(
		uintptr(unsafe.Pointer(&buf[0])),
		uintptr(len(buf)),
	)

	// Получаем строку вида "ru-RU", "en-US", "uk-UA" и т.д.
	locale := syscall.UTF16ToString(buf)
	return localeToLang(locale)
}

// localeToLang преобразует BCP-47 тег (ru-RU) в код языка (ru).
func localeToLang(locale string) string {
	locale = strings.ToLower(locale)

	// Берём только основной тег до дефиса: "ru-RU" → "ru"
	lang := strings.SplitN(locale, "-", 2)[0]

	switch lang {
	case "ru", "uk", "be":
		return "ru" // Украинский и белорусский → русский файл
	case "en":
		return "en"
	default:
		return "en" // Фолбэк
	}
}
