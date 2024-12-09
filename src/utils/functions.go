package utils

import (
	"backend_template/src/utils/validator"
	"regexp"
	"strings"
)

func GetNullableValue[T any](i interface{}) *T {
	switch v := i.(type) {
	case T:
		return &v
	default:
		return nil
	}
}

func ExtractExtensionFromFile(filename string) string {
	filename = strings.ToLower(filename)
	if validator.IsTextBlank(filename) {
		return ""
	}
	re := regexp.MustCompile(`\.([^\.]+)$`)
	match := re.FindStringSubmatch(filename)

	if len(match) < 2 {
		return ""
	}
	return match[1]
}

func RemoveExtensionFromFileName(filename string) string {
	lastDotIndex := strings.LastIndex(filename, ".")
	if lastDotIndex == -1 {
		return filename
	}
	return filename[:lastDotIndex]
}
