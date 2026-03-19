package main

import "fmt"

func cetakDeret(n int64) {
	for {
		fmt.Print(n)
		if n == 1 {
			break
		}
		fmt.Print(" ")
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	fmt.Println()
}

func main() {
	var n int64

	fmt.Print("Masukkan bilangan : ")
	fmt.Scan(&n)

	cetakDeret(n)
}
