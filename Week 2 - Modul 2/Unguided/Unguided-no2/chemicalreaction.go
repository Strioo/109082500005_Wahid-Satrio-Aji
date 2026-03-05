package main

import "fmt"

func checkColor(c1, c2, c3, c4 string) bool {
	return c1 == "merah" && c2 == "kuning" && c3 == "hijau" && c4 == "ungu"
}

func main() {
	var c1, c2, c3, c4 string
	result := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d : ", i)
		fmt.Scan(&c1, &c2, &c3, &c4)

		if !checkColor(c1, c2, c3, c4) {
			result = false
		}
	}

	fmt.Println("BERHASIL :", result)
}
