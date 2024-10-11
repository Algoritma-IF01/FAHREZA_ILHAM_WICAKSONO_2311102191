package main

import (
	"fmt"
)

func main() {
	var (
		nam float64
		nmk string
	)

	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scanln(&nam)

	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 && nam <= 80 {
		nmk = "AB"
	} else if nam > 65 && nam <= 72.5 {
		nmk = "B"
	} else if nam > 57.5 && nam <= 65 {
		nmk = "BC"
	} else if nam > 50 && nam <= 57.5 {
		nmk = "C"
	} else if nam > 40 && nam <= 50 {
		nmk = "D"
	} else {
		nmk = "E"
	}

	fmt.Println("Nilai mata kuliah: ", nmk)
}

/*
a. jika nam adalah 80.1 maka program akan error karena pada pogram tersebut variabel yang dimodifikasi adalah nam bukan nmk yang mana nam merupakan float dan nam dimodifikasi menjadi string dan jika pun yang dimodifikasi adalah nmk hasilnya akan mencetak D karena kurangnya sepisifikasi kondisi 
b. Pengkondisian kurang lengkap dan modifikasi variabel salah, seharusnya pengkondisian juga mengecek apakah nilai yang diinput kurang dari batas maksimum suatu NMK dan variabel yang seharusnya dimodifikasi adalah nmk bukan nma
*/
