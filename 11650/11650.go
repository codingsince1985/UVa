// UVa 11650 - Mirror Clock

package main

import (
	"fmt"
	"os"
)

func solve(h, m int) (int, int) {
	if h < 12 {
		h = 12 - h
	}
	if m > 0 {
		m = 60 - m
		h--
	}
	return h, m
}

func main() {
	in, _ := os.Open("11650.in")
	defer in.Close()
	out, _ := os.Create("11650.out")
	defer out.Close()

	var t, h, m int
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		fmt.Fscanf(in, "%d:%d", &h, &m)
		h, m = solve(h, m)
		fmt.Fprintf(out, "%02d:%02d\n", h, m)
	}
}
