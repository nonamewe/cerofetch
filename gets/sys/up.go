package sys

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Uptime() string {
	out, err := os.ReadFile("/proc/uptime")
	if err != nil {
		return "Unknown"
	}

	parts := strings.Fields(string(out))
	if len(parts) < 1 {
		return "Unknown"
	}
	sec, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return "Unknown"
	}
	hour := 0
	mins := sec / 60
	for {
		if mins-60 >= 0 {
			hour++
			mins -= 60
		} else {
			return fmt.Sprintf("%dh %dm", hour, int(mins))
		}
	}
}
