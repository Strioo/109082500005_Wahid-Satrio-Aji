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

	var ketua, wakil int
	var maxKetua, maxWakil int

	for i := 1; i <= 20; i++ {
		if suara[i] > maxKetua {
			wakil = ketua
			maxWakil = maxKetua
			
			ketua = i
			maxKetua = suara[i]
		} else if suara[i] > maxWakil {
			wakil = i
			maxWakil = suara[i]
		} else if suara[i] == maxKetua && ketua != 0 && wakil == 0 {
             wakil = i
             maxWakil = suara[i]
        } else if suara[i] == maxKetua && suara[i] > maxWakil {
             wakil = i
             maxWakil = suara[i]
        }
	}

	fmt.Printf("Ketua RT: %d\n", ketua)
	fmt.Printf("Wakil ketua: %d\n", wakil)
}