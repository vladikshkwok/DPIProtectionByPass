package utils

import (
	"fmt"
	"strconv"
)

func ToInt(raw string) int {
	if raw == "" {
		return 0
	}
	res, err := strconv.Atoi(raw)
	if err != nil {
		fmt.Println(err)
	}
	return res
}

func ToFloat(raw string) float64 {
	if raw == "" {
		return 0
	}
	res, err := strconv.ParseFloat(raw, 64)
	if err != nil {
		fmt.Println(err)
	}
	return res
}
