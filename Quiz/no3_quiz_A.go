package main

import (
	"fmt"
)

func jumlahPertemuanIteratif(x, y int) int {
	jumlahHari := 0

	for i := 1; i <= 365; i++ {
		if i%x == 0 && i%y != 0 {
			jumlahHari++
		}
	}

	return jumlahHari
}

func jumlahPertemuanRekursif(x, y, hari int) int {
	if hari > 365 {
		return 0
	}

	if hari%x == 0 && hari%y != 0 {
		return 1 + jumlahPertemuanRekursif(x, y, hari+1)
	}

	return jumlahPertemuanRekursif(x, y, hari+1)
}

func main() {
	var x, y int

	fmt.Print("Masukkan 2 bilangan: ")
	fmt.Scanln(&x, &y)

	// fmt.Println(jumlahPertemuanIteratif(x, y))
	fmt.Println(jumlahPertemuanRekursif(x, y, 1))
}
