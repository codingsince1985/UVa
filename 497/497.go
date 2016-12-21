// UVa 497 - Strategic Defense Initiative

package main

import (
	"fmt"
	"os"
	"sort"
)

type node struct{ num, idx int }

func find(n int, s [][]node) int {
	// for home-made binary search, see UVa 481
	return sort.Search(len(s), func(i int) bool {
		nodes := s[i]
		return nodes[len(nodes)-1].num >= 2
	})
}

func lis(h []int) [][]node {
	var s [][]node
	s = append(s, []node{{h[0], 0}})
	for i := 1; i < len(h); i++ {
		last := s[len(s)-1]
		if h[i] > last[len(last)-1].num {
			t := []node{{h[i], i}}
			s = append(s, t)
		} else {
			idx := find(h[i], s)
			nd := s[idx]
			tn := node{h[i], i}
			nd = append(nd, tn)
			s[idx] = nd
		}
	}
	return s
}

func output(out *os.File, s [][]node, l int) {
	var f []int
	k := len(s)
	fmt.Fprintf(out, "Max hits: %d\n", k)
	for i := k - 1; i >= 0; i-- {
		nodes := s[i]
		for j := len(nodes) - 1; j >= 0; j-- {
			if nodes[j].idx < l {
				f = append(f, nodes[j].num)
				k = nodes[j].idx
				break
			}
		}
	}
	for i := len(f) - 1; i >= 0; i-- {
		fmt.Fprintln(out, f[i])
	}
}

func main() {
	in, _ := os.Open("497.in")
	defer in.Close()
	out, _ := os.Create("497.out")
	defer out.Close()

	var n, tmp int
	for fmt.Fscanf(in, "%d\n\n", &n); n > 0; n-- {
		var h []int
		for {
			if _, err := fmt.Fscanf(in, "%d", &tmp); err != nil {
				break
			}
			h = append(h, tmp)
		}
		output(out, lis(h), len(h))
	}
}
