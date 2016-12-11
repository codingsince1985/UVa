// UVa 10125 - Sumsets

package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type (
	node  struct{ n, n1, n2 int }
	nodes []node
)

func (n nodes) Len() int { return len(n) }

func (n nodes) Less(i, j int) bool { return n[i].n < n[j].n }

func (n nodes) Swap(i, j int) { n[i], n[j] = n[j], n[i] }

func preCalculate(num []int, f func(int, int) int) nodes {
	var pre nodes
	for i, vi := range num {
		for j, vj := range num {
			if i != j {
				pre = append(pre, node{f(vi, vj), vi, vj})
			}
		}
	}
	sort.Sort(pre)
	return pre
}

func binarySearch(n int, sub nodes) (int, bool) {
	l, r := 0, len(sub)-1
	for l+1 < r {
		mid := (l + r) / 2
		switch {
		case sub[mid].n == n:
			return sub[mid].n1, true
		case sub[mid].n > n:
			r = mid
		default:
			l = mid
		}
	}
	switch {
	case sub[l].n == n:
		return sub[l].n1, true
	case sub[r].n == n:
		return sub[r].n1, true
	default:
		return 0, false
	}
}

func solve(num []int) (int, bool) {
	sum := preCalculate(num, func(a, b int) int { return a + b })
	sub := preCalculate(num, func(a, b int) int { return a - b })
	max, found := math.MinInt32, false
	for _, vi := range sum {
		if d, ok := binarySearch(vi.n, sub); ok {
			if d > max {
				max, found = d, true
			}
		}
	}
	return max, found
}

func main() {
	in, _ := os.Open("10125.in")
	defer in.Close()
	out, _ := os.Create("10125.out")
	defer out.Close()

	var n int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		num := make([]int, n)
		for i := range num {
			fmt.Fscanf(in, "%d", &num[i])
		}
		if d, found := solve(num); found {
			fmt.Fprintln(out, d)
		} else {
			fmt.Fprintln(out, "no solution")
		}
	}
}
