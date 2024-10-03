package main

import (
	"fmt"
	// "strconv"
)

func main() {
	var (
		angka1, angka2, angka3, angka4, angka5 int16
		char1, char2, char3 string
	)

	fmt.Print("Masukkan 5 angka: ")
	fmt.Scanln(&angka1, &angka2, &angka3, &angka4, &angka5)

	fmt.Print("Masukkan 3 karakter: ")
	fmt.Scanln(&char1, &char2, &char3)

	fmt.Printf("%c%c%c%c%c \n", angka1, angka2, angka3, angka4, angka5)
	fmt.Printf("%c %c %c \n", char1[0]+1, char2[0]+1, char3[0]+1)
}
