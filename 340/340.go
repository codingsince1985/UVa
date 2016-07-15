// UVa 340 - Master-Mind Hints

package main

import (
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func match(s, g []int) (int, int) {
	ms, mg := make(map[int]int), make(map[int]int)
	strong, weak := 0, 0
	for i, v := range s {
		if g[i] == v {
			strong++
		} else {
			ms[v]++
			mg[g[i]]++
		}
	}

	for i, vi := range ms {
		vj := mg[i]
		weak += min(vi, vj)
	}
	return strong, weak
}

func main() {
	in, _ := os.Open("340.in")
	defer in.Close()
	out, _ := os.Create("340.out")
	defer out.Close()

	var n, kase int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		kase++
		fmt.Fprintf(out, "Game %d:\n", kase)
		s, g := make([]int, n), make([]int, n)
		for i := range s {
			fmt.Fscanf(in, "%d", &s[i])
		}
		for {
			stop := true
			for i := range g {
				fmt.Fscanf(in, "%d", &g[i])
				if g[i] != 0 {
					stop = false
				}
			}
			if stop {
				break
			}
			strong, weak := match(s, g)
			fmt.Fprintf(out, "    (%d,%d)\n", strong, weak)
		}
	}
}
