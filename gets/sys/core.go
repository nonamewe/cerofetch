package sys

import (
	"os/exec"
	"strings"
)

func LINUX() string {
	linuxver := exec.Command("uname", "-r")
	linuxvero, err := linuxver.Output()
	if err != nil {
		return "Unknown"
	}

	core := strings.TrimSpace(string(linuxvero))
	return core
}
