package main

import (
	"fmt"
)

var ColorGRN = "\033[92m"
var ColorEND = "\033[0m"

func IndexOf(s string, arr []string) int {
	for i, d := range arr {
		if s == d {
			return i
		}
	}
	return -1
}

func PrintOkMsg(s string) {
	fmt.Println(fmt.Sprintf("%s%s%s", ColorGRN, s, ColorEND))
}