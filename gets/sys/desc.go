package sys

import (
	"os"
	"strings"
)

func DESCTOP() string {
	outm := os.Getenv("XDG_CURRENT_DESKTOP")
	out := strings.TrimPrefix(outm, "/usr/bin/")
	return out

}
