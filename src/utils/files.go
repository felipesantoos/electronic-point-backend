package utils

import (
	"path/filepath"
)

func ExtractFileExtension(filename string) string {
	return filepath.Ext(filename)
}
