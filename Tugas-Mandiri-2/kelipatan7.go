package main

import (
	"fmt"
)

func isMultipleOfSeven(n int) bool {
	if n%7 == 0 {
		return true
	}
	return false
}

func main() {
	// Menerima input dari pengguna
	var num int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scanln(&num)

	// Memeriksa apakah bilangan merupakan kelipatan 7 atau bukan
	if isMultipleOfSeven(num) {
		fmt.Printf("%d adalah kelipatan dari 7.\n", num)
	} else {
		fmt.Printf("%d bukan kelipatan dari 7.\n", num)
	}
}
