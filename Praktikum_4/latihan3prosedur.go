package main

import (
	"fmt"
)

func cetakDeret(n int) {
	if n > 1000000 {
		fmt.Println("Bilangan harus lebih kecil dari 1000000")
	}
	fmt.Print(n, " ")

	for {
		if n == 1 {
			break
		}

		if n%2 == 0 {
			n = n / 2
			fmt.Print(n, " ")
		} else {
			n = 3*n + 1
			fmt.Print(n, " ")
		}
	}
}

func main() {
	var n int

	fmt.Print("Masukkan sebuah bilangan: ")
	fmt.Scanln(&n)

	cetakDeret(n)
}
