package main

import (
	"fmt"
)

func main() {
	var suara [21]int
	var masuk, sah, vote int

	for {
		fmt.Scan(&vote)
		if vote == 0 {
			break
		}
		
		masuk++
		
		if vote >= 1 && vote <= 20 {
			sah++
			suara[vote]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", masuk)
	fmt.Printf("Suara sah: %d\n", sah)

	for i := 1; i <= 20; i++ {
		if suara[i] > 0 {
			fmt.Printf("%d: %d\n", i, suara[i])
		}
	}
}