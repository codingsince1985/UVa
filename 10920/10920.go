// UVa 10920 - Spiral Tap

package main

import (
	"fmt"
	"os"
)

var directions = [][2]int64{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}

func findRing(size, p int64) int64 {
	for i := int64(0); i < size; i++ {
		if (i*2+1)*(i*2+1) >= p {
			return i
		}
	}
	return -1
}

func solve(size, p int64) (int64, int64) {
	ring := findRing(size, p)
	y := size/2 + 1 + ring
	x := y
	side := ring*2 + 1
	delta := side*side - p
	if delta == 0 {
		return y, x
	}
	for i, step := range []int64{side - 1, side - 1, side - 1, side - 2} {
		for ; step > 0; step-- {
			y, x = y+directions[i][0], x+directions[i][1]
			if delta--; delta == 0 {
				return y, x
			}
		}
	}
	return -1, -1
}

func main() {
	in, _ := os.Open("10920.in")
	defer in.Close()
	out, _ := os.Create("10920.out")
	defer out.Close()

	var size, p int64
	for {
		if fmt.Fscanf(in, "%d%d", &size, &p); size == 0 && p == 0 {
			break
		}
		y, x := solve(size, p)
		fmt.Fprintf(out, "Line = %d, column = %d.\n", y, x)
	}
}
