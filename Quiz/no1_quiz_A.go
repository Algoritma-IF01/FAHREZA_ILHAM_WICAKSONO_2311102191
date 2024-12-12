package main

import (
	"fmt"
)

func hitungTarif(jam, menit float64, member bool) float64 {
	var totalTarif float64

	if member {
		totalTarif += jam * 3500

		if jam > 1 && menit > 10 {
			totalTarif += menit * (3500.0 / 60.0)
		} else if jam < 1 {
			totalTarif += menit * (3500.0 / 60.0)
		}
	} else {
		totalTarif += jam * 5000

		if jam > 1 && menit > 10 {
			totalTarif += menit * (5000.0 / 60.0)
		} else if jam < 1 {
			totalTarif += menit * (5000.0 / 60.0)
		}
	}

	if jam > 3 {
		totalTarif -= totalTarif * (10.0 / 100.0)
	}

	return totalTarif
}

func main() {
	var (
		jam, menit float64
		member     bool
	)

	fmt.Print("Masukkan jam, menit, dan membership: ")
	fmt.Scanln(&jam, &menit, &member)

	fmt.Println(hitungTarif(jam, menit, member))
}
