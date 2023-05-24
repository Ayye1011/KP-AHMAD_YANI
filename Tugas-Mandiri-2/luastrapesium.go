package main

import (
	"fmt"
)

func calculateArea(a float64, b float64, h float64) float64 {
	// Menghitung luas trapesium menggunakan rumus: (a + b) * h / 2
	area := (a + b) * h / 2
	return area
}

func main() {
	// Menerima input dari pengguna
	var base1, base2, height float64
	fmt.Println("Program Menghitung Luas Trapesium")
	fmt.Print("Masukkan panjang alas atas: ")
	fmt.Scanln(&base1)
	fmt.Print("Masukkan panjang alas bawah: ")
	fmt.Scanln(&base2)
	fmt.Print("Masukkan tinggi trapesium: ")
	fmt.Scanln(&height)

	// Menghitung luas trapesium
	area := calculateArea(base1, base2, height)

	// Menampilkan hasil
	fmt.Printf("Luas trapesium: %.2f\n", area)
}