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

func MeiliSearchHost() string {
	if isRunningInContainer() {
		return "http://meilisearch:7700"
	}

	return "http://127.0.0.1:7700"
}
