package main

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pemenang [100]string
	var jumlahPertandingan int

	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	for {
		fmt.Printf("Pertandingan %d: ", jumlahPertandingan+1)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			pemenang[jumlahPertandingan] = klubA
		} else if skorB > skorA {
			pemenang[jumlahPertandingan] = klubB
		} else {
			pemenang[jumlahPertandingan] = "Draw"
		}

		jumlahPertandingan++
	}

	for i := 0; i < jumlahPertandingan; i++ {
		fmt.Printf("Hasil %d: %s\n", i+1, pemenang[i])
	}
	fmt.Println("Pertandingan selesai")
}