// UVa 10474 - Where is the Marble?

package main

import (
	"fmt"
	"os"
	"sort"
)

type marbles []int

func binarySearch(a []int, q, l, r int) int {
	if l+2 < r {
		m := (l + r) / 2
		if a[m] >= q {
			return binarySearch(a, q, l, m)
		}
		return binarySearch(a, q, m+1, r)
	}

	switch {
	case a[l] == q:
		return l
	case a[r] == q:
		return r
	default:
		return -1
	}
}

func main() {
	in, _ := os.Open("10474.in")
	defer in.Close()
	out, _ := os.Create("10474.out")
	defer out.Close()

	var N, Q, q, kase int
	var a marbles
	for {
		if fmt.Fscanf(in, "%d%d", &N, &Q); N == 0 && Q == 0 {
			break
		}
		a = make([]int, N)
		for i := range a {
			fmt.Fscanf(in, "%d", &a[i])
		}
		sort.Ints(a)

		kase++
		fmt.Fprintf(out, "CASE# %d:\n", kase)
		for i := 0; i < Q; i++ {
			fmt.Fscanf(in, "%d", &q)
			// manual binary search
			//			pos := binarySearch(a, q, 0, N-1) + 1
			//			switch pos {
			//			case 0:
			//				fmt.Fprintf(out, "%d not found\n", q)
			//			default:
			//				fmt.Fprintf(out, "%d found at %d\n", q, pos)
			//			}

			// built-in binary search
			pos := sort.Search(N, func(j int) bool { return a[j] >= q })
			switch {
			case pos < N && a[pos] == q:
				fmt.Fprintf(out, "%d found at %d\n", q, pos+1)
			default:
				fmt.Fprintf(out, "%d not found\n", q)
			}
		}
	}
}
