package main

import (
	"fmt"
)

func main() {
	var (
		beratKanan, beratKiri float64
		isOleng               bool
	)

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&beratKiri, &beratKanan)

		if beratKanan+beratKiri >= 150.0 || beratKanan < 0 || beratKiri < 0 {
			fmt.Println("Proses selesai")
			break
		}

		if beratKanan-beratKiri >= 9 || beratKiri-beratKanan >= 9 {
			isOleng = true
		} else {
			isOleng = false
		}

		fmt.Println("Sepeda motor pak Andi akan oleng: ", isOleng)
	}
}
