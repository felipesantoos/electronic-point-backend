package validator

import "strings"

func TextIsBlank(text string) bool {
	return len(strings.TrimSpace(text)) == 0
}
