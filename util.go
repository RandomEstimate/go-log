package TradeLog

import (
	"os"
	"path/filepath"
)

func Exist(filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
		return true
	}
	return false
}

func joinFilePath(path, file string) string {
	return filepath.Join(path, file)
}

func shortFileName(file string) string {
	return filepath.Base(file)
}
