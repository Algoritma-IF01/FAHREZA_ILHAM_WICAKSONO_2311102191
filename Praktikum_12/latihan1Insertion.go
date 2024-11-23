package main

import "fmt"

type arrBilangan []int

func insertionSort(array *arrBilangan, n int) {
	var temp, i, j int

	for i = 1; i < n; i++ {
		temp = (*array)[i]
		j = i

		for j > 0 && temp < (*array)[j-1] {
			(*array)[j] = (*array)[j-1]
			j--
		}

		(*array)[j] = temp
	}
}

func cekJarak(array *arrBilangan, n int) {
	if n < 2 {
		fmt.Println("Data tidak cukup untuk menentukan jarak")
		return
	}

	var isEqual = true
	var distance = (*array)[1] - (*array)[0]

	for i := 1; i < n-1; i++ {
		if (*array)[i+1]-(*array)[i] != distance {
			isEqual = false
			break
		}
	}

	if isEqual {
		fmt.Printf("Data berjarak %d\n", distance)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}

func main() {
	var (
		array       arrBilangan
		n, bilangan int
	)

	for {
		fmt.Print("Masukkan bilangan: ")
		fmt.Scanln(&bilangan)

		if bilangan < 0 {
			break
		}

		array = append(array, bilangan)
		n++
	}

	insertionSort(&array, n)
	fmt.Println("Data setelah diurutkan:", array)

	cekJarak(&array, n)
}
