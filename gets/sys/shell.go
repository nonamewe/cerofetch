package sys

import (
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func SHELL() string {
	ppid := strconv.Itoa(os.Getppid())
	command := exec.Command("ps", "-p", ppid, "-o", "comm=")
	out, err := command.Output()
	if err != nil {
		return "Unknown"
	}
	return strings.TrimSpace(string(out))
}
