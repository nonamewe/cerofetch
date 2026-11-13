package hardware

import (
	"fmt"
	"os/exec"
	"strings"
)

func Disk() string {
	command := exec.Command("df", "-hP", "/")
	out, err := command.Output()
	if err != nil {
		return "Unknown"
	}
	for _, out := range strings.Split(string(out), "\n") {
		out = strings.TrimLeft(out, " \t")
		parts := strings.Fields(out)

		if len(parts) < 6 {
			return "Unknown"
		}
		if parts[5] == "/" {
			return fmt.Sprintf("%s / %s", parts[2], parts[1])
		}

	}
	return "Unknown"
}
