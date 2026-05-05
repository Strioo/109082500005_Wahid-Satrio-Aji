package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func cariNilaiPertama(data arrayMahasiswa, n int, nim string) int {
	for i := 0; i < n; i++ {
		if data[i].NIM == nim {
			return data[i].nilai
		}
	}
	return -1
}

func cariNilaiTerbesar(data arrayMahasiswa, n int, nim string) int {
	maxNilai := -1
	for i := 0; i < n; i++ {
		if data[i].NIM == nim {
			if data[i].nilai > maxNilai {
				maxNilai = data[i].nilai
			}
		}
	}
	return maxNilai
}

func main() {
	var data arrayMahasiswa
	var n int
	var nimCari string

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&data[i].NIM, &data[i].nama, &data[i].nilai)
	}

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimCari)

	nilaiPert := cariNilaiPertama(data, n, nimCari)
	nilaiMax := cariNilaiTerbesar(data, n, nimCari)

	fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", nimCari, nilaiPert)
	fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", nimCari, nilaiMax)
}