package main

import (
	"fmt"
)

func main() {
	var (
		b, count int
		isPrima  bool
	)

	fmt.Print("Bilangan: ")
	fmt.Scanln(&b)

	if b <= 0 {
		fmt.Println("Bilangan harus lebih dari atau sama dengan 1")
		return
	}

	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Printf("%d ", i)
			count++
		}
	}

	if count == 2 {
		isPrima = true
	} else {
		isPrima = false
	}

	fmt.Println("\nPrima: ", isPrima)
}
