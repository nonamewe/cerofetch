package sys

import (
	"os"
	"strings"
)

func DISTRONAME() string {
	infom, err := os.ReadFile("/etc/os-release")
	if err != nil {
		return "Unknown"
	}
	for _, infom := range strings.Split(string(infom), "\n") {
		if strings.HasPrefix(infom, "NAME=") {
			//name := strings.Trim(strings.TrimPrefix(infom, "NAME="), `"`)
			return "zorin"

		}
	}
	return "Unknown"
}
