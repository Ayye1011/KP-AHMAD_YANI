package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}

	// Mengecek apakah bilangan n habis dibagi oleh bilangan dari 2 hingga akar(n)
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	// Menerima input dari pengguna
	var num int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scanln(&num)

	// Memeriksa apakah bilangan merupakan bilangan prima atau bukan
	if isPrime(num) {
		fmt.Printf("%d adalah bilangan prima.\n", num)
	} else {
		fmt.Printf("%d bukan bilangan prima.\n", num)
	}
}
