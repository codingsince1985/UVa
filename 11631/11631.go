// UVa 11631 - Dark roads

package main

import (
	"fmt"
	"os"
	"sort"
)

type road struct{ c1, c2, length int }

func unionFind(x int, f []int) int {
	if f[x] != x {
		f[x] = unionFind(f[x], f)
	}
	return f[x]
}

func solve(roads []road, m, n int) int {
	sort.Slice(roads, func(i, j int) bool { return roads[i].length < roads[j].length })
	f := make([]int, n)
	for i := range f {
		f[i] = i
	}
	var total int
	for _, r := range roads {
		if f1, f2 := unionFind(r.c1, f), unionFind(r.c2, f); f1 != f2 {
			f[f1] = f2
			total += r.length
		}
	}
	return total
}

func main() {
	in, _ := os.Open("11631.in")
	defer in.Close()
	out, _ := os.Create("11631.out")
	defer out.Close()

	var m, n int
	for {
		if fmt.Fscanf(in, "%d%d", &m, &n); m == 0 && n == 0 {
			break
		}
		var total int
		roads := make([]road, n)
		for i := range roads {
			fmt.Fscanf(in, "%d%d%d", &roads[i].c1, &roads[i].c2, &roads[i].length)
			total += roads[i].length
		}
		fmt.Fprintln(out, total-solve(roads, m, n))
	}
}
