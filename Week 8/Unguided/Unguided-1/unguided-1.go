package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y float64
}

type lingkaran struct {
	cx, cy float64
	r      float64
}

func jarak(p, q titik) float64 {
	return math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
}

func didalam(c lingkaran, p titik) bool {
	pusat := titik{x: c.cx, y: c.cy}
	return jarak(pusat, p) <= c.r
}

func main() {
	var l1, l2 lingkaran
	var p titik

	fmt.Scan(&l1.cx, &l1.cy, &l1.r)
	fmt.Scan(&l2.cx, &l2.cy, &l2.r)
	fmt.Scan(&p.x, &p.y)

	inL1 := didalam(l1, p)
	inL2 := didalam(l2, p)

	if inL1 && inL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}