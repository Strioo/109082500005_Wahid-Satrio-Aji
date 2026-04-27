package main

import (
	"fmt"
	"math"
)

const MAX = 1000

type ArrayInt [MAX]int

func tampilArray(arr ArrayInt, n int) {
	for i := 0; i < n; i++ {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

func rataRata(arr ArrayInt, n int) float64 {
	var total int
	for i := 0; i < n; i++ {
		total += arr[i]
	}
	return float64(total) / float64(n)
}

func standarDeviasi(arr ArrayInt, n int) float64 {
	rata := rataRata(arr, n)
	var jumlah float64

	for i := 0; i < n; i++ {
		selisih := float64(arr[i]) - rata
		jumlah += selisih * selisih
	}

	return math.Sqrt(jumlah / float64(n))
}

func frekuensi(arr ArrayInt, n int, x int) int {
	var count int
	for i := 0; i < n; i++ {
		if arr[i] == x {
			count++
		}
	}
	return count
}

func hapusElemen(arr *ArrayInt, n *int, idx int) {
	for i := idx; i < *n-1; i++ {
		arr[i] = arr[i+1]
	}
	*n--
}

func main() {
	var arr ArrayInt
	var n int

	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

 	fmt.Println("Isi array:")
	tampilArray(arr, n)

 	fmt.Println("Elemen dengan indeks ganjil:")
	for i := 1; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

 	fmt.Println("Elemen dengan indeks genap:")
	for i := 0; i < n; i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()

 	var x int
	fmt.Scan(&x)

	fmt.Println("Elemen dengan indeks kelipatan", x, ":")
	for i := 0; i < n; i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()

 	var idx int
	fmt.Scan(&idx)

	hapusElemen(&arr, &n, idx)

	fmt.Println("Array setelah penghapusan:")
	tampilArray(arr, n)

 	fmt.Println("Rata-rata:", rataRata(arr, n))

 	fmt.Println("Standar deviasi:", standarDeviasi(arr, n))

 	var cari int
	fmt.Scan(&cari)

	fmt.Println("Frekuensi", cari, ":", frekuensi(arr, n, cari))
}