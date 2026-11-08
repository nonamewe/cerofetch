package main

import (
	"fmt"
	"os/exec"
	"strings"

	hardware "github.com/nonamewe/ceofetch/gets/hardware"
	sys "github.com/nonamewe/ceofetch/gets/sys"
)

func main() {
	cpu := hardware.CPU()
	gpu := hardware.GPU()
	ram := hardware.RAM()
	shell := sys.SHELL()
	distr := sys.DISTRONAME()
	core := sys.LINUX()
	desc := sys.DESCTOP()
	USER()
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

func USER() {
	usrname := exec.Command("whoami")
	usrnameo, err := usrname.Output()
	if err != nil {
		fmt.Println("user name get error")
		panic(err)
	}
	fmt.Println(strings.TrimSpace(string(usrnameo)))
}
