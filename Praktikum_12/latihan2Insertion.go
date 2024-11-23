package main

import (
	"fmt"
)

const nMax int = 7919

type DaftarBuku = [nMax]Buku

type Buku struct {
	id, judul, penulis, penerbit string
	eksemplar, tahun, rating     int
}

func daftarkanBuku(pustaka *DaftarBuku, n int) {
	for i := 0; i < n; i++ {
		fmt.Println("|========= Masukkan Data Buku =========|")

		fmt.Print("Masukkan ID buku: ")
		fmt.Scanln(&pustaka[i].id)

		fmt.Print("Masukkan judul buku: ")
		fmt.Scanln(&pustaka[i].judul)

		fmt.Print("Masukkan penulis buku: ")
		fmt.Scanln(&pustaka[i].penulis)

		fmt.Print("Masukkan penerbit buku: ")
		fmt.Scanln(&pustaka[i].penerbit)

		fmt.Print("Masukkan eksemplar buku: ")
		fmt.Scanln(&pustaka[i].eksemplar)

		fmt.Print("Masukkan tahun buku: ")
		fmt.Scanln(&pustaka[i].tahun)

		fmt.Print("Masukkan rating buku: ")
		fmt.Scanln(&pustaka[i].rating)
	}
}

func cetakFavorit(pustaka DaftarBuku, n int) {
	idxMax := 0

	for i := 1; i < n; i++ {
		if pustaka[i].rating > pustaka[idxMax].rating {
			idxMax = i
		}
	}

	fmt.Println("|========= BUKU TERFAVORIT =========|")
	fmt.Println("Judul		:", pustaka[idxMax].judul)
	fmt.Println("Penulis	:", pustaka[idxMax].penulis)
	fmt.Println("Penerbit	:", pustaka[idxMax].penerbit)
	fmt.Println("Tahun		:", pustaka[idxMax].tahun)
	fmt.Println("|===================================|")
}

func urutBuku(pustaka *DaftarBuku, n int) {
	var temp Buku

	for i := 0; i < n; i++ {
		temp = (*pustaka)[i]
		j := i

		for j > 0 && temp.rating > (*pustaka)[j-1].rating {
			(*pustaka)[j] = (*pustaka)[j-1]
			j--
		}

		(*pustaka)[j] = temp
	}
}

func cetak5Tertinggi(pustaka DaftarBuku, n int) {
	fmt.Println("|========= 5 BUKU TERBAIK =========|")
	for i := 0; i < 5 && i < n; i++ {
		fmt.Printf("%d. %s\n", i+1, pustaka[i].judul)
	}
	fmt.Println("|==================================|")
}

func cariBuku(pustaka DaftarBuku, n int, rating int) {
	start := 0
	end := n - 1

	for start <= end {
		mid := (start + end) / 2

		if pustaka[mid].rating == rating {
			fmt.Println("|========= DATA BUKU =========|")
			fmt.Println("Judul     :", pustaka[mid].judul)
			fmt.Println("Penulis   :", pustaka[mid].penulis)
			fmt.Println("Penerbit  :", pustaka[mid].penerbit)
			fmt.Println("Tahun     :", pustaka[mid].tahun)
			fmt.Println("|=============================|")
			return
		} else if pustaka[mid].rating < rating {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	fmt.Println("Tidak ada buku dengan rating seperti itu")
}

func main() {
	var (
		Pustaka          DaftarBuku
		nPustaka, rating int
	)

	fmt.Print("Masukkan jumlah buku yang ingin didaftarkan: ")
	fmt.Scanln(&nPustaka)

	if nPustaka <= 0 || nPustaka > nMax {
		fmt.Println("Jumlah buku harus diantara 1 -", nMax)
		return
	}

	daftarkanBuku(&Pustaka, nPustaka)
	fmt.Println()

	cetakFavorit(Pustaka, nPustaka)
	fmt.Println()

	urutBuku(&Pustaka, nPustaka)
	fmt.Println()

	cetak5Tertinggi(Pustaka, nPustaka)
	fmt.Println()

	fmt.Print("Masukkan rating buku yang dicari: ")
	fmt.Scanln(&rating)

	cariBuku(Pustaka, nPustaka, rating)
}
