package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"regexp"
)

func CommandRunner(cmd string, args []string) error {
	command := exec.Command(cmd, args...)
	command.Stderr = os.Stderr
	command.Stdin = os.Stdin
	return command.Run()

}

func macChanger(newMAC string, iface string) error {

	commands := [][]string{
		[]string{"ifcongig", iface, "down"},
		[]string{"ifconfig", iface, "hw", "ether", newMAC},
		[]string{"ifcongig", iface, "up"},
	}

	for _, args := range commands {
		if err := CommandRunner("sudo", args); err != nil {
			return err
		}
	}
	return nil
}

func main() {

	const MAC_PATTERN = "(.){18}"
	newMac := flag.String("mac", "", "Enter a new and valid mac address")
	iFace := flag.String("iface", "", "Enter the trageting interface")
	flag.Parse()

	reg, err := regexp.Compile(MAC_PATTERN)
	if err != nil {
		panic("error regix")
	}

	if !reg.Match([]byte(*newMac)) {
		panic("not matched")
	}

	
	fmt.Println("result : ", *newMac, *iFace)

}
