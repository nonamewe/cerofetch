package hardware

import (
	"fmt"
	"os/exec"
	"strings"
)

func RAM() string {
	command := exec.Command("free", "-h")
	out, err := command.Output()
	if err != nil {
		return "Unknown"
	}
	for _, out := range strings.Split(string(out), "\n") {
		if strings.HasPrefix(out, "Mem: ") {
			out = strings.Trim(strings.TrimPrefix(out, "Mem: "), ``)
			out = strings.TrimLeft(out, " \t")
			parts := strings.Fields(out)

			if len(parts) < 3 {
				return "Unknown"
			}
			return fmt.Sprintf("%s / %s", parts[1], parts[0])

		}
	}
	return "Unknown"
}
