package main

import "fmt"

func checkLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 == 0)
}

func main() {
	var year int

	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&year)

	fmt.Println("Kabisat:", checkLeapYear(year))
}
