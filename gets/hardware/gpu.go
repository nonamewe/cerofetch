package hardware

import (
	"os/exec"
	"strings"
)

func GPU() string {
	command := exec.Command("lspci")
	out, err := command.Output()
	if err != nil {
		return "Unknown"
	}
	for _, out := range strings.Split(string(out), "\n") {
		if strings.Contains(out, "VGA compatible controller:") || strings.Contains(out, "3D controller:") || strings.Contains(out, "Display controller:") {
			parts := strings.SplitN(out, ": ", 2)
			parts[1] = strings.Trim(strings.TrimPrefix(parts[1], "NVIDIA Corporation"), ``)
			parts[1] = strings.Trim(strings.TrimPrefix(parts[1], "Advanced Micro Devices"), ``)
			parts[1] = strings.Trim(strings.TrimPrefix(parts[1], "Intel Corporation"), ``)
			return parts[1]

		}
	}
	return "Unknown"
}
