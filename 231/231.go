// UVa 231 - Testing the CATCHER

package main

import (
	"fmt"
	"os"
)

func lds(h []int) int {
	l := len(h)
	m := make([]int, l)
	for i := range m {
		m[i] = 1
	}
	max := 1
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
	for {
		if fmt.Fscanf(in, "%d", &tmp); tmp == -1 && len(h) == 0 {
			break
		}
		if tmp == -1 {
			if count++; count > 1 {
				fmt.Fprintln(out)
			}
			fmt.Fprintf(out, "Test #%d:\n", count)
			fmt.Fprintf(out, "  maximum possible interceptions: %d\n", lds(h))
			h = nil
		} else {
			h = append(h, tmp)
		}
	}
}
