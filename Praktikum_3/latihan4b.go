package main

import (
	"fmt"
)

func akar(k int) float64 {
	hasil := 1.0

	for i := 0; i < k; i++ {
		pembilang := float64((4*i + 2) * (4*i + 2))
		penyebut := float64((4*i + 1) * (4*i + 3))
		hasil *= pembilang / penyebut
	}

	return hasil
}

func main() {
	var (
		k int
	)

	fmt.Print("Nilai K = ")
	fmt.Scanln(&k)

	fmt.Printf("Nilai akar 2 = %.10f", akar(k))
}
