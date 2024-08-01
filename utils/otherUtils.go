package utils

import (
	"awesomeProject/domain"
	"errors"
	"fmt"
	"log"
	"strconv"
)

func ToInt(raw string) (int, error) {
	res, err := strconv.Atoi(raw)
	if err != nil {
		log.Printf("Error converting string to int: %v", err)
		return 0, err
	}
	return res, nil
}

func ToFloat(raw string) (float64, error) {
	res, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		log.Printf("Error converting string to float: %v", err)
		return 0, err
	}
	return res, nil
}

func GetRouterStats(mock bool) []domain.RouterStat {
	if mock {
		return []domain.RouterStat{
			domain.NewRouterStat("MockStat", "Some cool stat"),
			domain.NewRouterStat("MockStat2", "Another cool stat"),
		}
	}

	return []domain.RouterStat{
		domain.NewRouterStat("LoadAverage", GetLoadAverage().String()),
		domain.NewRouterStat("MemoryStats", GetMemoryStats().String()),
	}
}

func SwitchProtection() error {
	currentStatus := GetDpiProtectionStatus()
	fmt.Println("Switch protection")

	switch currentStatus {
	case "OFF":
		fmt.Println("Process doesn't exist")
		if err := ExecuteSimpleShellCommand("/etc/goodbyeDPI.sh"); err != nil {
			log.Printf("Error executing /etc/goodbyeDPI.sh: %v", err)
			return err
		}
	case "ON":
		fmt.Println("Process exists. Disable DPI protection")
		if err := ExecuteSimpleShellCommand("/etc/disableDPIProtection.sh"); err != nil {
			log.Printf("Error executing /etc/disableDPIProtection.sh: %v", err)
			return err
		}
	default:
		return errors.New("unknown DPI protection status")
	}

	return nil
}
