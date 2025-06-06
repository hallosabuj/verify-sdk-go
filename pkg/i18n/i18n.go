package i18n

import "fmt"

func TranslateWithCode(code string, defaultText string) string {
	return defaultText
}

func Translate(text string) string {
	return text
}

func TranslateWithArgs(text string, args ...any) string {
	return fmt.Sprintf(Translate(text), args...)
}
