// UVa 231 - Testing the CATCHER

package main

import (
	"fmt"
	"os"
)

func lds(h []int) int {
	max := 1
	l := len(h)
	m := make([]int, l)
	for i := range m {
		m[i] = 1
	}

	for i := 1; i < l; i++ {
		for j := 0; j < i; j++ {
			if h[i] < h[j] && m[i] <= m[j] {
				m[i] = m[j] + 1
				if max < m[i] {
					max = m[i]
				}
			}
		}
	}
	return max
}

func main() {
	in, _ := os.Open("231.in")
	defer in.Close()
	out, _ := os.Create("231.out")
	defer out.Close()

	var h []int
	var tmp int
	count := 0
	first := true
	for {
		if fmt.Fscanf(in, "%d", &tmp); first && tmp == -1 {
			break
		}
		if first {
			h = make([]int, 0)
			first = false
			count ++
		}
		if tmp == -1 {
			fmt.Fprintf(out, "Test #%d:\n", count)
			fmt.Fprintf(out, "  maximum possible interceptions: %d\n\n", lds(h))
			first = true
		}
		h = append(h, tmp)
	}
}
