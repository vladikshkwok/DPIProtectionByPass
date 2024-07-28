package utils

import (
	"fmt"
	"os/exec"
)

func ExecuteSimpleShellCommand(cmd string) {
	command := exec.Command(cmd)
	output, err := command.Output()
	if err != nil {
		return
	}
	fmt.Println(string(output))
}
