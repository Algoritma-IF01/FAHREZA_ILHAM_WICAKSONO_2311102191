package main

import (
	"fmt"
	"math"
)

func isiArray(arr []int, n int) {
	var bil int

	for i := 0; i < n; i++ {
		fmt.Print("Masukkan bilangan ke-", i+1, ": ")
		fmt.Scanln(&bil)

		arr[i] = bil
	}

	fmt.Println()
}

func cetakArray(arr []int, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}

	fmt.Println()

}

func cetakArrayGanjil(arr []int, n int) {
	for i := 0; i < n; i++ {
		if arr[i]%2 != 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

func cetakArrayGenap(arr []int, n int) {
	for i := 0; i < n; i++ {
		if arr[i]%2 == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

func cetakArrayKelipatan(arr []int, n int) {
	var x int

	fmt.Print("Masukkan bilangan x: ")
	fmt.Scanln(&x)

	for i := 0; i < n; i++ {
		if arr[i]%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}

	fmt.Println()
}

func hapusArray(arr []int, n *int) {
	var index int

	fmt.Print("Masukkan index array : ")
	fmt.Scanln(&index)

	if index > *n {
		fmt.Println("Index tidak valid")
		return
	}

	arr = append(arr[:index], arr[index+1:]...)
	*n--
	cetakArray(arr[:], *n)
}

func rataRataArray(arr []int, n int) float64 {
	var (
		hasil int
	)
	for i := 0; i < n; i++ {
		hasil += arr[i]
	}

	return float64(hasil) / float64(n)
}

func simpanganBaku(arr []int, n int) float64 {
	var (
		jumlah float64
	)

	mean := rataRataArray(arr[:], n)

	for i := 0; i < n; i++ {
		diff := float64(arr[i]) - mean
		jumlah += diff * diff
	}

	variansi := jumlah / float64(n)

	return math.Sqrt(variansi)
}

func frekuensiBilangan(arr []int, n int) int {
	var bil, count int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&bil)

	for i := 0; i < n; i++ {
		if arr[i] == bil {
			count++
		}
	}

	return count
}

func main() {
	const nMax = 1000

	var (
		arr [nMax]int
		n   int
	)

	fmt.Print("Masukkan index array: ")
	fmt.Scanln(&n)

	if n > nMax {
		fmt.Println("Index melebihi batas (1000)")
		return
	}

	fmt.Println("Mengisi Array")
	isiArray(arr[:], n)

	fmt.Println("Menampilkan Array")
	cetakArray(arr[:], n)

	fmt.Println("Menampilkan Array Ganjil")
	cetakArrayGanjil(arr[:], n)

	fmt.Println("Menampilkan Array Genap")
	cetakArrayGenap(arr[:], n)

	fmt.Println("Menampilkan Array Kelipatan x")
	cetakArrayKelipatan(arr[:], n)

	fmt.Println("Menghapus elemen array")
	hapusArray(arr[:], &n)

	rataRata := rataRataArray(arr[:], n)
	fmt.Print("Rata-rata Array: ", rataRata)
	fmt.Println()

	simpanganBaku := simpanganBaku(arr[:], n)
	fmt.Print("Simpangan Baku Array: ", simpanganBaku)
	fmt.Println()

	fmt.Println("Frekuensi suatu Array")
	frekuensi := frekuensiBilangan(arr[:], n)
	fmt.Println("Frekuensi: ", frekuensi)
}
