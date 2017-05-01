// UVa 11428 - Cubes

package main

import (
	"fmt"
	"os"
)

const max = 60

type pair struct{ x, y int }

var m = func() map[int]pair {
	m := make(map[int]pair)
	var n int
	for x := 2; x < max; x++ {
		for y := 1; y < x; y++ {
			n = x*x*x - y*y*y
			if s, ok := m[n]; !ok || y < s.y {
				m[n] = pair{x, y}
			}
		}
	}
	return m
}()

func main() {
	in, _ := os.Open("11428.in")
	defer in.Close()
	out, _ := os.Create("11428.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		if s, ok := m[n]; ok {
			fmt.Fprintln(out, s.x, s.y)
		} else {
			fmt.Fprintln(out, "No solution")
		}
	}
}
