package validator

import "strings"

func IsTextBlank(text string) bool {
	text = strings.TrimSpace(text)
	return len(text) <= 0
}
