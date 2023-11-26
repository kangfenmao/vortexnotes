package utils

import (
	"os"
	"strings"
)

func isRunningInContainer() bool {
	osReleaseFile := "/etc/os-release"

	data, err := os.ReadFile(osReleaseFile)
	if err != nil {
		return false
	}

	return strings.Contains(string(data), "Alpine")
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
