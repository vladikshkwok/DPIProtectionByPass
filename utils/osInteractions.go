package utils

import (
	"awesomeProject/domain"
	"bufio"
	"errors"
	"log"
	"os"
	"strings"
	"syscall"
)

func GetMemoryStats() domain.Memory {
	file, err := os.Open("/proc/meminfo")
	if err != nil {
		log.Printf("Error opening /proc/meminfo: %v", err)
		return domain.Memory{}
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
	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning /proc/meminfo: %v", err)
	}
	return res
}

func GetLoadAverage() domain.LoadAvg {
	file, err := os.Open("/proc/loadavg")
	if err != nil {
		log.Printf("Error opening /proc/loadavg: %v", err)
		return domain.LoadAvg{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	res := domain.LoadAvg{}
	if scanner.Scan() {
		text := scanner.Text()
		splittedText := strings.Split(text, " ")
		res.Load1, _ = ToFloat(splittedText[0])
		res.Load5, _ = ToFloat(splittedText[1])
		res.Load15, _ = ToFloat(splittedText[2])
		res.LastCreatedPid, _ = ToInt(splittedText[4])
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning /proc/loadavg: %v", err)
	}
	return res
}

func GetDpiProtectionStatus() string {
	fileName := "/tmp/dpi.run"
	if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
		return "OFF"
	}

	file, err := os.Open(fileName)
	if err != nil {
		log.Printf("Error opening %s: %v", fileName, err)
		return "OFF"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	pid := 0
	for scanner.Scan() {
		pid, _ = ToInt(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error scanning %s: %v", fileName, err)
	}

	if !findProcessAndCheckForLiveness(pid) {
		return "OFF"
	}

	return "ON"
}

func CheckProcessLiveliness(process *os.Process) bool {
	err := process.Signal(syscall.Signal(0))
	return err == nil
}

func findProcessAndCheckForLiveness(pid int) bool {
	process, err := os.FindProcess(pid)
	if err != nil {
		log.Printf("Error finding process with PID %d: %v", pid, err)
		return false
	}
	return CheckProcessLiveliness(process)
}

func parseMemoryStatsLine(raw string) (string, int) {
	text := strings.TrimSpace(raw)
	keyValue := strings.Split(text, ":")
	key := strings.TrimSpace(keyValue[0])
	value, _ := ToInt(strings.Fields(keyValue[1])[0])
	return key, value
}
