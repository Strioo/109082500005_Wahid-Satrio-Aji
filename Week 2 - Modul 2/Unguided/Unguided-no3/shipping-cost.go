package main

import "fmt"

func main() {
	var gram int

	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&gram)

	kg := gram / 1000
	sisa := gram % 1000

	biayaKg := kg * 10000

	var biayaSisa int
	if kg > 10 {
		biayaSisa = 0
	} else if sisa >= 500 {
		biayaSisa = sisa * 5
	} else {
		biayaSisa = sisa * 15
	}

	total := biayaKg + biayaSisa

	fmt.Println()
	fmt.Println("===== Detail Perhitungan =====")
	fmt.Printf("Detail berat  : %2d kg + %4d gram\n", kg, sisa)
	fmt.Printf("Detail biaya  : Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp %d\n", total)
}
