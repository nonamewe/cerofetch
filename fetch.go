package main

import (
	"fmt"
	"os/exec"
	"strings"

	hardware "github.com/nonamewe/ceofetch/gets/hardware"
	sys "github.com/nonamewe/ceofetch/gets/sys"
)

type SystemInfo struct {
	Name  string
	CPU   string
	GPU   string
	RAM   string
	Shell string
	Distr string
	Core  string
	Desc  string
}

func main() {
	info := SystemInfo{
		Name:  USER(),
		CPU:   hardware.CPU(),
		GPU:   hardware.GPU(),
		RAM:   hardware.RAM(),
		Shell: sys.SHELL(),
		Distr: sys.DISTRONAME(),
		Core:  sys.LINUX(),
		Desc:  sys.DESCTOP(),
	}
	fmt.Println(info.Name)
	fmt.Println("---------------")

	fmt.Println("NAME:", info.Distr)
	fmt.Println("KERNEL:", info.Core)
	fmt.Println("DESCTOP:", info.Desc)
	fmt.Println("SHELL:", info.Shell)

	fmt.Println("---------------")

	fmt.Println("CPU:", info.CPU)
	fmt.Println("GPU:", info.GPU)
	fmt.Println("RAM:", info.RAM)

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
