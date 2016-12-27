// UVa 350 - Pseudo-Random Numbers

package main

import (
	"fmt"
	"os"
)

func findCycle(z, i, m, l int) int {
	cache := make(map[int]int)
	cnt := 1
	for cache[l] == 0 {
		cache[l] = cnt
		l = (z*l + i) % m
		cnt++
	}
	return cnt - cache[l]
}

func main() {
	in, _ := os.Open("350.in")
	defer in.Close()
	out, _ := os.Create("350.out")
	defer out.Close()

	var z, i, m, l int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d%d%d%d", &z, &i, &m, &l); m == 0 {
			break
		}
		fmt.Fprintf(out, "Case %d: %d\n", kase, findCycle(z, i, m, l))
	}
}
