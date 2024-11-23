package main

import "fmt"

type arrRumahKerabat [][]int

func selectionSort(rumahKerabat *arrRumahKerabat) {
	for i := 0; i < len(*rumahKerabat); i++ {
		for j := 0; j < len((*rumahKerabat)[i])-1; j++ {
			var idxMin = j

			for k := j + 1; k < len((*rumahKerabat)[i]); k++ {
				if (*rumahKerabat)[i][k] < (*rumahKerabat)[i][idxMin] {
					idxMin = k
				}
			}

			if idxMin != j {
				(*rumahKerabat)[i][j], (*rumahKerabat)[i][idxMin] = (*rumahKerabat)[i][idxMin], (*rumahKerabat)[i][j]
			}
		}
	}
}

func main() {
	var (
		rumahKerabat    arrRumahKerabat
		nDaerah, nRumah int
	)

	fmt.Print("Masukkan jumlah daerah: ")
	fmt.Scanln(&nDaerah)

	for i := 0; i < nDaerah; i++ {
		fmt.Print("Masukkan jumlah rumah: ")
		fmt.Scanln(&nRumah)

		daerah := make([]int, 0, nRumah)
		fmt.Println("Daerah ke", i+1)

		for j := 0; j < nRumah; j++ {
			var noRumah int
			fmt.Printf("Masukkan no. rumah ke-%d: ", j+1)
			fmt.Scanln(&noRumah)
			daerah = append(daerah, noRumah)
		}

		rumahKerabat = append(rumahKerabat, daerah)
	}

	fmt.Println("\nData sebelum diurutkan:")
	for i, daerah := range rumahKerabat {
		fmt.Printf("Daerah %d: %v\n", i+1, daerah)
	}

	selectionSort(&rumahKerabat)

	fmt.Println("\nData setelah diurutkan:")
	for i, daerah := range rumahKerabat {
		fmt.Printf("Daerah %d: %v\n", i+1, daerah)
	}
}
