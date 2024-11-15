package main

import "fmt"

type arrBalita [100]float64

func hitungMinMax(arrBerat arrBalita, n int, bMin, bMax *float64) {
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	for i := 1; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i]
		}

		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i]
		}
	}
}

func rerata(arrBerat arrBalita, n int) float64 {
	var sum float64

	for i := 0; i < n; i++ {
		sum += arrBerat[i]
	}

	return sum / float64(n)
}

func main() {
	var (
		n          int
		arrBerat   arrBalita
		bMin, bMax float64
	)

	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scanln(&n)

	if n <= 0 || n > 100 {
		fmt.Println("Banyak data balita harus diantara 1-100")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Print("Masukkan berat balita ke-", i+1, ": ")
		fmt.Scanln(&arrBerat[i])
	}
	hitungMinMax(arrBerat, n, &bMin, &bMax)
	rataRata := rerata(arrBerat, n)

	fmt.Printf("Berat balita minimum: %.2f kg\n", bMin)
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)
	fmt.Printf("Rerata berat balita: %.2f kg", rataRata)
}
