package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var char rune
	*n = 0
	for {
		fmt.Scanf("%c", &char)
		if char == '\n' || char == '\r' {
			continue
		}
		if char == '.' || *n >= NMAX {
			break
		}
		(*t)[*n] = char
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		temp := (*t)[i]
		(*t)[i] = (*t)[n-1-i]
		(*t)[n-1-i] = temp
	}
}

func palindrom(t tabel, n int) bool {
	var tCopy tabel
	for i := 0; i < n; i++ {
		tCopy[i] = t[i]
	}

	balikanArray(&tCopy, n)

	for i := 0; i < n; i++ {
		if t[i] != tCopy[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	fmt.Print("Teks: ")
	isiArray(&tab, &m)

	fmt.Print("Reverse teks: ")
	var tabRev tabel = tab
	balikanArray(&tabRev, m)
	cetakArray(tabRev, m)

	isPalindrom := palindrom(tab, m)
	fmt.Printf("Palindrom: %v\n", isPalindrom)
}
