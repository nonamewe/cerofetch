package main

import (
	"fmt"
	"os/exec"
	"strings"

	hardware "github.com/nonamewe/ceofetch/gets/hardware"
	sys "github.com/nonamewe/ceofetch/gets/sys"
)

func main() {
	name := USER()
	cpu := hardware.CPU()
	gpu := hardware.GPU()
	ram := hardware.RAM()
	shell := sys.SHELL()
	distr := sys.DISTRONAME()
	core := sys.LINUX()
	desc := sys.DESCTOP()
	fmt.Println(name)
	fmt.Println("---------------")

	fmt.Println("NAME:", distr)
	fmt.Println("CORE:", core)
	fmt.Println("DESCTOP:", desc)
	fmt.Println("SHELL:", shell)

	fmt.Println("---------------")

	fmt.Println("CPU:", cpu)
	fmt.Println("GPU:", gpu)
	fmt.Println("RAM:", ram)

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
