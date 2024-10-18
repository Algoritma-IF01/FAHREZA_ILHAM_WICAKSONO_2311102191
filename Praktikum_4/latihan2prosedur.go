package main

import (
	"fmt"
)

func hitungSkor(soalBenar *int, skor *int) {
	var soal1, soal2, soal3, soal4, soal5, soal6, soal7, soal8 int

	*soalBenar = 0
	*skor = 0

	fmt.Scan(&soal1, &soal2, &soal3, &soal4, &soal5, &soal6, &soal7, &soal8)

	// Deklarasi dan inisalisasi array soal untuk kemudahan pengecekan
	soal := []int{soal1, soal2, soal3, soal4, soal5, soal6, soal7, soal8}

	for i := 0; i < 8; i++ {
		if soal[i] < 301 {
			*soalBenar++
			*skor += soal[i]
		}
	}
}

func main() {
	var (
		skor, soalBenar                 int
		namaPeserta, pemenang           string
		soalBenarTerbesar, skorTerkecil int
	)

	for {
		fmt.Print("Masukkan nama peserta dan lama pengerjaan: ")
		fmt.Scan(&namaPeserta)

		if namaPeserta == "Selesai" || namaPeserta == "selesai" {
			break
		}

		hitungSkor(&soalBenar, &skor)

		if soalBenar > soalBenarTerbesar || (soalBenar == soalBenarTerbesar && skor < skorTerkecil) {
			soalBenarTerbesar = soalBenar
			skorTerkecil = skor
			pemenang = namaPeserta
		}
	}

	fmt.Println(pemenang, soalBenarTerbesar, skorTerkecil)
}
