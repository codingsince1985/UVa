// UVa 10608 - Friends

package main

import (
	"fmt"
	"os"
)

func unionFind(x int, f []int) int {
	if f[x] != x {
		f[x] = unionFind(f[x], f)
	}
	return f[x]
}

func max(f []int) int {
	var m int
	rootMap := make(map[int]int)
	for i := range f {
		root := unionFind(i, f)
		rootMap[root]++
		if rootMap[root] > m {
			m = rootMap[root]
		}
	}
	return m
}

func main() {
	in, _ := os.Open("10608.in")
	defer in.Close()
	out, _ := os.Create("10608.out")
	defer out.Close()

	var kase, n, m, m1, m2 int
	fmt.Fscanf(in, "%d", &kase)
	for kase > 0 {
		kase--
		fmt.Fscanf(in, "%d%d", &n, &m)
		f := make([]int, n+1)
		for i := range f {
			f[i] = i
		}
		for m > 0 {
			m--
			fmt.Fscanf(in, "%d%d", &m1, &m2)
			f1, f2 := unionFind(m1, f), unionFind(m2, f)
			if f1 != f2 {
				f[f1] = f2
			}
		}
		fmt.Fprintln(out, max(f))
	}
}
