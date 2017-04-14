// UVa 10600 - ACM Contest and Blackout

package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type (
	edge  struct{ s1, s2, cost int }
	edges []edge
)

func (e edges) Len() int           { return len(e) }
func (e edges) Less(i, j int) bool { return e[i].cost < e[j].cost }
func (e edges) Swap(i, j int)      { e[i], e[j] = e[j], e[i] }

func unionFind(x int, f []int) int {
	if f[x] != x {
		f[x] = unionFind(f[x], f)
	}
	return f[x]
}

func solve1(n int, e []edge) (int, []edge) {
	f := make([]int, n)
	for i := range f {
		f[i] = i
	}
	var selected []edge
	var min int
	for _, ve := range e {
		if f1, f2 := unionFind(ve.s1-1, f), unionFind(ve.s2-1, f); f1 != f2 {
			f[f1] = f2
			min += ve.cost
			if selected = append(selected, ve); len(selected) == n {
				break
			}
		}
	}
	return min, selected
}

func solve2(n int, e, selected []edge) int {
	min := math.MaxInt32
	for _, vs := range selected {
		f := make([]int, n)
		for i := range f {
			f[i] = i
		}
		cost, count := 0, 0
		for _, ve := range e {
			if vs != ve {
				if f1, f2 := unionFind(ve.s1-1, f), unionFind(ve.s2-1, f); f1 != f2 {
					f[f1] = f2
					cost += ve.cost
					if count++; count == n {
						break
					}
				}
			}
		}
		if cost < min {
			min = cost
		}
	}
	return min
}

func main() {
	in, _ := os.Open("10600.in")
	defer in.Close()
	out, _ := os.Create("10600.out")
	defer out.Close()

	var t, n, m int
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		fmt.Fscanf(in, "%d%d", &n, &m)
		e := make(edges, m)
		for i := range e {
			fmt.Fscanf(in, "%d%d%d", &e[i].s1, &e[i].s2, &e[i].cost)
		}
		sort.Sort(e)
		min, selected := solve1(n, e)
		fmt.Fprintln(out, min, solve2(n, e, selected))
	}
}
