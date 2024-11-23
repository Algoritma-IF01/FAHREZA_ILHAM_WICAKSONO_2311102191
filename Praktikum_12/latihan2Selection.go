package main

import "fmt"

type arrRumahKerabat [][]int

func selectionSortGanjil(arrRumah *[]int) {
	for i := 0; i < len(*arrRumah); i++ {
		idxMin := i

		for j := i + 1; j < len(*arrRumah); j++ {
			if (*arrRumah)[j] < (*arrRumah)[idxMin] {
				idxMin = j
			}
		}

		if idxMin != i {
			(*arrRumah)[i], (*arrRumah)[idxMin] = (*arrRumah)[idxMin], (*arrRumah)[i]
		}
	}
}

func selectionSortGenap(arrRumah *[]int) {
	for i := 0; i < len(*arrRumah); i++ {
		idxMin := i

		for j := i + 1; j < len(*arrRumah); j++ {
			if (*arrRumah)[j] > (*arrRumah)[idxMin] {
				idxMin = j
			}
		}

		if idxMin != i {
			(*arrRumah)[i], (*arrRumah)[idxMin] = (*arrRumah)[idxMin], (*arrRumah)[i]
		}
	}
}

func urutkanGanjilGenap(rumahKerabat *arrRumahKerabat) {
	for i := 0; i < len(*rumahKerabat); i++ {
		var rumahGanjil, rumahGenap []int

		for _, noRumah := range (*rumahKerabat)[i] {
			if noRumah%2 == 0 {
				rumahGenap = append(rumahGenap, noRumah)
			} else {
				rumahGanjil = append(rumahGanjil, noRumah)
			}
		}

		selectionSortGanjil(&rumahGanjil)
		selectionSortGenap(&rumahGenap)

		hasil := append(rumahGanjil, rumahGenap...)
		fmt.Printf("Daerah %d setelah diurutkan: %v\n", i+1, hasil)
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

	urutkanGanjilGenap(&rumahKerabat)
}
