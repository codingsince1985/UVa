// UVa 696 - How Many Knights

package main

import (
	"fmt"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int { return a + b - min(a, b) }

func solve(m, n int) int {
	m, n = min(m, n), max(m, n)
	switch m {
	case 1:
		return n
	case 2:
		answer := (n / 4) * 4
		if (n % 4) > 2 {
			return answer + 4
		}
		return answer + (n%4)*2
	default:
		return (m*n + 1) / 2
	}
}

func main() {
	in, _ := os.Open("696.in")
	defer in.Close()
	out, _ := os.Create("696.out")
	defer out.Close()

	var m, n int
	for {
		if fmt.Fscanf(in, "%d%d", &m, &n); m == 0 && n == 0 {
			break
		}
		fmt.Fprintf(out, "%d knights may be placed on a %d row %d column board.\n", solve(m, n), m, n)
	}
}
