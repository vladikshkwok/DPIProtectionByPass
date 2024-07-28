package utils

import (
	"awesomeProject/domain"
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"syscall"
)

func GetMemoryStats() domain.Memory {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := domain.Memory{}
	for scanner.Scan() {
		key, value := parseMemoryStatsLine(scanner.Text())
		switch key {
		case "MemTotal":
			res.MemTotal = value
		case "MemFree":
			res.MemFree = value
		case "MemAvailable":
			res.MemAvailable = value
		}
	}
	return res
}

func GetLoadAverage() domain.LoadAvg {
	file, err := os.Open("/proc/loadavg")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	res := domain.LoadAvg{}
	for scanner.Scan() {
		text := scanner.Text()
		splittedText := strings.Split(text, " ")
		res.Load1 = ToFloat(splittedText[0])
		res.Load5 = ToFloat(splittedText[1])
		res.Load15 = ToFloat(splittedText[2])
		res.LastCreatedPid = ToInt(splittedText[4])
	}
	return res
}

func GetDpiProtectionStatus() string {
	fileName := "/tmp/dpi.run"
	pid := 0
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return "OFF"
	}

	file, err := os.Open(fileName)
	if err != nil {
		return "OFF"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pid = ToInt(scanner.Text())
	}
	fmt.Printf("Process pid=%d\n", pid)

	if !findProcessAndCheckForLiveness(pid) {
		return "OFF"
	}

	return "ON"
}

func CheckProcessLiveliness(process *os.Process) bool {
	err := process.Signal(syscall.Signal(0))
	if err != nil {
		return false
	}

	return true
}

func findProcessAndCheckForLiveness(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if !CheckProcessLiveliness(process) {
		return false
	}
	return true
}

func parseMemoryStatsLine(raw string) (key string, value int) {
	text := strings.ReplaceAll(raw[:len(raw)-2], " ", "")
	keyValue := strings.Split(text, ":")
	return keyValue[0], ToInt(keyValue[1])
}
