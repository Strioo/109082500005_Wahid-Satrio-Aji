package main

import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	if tujuan == "domestik" {
		if jumlahHari > 3 {
			return 3
		}
		return jumlahHari
	} else if tujuan == "mancanegara" {
		if jumlahHari > 8 {
			return 8
		}
		return jumlahHari
	}
	return 0
}

func biayaPerHari(jumlahMhs int) int {
	biayaMakan := 35000 * 2
	biayaPenginapan := 250000
	uangSaku := 300000
	totalPerMhs := biayaMakan + biayaPenginapan + uangSaku

	return jumlahMhs * totalPerMhs
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	hariDitanggung := tanggunganHari(lamaPerjalanan, tujuan)

	biayaHarianDomestik := biayaPerHari(jumlahMhs)

	if tujuan == "domestik" {
		*totalBiaya = float64(hariDitanggung * biayaHarianDomestik)
	} else if tujuan == "mancanegara" {
		*totalBiaya = float64(hariDitanggung) * float64(biayaHarianDomestik) * 1.5
	} else {
		*totalBiaya = 0
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	fmt.Print("masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)
	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	fmt.Println()
	fmt.Printf("Biaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biaya)
}
