package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/switch", switchProtection)

	err := http.ListenAndServe(":8082", nil)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func homePage(w http.ResponseWriter, r *http.Request) {
	err := printDefaultPage(w)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func switchProtection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Switch protection")

	if getDpiProtectionStatus() == "OFF" {
		fmt.Println("Process doesn't exist")
		executeSimpleShellCommand("/etc/goodbyeDPI1.sh")
	} else {
		fmt.Println("Process exist. Disable DPI protection")
		executeSimpleShellCommand("/etc/disableDPIProtection.sh")
	}
}
