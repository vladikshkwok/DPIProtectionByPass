package main

import (
	"awesomeProject/utils"
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
	err := utils.PrintDefaultPage(w)

	if err != nil {
		fmt.Println(err)
		return
	}
}

func switchProtection(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Switch protection")

	if utils.GetDpiProtectionStatus() == "OFF" {
		fmt.Println("Process doesn't exist")
		utils.ExecuteSimpleShellCommand("/etc/goodbyeDPI.sh")
	} else {
		fmt.Println("Process exist. Disable DPI protection")
		utils.ExecuteSimpleShellCommand("/etc/disableDPIProtection.sh")
	}
}
