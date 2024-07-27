package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"syscall"
)

func getDpiProtectionStatus() string {
	fileName := "/tmp/dpi.run"
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return "OFF"
	}

	file, err := os.Open(fileName)
	if err != nil {
		return "OFF"
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			fmt.Println("Problem closing file", err)
		}
	}(file)
	buf := make([]byte, 1024)
	readed, err := file.Read(buf)
	if err != nil {
		fmt.Println(err)
		return "OFF"
	}
	pid := string(buf[:readed])
	pidInt, err := strconv.Atoi(pid)
	if err != nil {
		fmt.Println(err)
		return "OFF"
	}
	fmt.Printf("Parsed pid=%d\n", pidInt)
	process, err := os.FindProcess(pidInt)
	if err != nil {
		fmt.Println(err)
		return "OFF"
	}
	err = process.Signal(syscall.Signal(0))
	if err != nil {
		return "OFF"
	}

	return "ON"
}

func executeSimpleShellCommand(cmd string) {
	command := exec.Command(cmd)
	output, err := command.Output()
	if err != nil {
		return
	}
	fmt.Println(string(output))
}
