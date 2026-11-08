package hardware

import (
	"os/exec"
	"strings"
)

func CPU() string {
	cpumachine := exec.Command("lscpu")
	cpum, err := cpumachine.Output()
	if err != nil {
		return "Unknown"
	}
	for _, cpum := range strings.Split(string(cpum), "\n") {
		if strings.HasPrefix(cpum, "Model name:") {
			cpu := strings.Trim(strings.TrimPrefix(cpum, "Model name:"), `"`)
			cpu = strings.TrimLeft(cpu, " \t")
			return cpu
		}
	}
	return "Unknown"
}
