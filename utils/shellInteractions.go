package utils

import (
	"log"
	"os/exec"
)

func ExecuteSimpleShellCommand(cmd string) error {
	command := exec.Command(cmd)
	output, err := command.Output()
	if err != nil {
		return err
	}
	log.Println(string(output))
	return nil
}
