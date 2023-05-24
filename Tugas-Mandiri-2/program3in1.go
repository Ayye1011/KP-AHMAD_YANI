package main

import (
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Println("Pilih program yang akan digunakan:")
	fmt.Println("1. Menentukan bilangan prima")
	fmt.Println("2. Menentukan bilangan kelipatan 7")
	fmt.Println("3. Menghitung luas trapesium")
	fmt.Println("0. Keluar")

	// Menerima input pilihan dari pengguna
	var choice int
	fmt.Print("Masukkan pilihan Anda: ")
	fmt.Scanln(&choice)

	// Melakukan pengkondisian berdasarkan pilihan pengguna
	switch choice {
	case 1:
		runPrimeProgram()
	case 2:
		runMultipleOf7Program()
	case 3:
		runTrapeziumAreaProgram()
	case 0:
		fmt.Println("Terima kasih telah menggunakan program ini!")
		os.Exit(0)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}

func runPrimeProgram() {
	fmt.Println("Anda memilih program Menentukan Bilangan Prima.")

	// Menerima input bilangan dari pengguna
	var num int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scanln(&num)

	isPrime := true
	if num <= 1 {
		isPrime = false
	} else {
		// Mengecek apakah ada faktor yang membagi num selain dari 1 dan num itu sendiri
		// Karena faktor yang membagi num tidak akan melebihi akar kuadrat dari num
		sqrt := int(math.Sqrt(float64(num)))
		for i := 2; i <= sqrt; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
	}

	if isPrime {
		fmt.Println(num, "adalah bilangan prima")
	} else {
		fmt.Println(num, "bukan bilangan prima")
	}
}

func runMultipleOf7Program() {
	fmt.Println("Anda memilih program Menentukan Bilangan Kelipatan 7.")

	// Menerima input bilangan dari pengguna
	var num int
	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scanln(&num)

	if num%7 == 0 {
		fmt.Println(num, "adalah bilangan kelipatan 7")
	} else {
		fmt.Println(num, "bukan bilangan kelipatan 7")
	}
}

func runTrapeziumAreaProgram() {
	fmt.Println("Anda memilih program Menghitung Luas Trapesium.")

	// Menerima input panjang alas, panjang atas, dan tinggi dari pengguna
	var a, b, h float64
	fmt.Print("Masukkan panjang alas: ")
	fmt.Scanln(&a)
	fmt.Print("Masukkan panjang atas: ")
	fmt.Scanln(&b)
	fmt.Print("Masukkan tinggi: ")
	fmt.Scanln(&h)

	area := (a + b) * h / 2
	fmt.Println("Luas trapesium:", area)
}
