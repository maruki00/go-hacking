package main

import (
	"log"
	"os"
	"os/exec"
)

func command(cmd string) error {

	cmd_obj := exec.Command(cmd, []string{"-lag"}...)
	cmd_obj.Stdout = os.Stdout
	cmd_obj.Stderr = os.Stderr

	return cmd_obj.Run()

}
func main() {
	if err := command("ls"); err != nil {
		log.Fatal("command error, ", err.Error())
	}
}
