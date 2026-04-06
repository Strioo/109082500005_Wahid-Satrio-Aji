package main

import "fmt"

func printOdd(n int, current int) {
	if current > n {
		return
	}
	if current%2 != 0 {
		fmt.Print(current, " ")
	}
	printOdd(n, current+1)
}

func main() {
	var n int
	fmt.Print("Masukkan N: ")
	fmt.Scan(&n)
	printOdd(n, 1)
	fmt.Println()
}