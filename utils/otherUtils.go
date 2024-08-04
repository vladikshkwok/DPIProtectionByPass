package utils

import (
	"awesomeProject/domain"
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
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

func GetDomains() []string {
	file, err := os.Open("/etc/domains.txt")
	if err != nil {
		log.Printf("Error opening /etc/domains.txt: %v", err)
		return []string{}
	}
	defer file.Close()

	var domains []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domains = append(domains, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Printf("Error reading /etc/domains.txt: %v", err)
	}
	return domains
}

func AddDomain(domain string) error {
	file, err := os.OpenFile("/etc/domains.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Printf("Error opening /etc/domains.txt: %v", err)
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(domain + "\n"); err != nil {
		log.Printf("Error writing to /etc/domains.txt: %v", err)
		return err
	}
	return nil
}

func UpdateDomain(oldDomain, newDomain string) error {
	domains := GetDomains()
	for i, d := range domains {
		if d == oldDomain {
			domains[i] = newDomain
			break
		}
	}

	return writeDomainsToFile(domains)
}

func DeleteDomain(domain string) error {
	domains := GetDomains()
	for i, d := range domains {
		if d == domain {
			domains = append(domains[:i], domains[i+1:]...)
			break
		}
	}

	return writeDomainsToFile(domains)
}

func writeDomainsToFile(domains []string) error {
	file, err := os.OpenFile("/etc/domains.txt", os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		log.Printf("Error opening /etc/domains.txt: %v", err)
		return err
	}
	defer file.Close()

	for _, domain := range domains {
		if _, err := file.WriteString(domain + "\n"); err != nil {
			log.Printf("Error writing to /etc/domains.txt: %v", err)
			return err
		}
	}
	return nil
}
