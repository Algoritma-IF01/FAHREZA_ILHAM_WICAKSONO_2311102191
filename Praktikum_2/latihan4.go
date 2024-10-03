package main

import (
	"fmt"
	// "strconv"
)

func main() {
	var (
		celcius, fahrenheit, kelvin, reamur float32
	)

	fmt.Print("Masukkan suhu celcius: ")
	fmt.Scanln(&celcius)

	fahrenheit = (9.0 / 5.0) * celcius + 32
	reamur = (4.0 / 5.0) * celcius
	kelvin = celcius + 273.15

	fmt.Printf("Derajat Reamur %.2f \n", reamur)
	fmt.Printf("Derajat Fahrenheit %.2f \n", fahrenheit)
	fmt.Printf("Derajat Kelvin %.2f \n", kelvin)
}
