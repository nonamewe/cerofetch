package logocheck

import (
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
	"strings"

	sys "github.com/nonamewe/cerofetch/gets/sys"
)

var Distro string

func LoadImage(path string) image.Image {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	Img, _, err := image.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	return Img

}
func Image() image.Image {

	var Img image.Image
	switch {
	case strings.Contains(strings.ToLower(sys.DISTRONAME()), "arch"):
		Distro = "arch"
		Img = LoadImage("logo/arch.png")
	case strings.Contains(strings.ToLower(sys.DISTRONAME()), "gentoo"):
		Distro = "gentoo"
		Img = LoadImage("logo/gentoo.png")
	case strings.Contains(strings.ToLower(sys.DISTRONAME()), "debian"):
		Distro = "debian"
		Img = LoadImage("logo/debian.png")
	case strings.Contains(strings.ToLower(sys.DISTRONAME()), "mx"):
		Distro = "MX"
		Img = LoadImage("logo/MX.png")
	case strings.Contains(strings.ToLower(sys.DISTRONAME()), "mint"):
		Distro = "mint"
		Img = LoadImage("logo/mint.png")
	case strings.Contains(strings.ToLower(sys.DISTRONAME()), "cachyos"):
		Distro = "cachyos"
		Img = LoadImage("logo/cachyos.png")
	default:
		Distro = ""

	}
	return Img
}

func RenderImage(img image.Image, scale int) []string {
	var lines []string
	bounds := img.Bounds()

	for y := bounds.Min.Y; y < bounds.Max.Y; y += scale {
		var line strings.Builder

		for x := bounds.Min.X; x < bounds.Max.X; x += scale {
			r, g, b, a := img.At(x, y).RGBA()

			if a == 0 {
				line.WriteString("  ")
				continue
			}

			fmt.Fprintf(&line,
				"\033[38;2;%d;%d;%dm██",
				uint8(r>>8),
				uint8(g>>8),
				uint8(b>>8),
			)
		}

		line.WriteString("\033[0m")
		lines = append(lines, line.String())
	}

	return lines
}
