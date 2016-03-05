// UVa 105 - The Skyline Problem

package main

import (
	"fmt"
	"os"
)

func main() {
	in, _ := os.Open("105.in")
	defer in.Close()
	out, _ := os.Create("105.out")
	defer out.Close()

	var l, h, r, maxr int
	var city [10001]int
	for {
		if _, err := fmt.Fscanf(in, "%d%d%d", &l, &h, &r); err != nil {
			break
		}
		for i := l; i < r; i++ {
			if city[i] < h {
				city[i] = h
			}
		}
		if r > maxr {
			maxr = r
		}
	}
	for i := 1; i < maxr; i++ {
		if city[i-1] != city[i] {
			fmt.Fprintf(out, "%d %d ", i, city[i])
		}
	}
	fmt.Fprintln(out, maxr, 0)
}
