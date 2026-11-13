package sys

import (
	"os"
)

func Locale() string {
	outm := os.Getenv("LANG")
	return outm
}
