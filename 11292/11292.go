// UVa 11292 - Dragon of Loowater

package main

import (
	"fmt"
	"os"
	"sort"
)

func solve(dragons, knights []int) int {
	sort.Ints(knights)
	sort.Ints(dragons)
	var total int
	for _, dragon := range dragons {
		l := len(knights)
		if pos := sort.Search(l, func(j int) bool { return knights[j] >= dragon }); pos < l {
			total += knights[pos]
			knights = knights[pos+1:]
		} else {
			return -1
		}
	}
	return total
}

func main() {
	in, _ := os.Open("11292.in")
	defer in.Close()
	out, _ := os.Create("11292.out")
	defer out.Close()

	var n, m int
	for {
		if fmt.Fscanf(in, "%d%d", &n, &m); n == 0 && m == 0 {
			break
		}
		dragons := make([]int, n)
		for i := range dragons {
			fmt.Fscanf(in, "%dd", &dragons[i])
		}
		knights := make([]int, m)
		for i := range knights {
			fmt.Fscanf(in, "%d", &knights[i])
		}
		if cost := solve(dragons, knights); cost == -1 {
			fmt.Fprintln(out, "Loowater is doomed!")
		} else {
			fmt.Fprintln(out, cost)
		}
	}
}
