package utils

import (
	"os"
	"strings"
)

func isRunningInContainer() bool {
	cgroupFile := "/proc/1/cgroup"

	data, err := os.ReadFile(cgroupFile)
	if err != nil {
		return false
	}

	content := string(data)

	return strings.Contains(content, "/docker/") || strings.Contains(content, "/containerd/")
}

func MeiliSearchHost() string {
	if isRunningInContainer() {
		return "http://meilisearch:7700"
	}

	return "http://127.0.0.1:7700"
}
