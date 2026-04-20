package main

import "fmt"

type suhu float64

func CelciusToReamur(celsius suhu) suhu {
	return celsius * 4 / 5
}

func CelciusToFahrenheit(celsius suhu) suhu {
	return (celsius * 9 / 5) + 32
}

func CelciusToKelvin(celsius suhu) suhu {
	return celsius + 273.15
}

func main() {
	var input float64
	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&input)
	celsius := suhu(input)

	fmt.Printf("\n%v celcius = %v reamur\n", input, float64(CelciusToReamur(celsius)))
	fmt.Printf("%v celcius = %v fahrenheit\n", input, float64(CelciusToFahrenheit(celsius)))
	fmt.Printf("%v celcius = %v kelvin\n", input, float64(CelciusToKelvin(celsius)))
}
