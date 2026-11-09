package main

import (
	"fmt"
	"image"
	_ "image/png"
	"os/exec"
	"strings"

	hardware "github.com/nonamewe/ceofetch/gets/hardware"
	sys "github.com/nonamewe/ceofetch/gets/sys"
	logocheck "github.com/nonamewe/ceofetch/logocheck"
)

type SystemInfo struct {
	Distro string
	Logo   image.Image
	Name   string
	CPU    string
	GPU    string
	RAM    string
	Shell  string
	Distr  string
	Core   string
	Desc   string
}

func main() {
	info := SystemInfo{
		Distro: logocheck.Distro,
		Logo:   logocheck.Image(),
		Name:   USER(),
		CPU:    hardware.CPU(),
		GPU:    hardware.GPU(),
		RAM:    hardware.RAM(),
		Shell:  sys.SHELL(),
		Distr:  sys.DISTRONAME(),
		Core:   sys.LINUX(),
		Desc:   sys.DESCTOP(),
	}
	logo := logocheck.RenderImage(info.Logo, 100)
	switch info.Distro {
	case "arch":
		logo = logocheck.RenderImage(info.Logo, 10)
	case "gentoo":
		logo = logocheck.RenderImage(info.Logo, 45)
	}

	text := []string{
		info.Name,
		"---------------",
		"NAME: " + info.Distr,
		"KERNEL: " + info.Core,
		"DESKTOP: " + info.Desc,
		"SHELL: " + info.Shell,
		"---------------",
		"CPU: " + info.CPU,
		"GPU: " + info.GPU,
		"RAM: " + info.RAM,
	}
	maxLines := len(logo)
	if len(text) > maxLines {
		maxLines = len(text)
	}

	for i := 0; i < maxLines; i++ {
		left := ""
		right := ""

		if i < len(logo) {
			left = logo[i]
		}

		if i < len(text) {
			right = text[i]
		}

		fmt.Printf("%-40s   %s\n", left, right)
	}

}

func USER() string {
	usrname := exec.Command("whoami")
	usrnameo, err := usrname.Output()
	if err != nil {
		return "Unknown"
	}
	name := strings.TrimSpace(string(usrnameo))
	return name
}
