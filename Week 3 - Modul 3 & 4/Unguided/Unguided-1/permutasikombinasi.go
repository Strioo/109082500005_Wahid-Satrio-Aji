package main

import "fmt"

func factorial(n int64) int64 {
	if n <= 1 {
		return 1
	}
	result := int64(1)
	for i := int64(2); i <= n; i++ {
		result *= i
	}
	return result
}

func permutasi(n, r int64) int64 {
	return factorial(n) / factorial(n-r)
}

func kombinasi(n, r int64) int64 {
	return permutasi(n, r) / factorial(r)
}

func main() {
	var a, b, c, d int64

	fmt.Print("masukkan nilai a : ")
	fmt.Scan(&a)
	fmt.Print("masukkan nilai b : ")
	fmt.Scan(&b)
	fmt.Print("masukkan nilai c : ")
	fmt.Scan(&c)
	fmt.Print("masukkan nilai d : ")
	fmt.Scan(&d)

	fmt.Printf("hasil permutasi %d dan %d adalah : %d\n", a, c, permutasi(a, c))
	fmt.Printf("hasil kombinasi %d dan %d adalah : %d\n", a, c, kombinasi(a, c))
	fmt.Printf("hasil permutasi %d dan %d adalah : %d\n", b, d, permutasi(b, d))
	fmt.Printf("hasil kombinasi %d dan %d adalah : %d\n", b, d, kombinasi(b, d))
}
