package main

import "fmt"

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func cetakFibonacci(n int, i int) {
	if i > n {
		return
	}
	fmt.Print(fibonacci(i), " ")
	cetakFibonacci(n, i+1)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)
	cetakFibonacci(n, 0)
	fmt.Println()
}
