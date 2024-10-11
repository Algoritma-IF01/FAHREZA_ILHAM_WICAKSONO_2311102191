package main

import (
	"fmt"
)

func main() {
	// deklarasi variabel
	var (
		warna1, warna2, warna3, warna4 string
	)

	urutanWarna := []string{"merah", "kuning", "hijau", "ungu"}
	hasil := true

	// perulangan 5 kali
	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d \n", i)

		fmt.Print("Masukkan warna pertama: ")
		fmt.Scanln(&warna1)
		fmt.Print("Masukkan warna kedua: ")
		fmt.Scanln(&warna2)
		fmt.Print("Masukkan warna ketiga: ")
		fmt.Scanln(&warna3)
		fmt.Print("Masukkan warna keempat: ")
		fmt.Scanln(&warna4)

		// Cek apakah urutan sama
		if warna1 != urutanWarna[0] || warna2 != urutanWarna[1] || warna3 != urutanWarna[2] || warna4 != urutanWarna[3] {
			hasil = false
		}
	}

	// tampilkan hasil
	fmt.Printf("Berhasil: %t", hasil)
}
