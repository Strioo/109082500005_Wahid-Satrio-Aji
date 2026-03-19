package main

import "fmt"

const pi = 3.14

func hitungPersegi(sisi int64) {
	luas := sisi * sisi
	keliling := 4 * sisi
	fmt.Printf("Luas persegi : %d\n", luas)
	fmt.Printf("Keliling persegi : %d\n", keliling)
}

func hitungPersegiPanjang(panjang, lebar int64) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	fmt.Printf("Luas persegi panjang : %d\n", luas)
	fmt.Printf("Keliling persegi panjang : %d\n", keliling)
}

func hitungLingkaran(jarijari float64) {
	luas := pi * jarijari * jarijari
	keliling := 2 * pi * jarijari
	fmt.Printf("Luas lingkaran : %.6f\n", luas)
	fmt.Printf("Keliling lingkaran : %.4f\n", keliling)
}

func main() {
	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan : ")

	var pilihan int
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		var sisi int64
		fmt.Print("Masukkan sisi : ")
		fmt.Scan(&sisi)
		hitungPersegi(sisi)
	case 2:
		var panjang, lebar int64
		fmt.Print("Masukkan panjang : ")
		fmt.Scan(&panjang)
		fmt.Print("Masukkan lebar : ")
		fmt.Scan(&lebar)
		hitungPersegiPanjang(panjang, lebar)
	case 3:
		var jarijari float64
		fmt.Print("Masukkan jari-jari : ")
		fmt.Scan(&jarijari)
		hitungLingkaran(jarijari)
	default:
		fmt.Println("Pilihan tidak valid")
	}
}
