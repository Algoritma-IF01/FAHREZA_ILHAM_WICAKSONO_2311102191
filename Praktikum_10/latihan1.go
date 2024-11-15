package main

import "fmt"

func main() {
	var N int
	fmt.Print("masukan jumlah anak kelinci : ")
	fmt.Scan(&N)

	if N <= 0 || N > 1000 {
		fmt.Println("jumlah anak kelinci harus antara 1 dan 1000 ")
		return
	}

	weights := make([]float64, N)
	fmt.Println("Masukan berat anak kelinci : ")
	for i := 0; i < N; i++ {
		fmt.Scan(&weights[i])
	}

	minWeight, maxWeight := weights[0], weights[0]
	for _, weight := range weights[1:] {
		if weight < minWeight {
			minWeight = weight
		}
		if weight > maxWeight {
			maxWeight = weight
		}

	}

	fmt.Printf("Berat kelinci terkecil : %.2f\n", minWeight)
	fmt.Printf("Berat kelinci terbesar : %.2f\n", maxWeight)
}
