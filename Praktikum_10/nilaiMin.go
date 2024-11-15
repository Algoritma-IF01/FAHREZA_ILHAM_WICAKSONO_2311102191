package main

import (
	"fmt"
)

type arrInt [2023]int

func terkecil_1(tabInt arrInt, n int) int {
	var min int = tabInt[0]
	var j int = 1

	for j < n {
		if min > tabInt[j] {
			min = tabInt[j]
		}

		j++
	}

	return min
}

func terkecil_2(tabInt arrInt, n int) int {
	var idk int = 0
	var j int = 1

	for j < n {
		if tabInt[idk] > tabInt[j] {
			idk = j
		}

		j++
	}

	return idk
}

func main() {
	var (
		n      int
		tabInt arrInt
	)

	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scanln(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&tabInt[i])
	}

	indexTerkecil := terkecil_2(tabInt, n)

	fmt.Printf("Elemen dengan index terkecil berada di index ke-%d yang bernilai %d", indexTerkecil, tabInt[indexTerkecil])
}
