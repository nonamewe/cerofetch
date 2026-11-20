package sys

import (
	"fmt"
	"os/exec"
	"strings"
)

func Package() string {
	distro := DISTRONAME()
	switch {
	case strings.Contains(strings.ToLower(distro), "arch"):
		comm := exec.Command("pacman", "-Qq")
		out, err := comm.Output()
		if err != nil {
			return "Unknown"
		}
		count := strings.Count(string(out), "\n")
		return fmt.Sprintf("%d", count)

	case strings.Contains(strings.ToLower(distro), "void"):
		comm := exec.Command("xbps-query", "-l")
		out, err := comm.Output()
		if err != nil {
			return "Unknown"
		}
		count := strings.Count(string(out), "\n")
		return fmt.Sprintf("%d", count)

	case strings.Contains(strings.ToLower(distro), "debian"):
		comm := exec.Command("dpkg", "-l")
		out, err := comm.Output()
		if err != nil {
			return "Unknown"
		}
		count := strings.Count(string(out), "\n")
		return fmt.Sprintf("%d", count)
	}

	return "Unknown"
}
