package main

import (
	"bufio"
	"fmt"
	"os"
)

const nMax int = 127

type Tabel [nMax]rune

func isiArray(t *Tabel, n *int) {
	fmt.Print("Teks	: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	for _, ch := range text {
		if ch == '.' || *n >= nMax {
			break
		}

		t[*n] = ch
		(*n)++
	}
}

func cetakArray(t Tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}

	fmt.Println()
}

func balikanArray(t *Tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t Tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}

	return true
}

func main() {
	var t Tabel
	var n int

	isiArray(&t, &n)

	balikanArray(&t, n)
	fmt.Print("Reverse teks	: ")
	cetakArray(t, n)

	isPalindrom := palindrom(t, n)
	fmt.Println("Palindrom	?", isPalindrom)

}
