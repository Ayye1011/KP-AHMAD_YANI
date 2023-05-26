package main

import "fmt"

type Car struct {
	Type   string
	FuelIn float64
}

func (c Car) CalculateDistance() float64 {
	fuelEfficiency := 1.5 // Bahan bakar yang digunakan per mill

	distance := c.FuelIn * fuelEfficiency
	return distance
}

func main() {
	var bahanBakar uint

	fmt.Print("Masukkan jumlah bahan bakar anda = ")
	fmt.Scanln(&bahanBakar)
	car := Car{FuelIn: float64(bahanBakar)}
	fmt.Printf("Jarak yang dapat ditempuh oleh Kendaraan Anda%s dengan bahan bakar %.2f L: %.2f mill\n", car.Type, car.FuelIn, car.CalculateDistance())

}
