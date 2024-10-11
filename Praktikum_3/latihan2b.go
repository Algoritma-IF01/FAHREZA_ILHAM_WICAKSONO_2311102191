package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var (
		pita       string
		bungaCount int
	)

	for {
		fmt.Printf("Bunga %d: ", bungaCount+1)
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == 	"selesai" {
			break
		}

		if pita == "" {
			pita += input
		} else {
			pita += " - " + input
		}

		bungaCount++
	}

	fmt.Println("Pita: ", pita)
	fmt.Println("Bunga: ", bungaCount)
}
