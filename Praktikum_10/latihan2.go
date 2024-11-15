package main

import "fmt"

func main() {
	const nMax = 1000

	var (
		x, y  int
		berat [nMax]float64
		wadah []float64
	)

	fmt.Print("Masukkan jumlah ikan dan kapasitas per wadah: ")
	fmt.Scanln(&x, &y)

	fmt.Printf("Masukkan berat masing-masing ikan (jumlah = %d):\n", x)
	for i := 0; i < x; i++ {
		fmt.Scan(&berat[i])
	}

	for i := 0; i < x; i += y {
		total := 0.0

		for j := i; j < i+y && j < x; j++ {
			total += berat[j]
		}

		wadah = append(wadah, total)
	}

	fmt.Println("Total berat setiap wadah")

	for i, beratWadah := range wadah {
		fmt.Printf("Berat wadah %d : %.2f\n", i+1, beratWadah)
	}

	fmt.Println("Rata rata berat ikan di setiap wadah")

	for i, beratWadah := range wadah {
		fmt.Printf("Rata rata berat wadah %d : %.2f\n", i+1, beratWadah/float64(y))
	}
}
