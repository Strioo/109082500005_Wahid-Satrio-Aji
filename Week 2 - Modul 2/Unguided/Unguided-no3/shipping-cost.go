package main

import "fmt"

func hitungBiaya(gram int) (kg, sisa, biayaKg, biayaSisa, total int) {
	kg = gram / 1000
	sisa = gram % 1000

	biayaKg = kg * 10000

	if kg > 10 {
		biayaSisa = 0
	} else if sisa >= 500 {
		biayaSisa = sisa * 5
	} else {
		biayaSisa = sisa * 15
	}

	total = biayaKg + biayaSisa
	return
}

func main() {
	var gram int

	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scan(&gram)

	kg, sisa, biayaKg, biayaSisa, total := hitungBiaya(gram)

	fmt.Println()
	fmt.Println("===== Detail Perhitungan =====")
	fmt.Printf("Detail berat  : %2d kg + %4d gram\n", kg, sisa)
	fmt.Printf("Detail biaya  : Rp. %d + Rp. %d\n", biayaKg, biayaSisa)
	fmt.Printf("Total biaya: Rp %d\n", total)
}
