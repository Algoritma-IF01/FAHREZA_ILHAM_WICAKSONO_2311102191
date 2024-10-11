package main

import (
	"fmt"
)

func main()  {
	var (
		beratParsel, detailKg, detailGr, biayaKg, biayaGr, totalBiaya int
	)

	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&beratParsel)

	detailKg = beratParsel / 1000
	detailGr = beratParsel % 1000

	biayaKg = detailKg * 10000

	if detailGr >= 500 {
		biayaGr = detailGr * 5 
	} else {
		biayaGr = detailGr * 15
	}

	if detailKg > 10 {
		totalBiaya = biayaKg
	} else {
		totalBiaya = biayaKg + biayaGr
	}

	fmt.Printf("Detail berat: %d kg + %d gr \n", detailKg, detailGr)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d \n", biayaKg, biayaGr )
	fmt.Println("Total biaya: Rp.", totalBiaya)
}