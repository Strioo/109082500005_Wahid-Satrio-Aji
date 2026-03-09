package main

import "fmt"

func main() {
	var year int

	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&year)

	if (year%4 == 0 && year%100 != 0) || (year%400 ==0) {
		fmt.Println("Kabisat : true")
	} else {
		fmt.Println("Kabisat : false")
	}

}
