package main

import (
	"fmt"
)

func cetakPerfectNumber(a, b int) {
	if a >= b {
		return
	}

	for i := a; i <= b; i++ {
		jumlah := 0

		for j := 1; j < i; j++ {
			if i%j == 0 {
				jumlah += j
			}
		}

		if jumlah == i {
			fmt.Print(i, " ")
		}
	}
}

func main() {
	var (
		a, b int
	)

	fmt.Print("Masukkan 2 bilangan: ")
	fmt.Scanln(&a, &b)

	cetakPerfectNumber(a, b)
}
