package utils

import (
	"awesomeProject/domain"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		key, value := parseLine(scanner.Text())
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
		res.Load1 = toFloat(splittedText[0])
		res.Load5 = toFloat(splittedText[1])
		res.Load15 = toFloat(splittedText[2])
		res.LastCreatedPid = toInt(splittedText[4])
	}
	return res
}

func parseLine(raw string) (key string, value int) {
	text := strings.ReplaceAll(raw[:len(raw)-2], " ", "")
	keyValue := strings.Split(text, ":")
	return keyValue[0], toInt(keyValue[1])
}

func toInt(raw string) int {
	if raw == "" {
		return 0
	}
	res, err := strconv.Atoi(raw)
	if err != nil {
		fmt.Println(err)
	}
	return res
}

func toFloat(raw string) float64 {
	if raw == "" {
		return 0
	}
	res, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		fmt.Println(err)
	}
	return res
}
