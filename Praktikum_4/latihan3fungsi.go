package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt((a-c)*(a-c) + (b-d)*(b-d))
}

func didalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {

	var (
		cx1, cx2, cy1, cy2, r1, r2, x, y float64
	)

	fmt.Print("Masukkan titik pusat dan radius lingkaran pertama: ")
	fmt.Scanln(&cx1, &cy1, &r1)

	fmt.Print("Masukkan titik pusat dan radius lingkaran kedua: ")
	fmt.Scanln(&cx2, &cy2, &r2)

	fmt.Print("Masukkan titik sembarang: ")
	fmt.Scanln(&x, &y)

	dalamLingkaran1 := didalam(cx1, cy1, r1, x, y)
	dalamLingkaran2 := didalam(cx2, cy2, r2, x, y)

	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
