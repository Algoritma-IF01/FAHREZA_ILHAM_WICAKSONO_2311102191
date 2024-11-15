package main

import "fmt"

type Mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

type arrMhs [2023]Mahasiswa

func IPK_2(T arrMhs, n int) int {
	var idx int = 0
	var j int = 1

	for j < n {
		if T[idx].ipk < T[j].ipk {
			idx = j
		}

		j++
	}

	return idx
}

func main() {
	var n int
	var data arrMhs

	fmt.Print("Masukkan jumlah mahasiswa (N <= 2023) : ")
	fmt.Scan(&n)

	if n <= 0 || n > 2023 {
		fmt.Println("Jumalah mahasiswa harus antara 1 dan 2023")
		return
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data mahasiswa ke-%d\n", i+1)
		fmt.Print("Nama: ")
		fmt.Scan(&data[i].nama)
		fmt.Print("NIM: ")
		fmt.Scan(&data[i].nim)
		fmt.Print("Kelas: ")
		fmt.Scan(&data[i].kelas)
		fmt.Print("Jurusan: ")
		fmt.Scan(&data[i].jurusan)
		fmt.Print("IPK: ")
		fmt.Scan(&data[i].ipk)
	}

	idxTertinggi := IPK_2(data, n)

	fmt.Println("\nMahasiswa dengan IPK tertinggi:")
	fmt.Printf("Nama    : %s\n", data[idxTertinggi].nama)
	fmt.Printf("NIM     : %s\n", data[idxTertinggi].nim)
	fmt.Printf("Kelas   : %s\n", data[idxTertinggi].kelas)
	fmt.Printf("Jurusan : %s\n", data[idxTertinggi].jurusan)
	fmt.Printf("IPK     : %.2f\n", data[idxTertinggi].ipk)
}
