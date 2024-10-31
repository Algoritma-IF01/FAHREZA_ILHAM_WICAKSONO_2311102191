package main

import (
	"fmt"
)

func main() {
	var (
		klubA, klubB string
		skorA, skorB int
		pemenang     []string
	)

	pertandingan := 1

	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)

	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	for {
		fmt.Print("Petandingan ", pertandingan, ": ")
		fmt.Scanln(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang = append(pemenang, klubA)
		} else if skorA == skorB {
			pemenang = append(pemenang, "Draw")
		} else {
			pemenang = append(pemenang, klubB)
		}

		pertandingan++
	}

	for i := 0; i < len(pemenang); i++ {
		fmt.Println("Hasil", i+1, ":", pemenang[i])
	}

	fmt.Println("Pertandingan selesai")
}
