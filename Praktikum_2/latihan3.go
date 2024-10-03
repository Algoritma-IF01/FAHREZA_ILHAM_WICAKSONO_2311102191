package main

import (
	"fmt"
	// "strconv"
)

func main() {
	var (
		r, luas, volume float64
	)

	const phi float64 = 3.1415926535

	fmt.Print("Masukkan jejari: ")
	fmt.Scanln(&r)

	luas = 4 * phi * r * r
	volume = (4.0 / 3.0) * phi * r * r * r

	fmt.Printf("Bola dengan jejari %.2f memiliki volume %.4f dan luas kulit %.4f\n", r, volume, luas)
}
