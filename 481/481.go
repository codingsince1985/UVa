// UVa 481 - What Goes Up

package main

import (
	"fmt"
	"os"
)

type node struct {
	num, idx int
}

func binarySearch(n int, l [][]node, s, e int) int {
	// find the 1st index that is not less than n
	// for build-in binary search, see UVa 497
	m := (s + e) / 2
	mid := l[m]
	if n == mid[len(mid) - 1].num {
		return m
	} else if s == e {
		return s
	} else if n > mid[len(mid) - 1].num {
		return binarySearch(n, l, m + 1, e)
	} else {
		return binarySearch(n, l, s, m)
	}
}

func lis(l []int) [][]node {
	var t []node
	t = make([]node, 1)
	t[0] = node{l[0], 0}
	var s [][]node
	s = append(s, t)
	for i := 1; i < len(l); i++ {
		last := s[len(s) - 1]
		if l[i] > last[len(last) - 1].num {
			t = []node{{l[i], i}}
			s = append(s, t)
		} else {
			idx := binarySearch(l[i], s, 0, len(s))
			t = s[idx]
			t = append(t, node{l[i], i})
			s[idx] = t
		}
	}
	return s
}

func output(out *os.File, l []int) {
	fmt.Fprintf(out, "%d\n-\n", len(l))
	for i := len(l) - 1; i >= 0; i-- {
		fmt.Fprintf(out, "%d\n", l[i])
	}
}

func main() {
	in, _ := os.Open("481.in")
	defer in.Close()
	out, _ := os.Create("481.out")
	defer out.Close()

	var tmp int
	var l []int
	for {
		if _, err := fmt.Fscanf(in, "%d", &tmp); err != nil {
			break
		}
		l = append(l, tmp)
	}
	s := lis(l)
	idx := len(l)
	var lst[]int
	for i := len(s) - 1; i >= 0; i -- {
		for j := len(s[i]) - 1; j >= 0; j-- {
			if s[i][j].idx < idx {
				lst = append(lst, s[i][j].num)
				idx = s[i][j].idx
				break
			}
		}
	}
	output(out, lst)
}
