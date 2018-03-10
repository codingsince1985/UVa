// UVa 481 - What Goes Up

package main

import (
	"fmt"
	"io"
	"os"
)

type node struct{ num, idx int }

func binarySearch(n int, l [][]node, s, e int) int {
	// find the 1st index that is not less than n
	// for build-in binary search, see UVa 497
	m := (s + e) / 2
	switch mid := l[m]; {
	case n == mid[len(mid)-1].num:
		return m
	case s == e:
		return s
	case n > mid[len(mid)-1].num:
		return binarySearch(n, l, m+1, e)
	default:
		return binarySearch(n, l, s, m)
	}
}

func lis(l []int) [][]node {
	t := []node{{l[0], 0}}
	s := [][]node{t}
	for i := 1; i < len(l); i++ {
		if last := s[len(s)-1]; l[i] > last[len(last)-1].num {
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

func output(out io.Writer, l []int) {
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
	s, idx := lis(l), len(l)
	var lst []int
	for i := len(s) - 1; i >= 0; i-- {
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
