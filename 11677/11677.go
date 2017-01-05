// UVa 11677 - Alarm Clock

package main

import (
	"fmt"
	"os"
)

func solve(h1, m1, h2, m2 int) int {
	minutes := m2 - m1 + 60*(h2-h1)
	if minutes < 0 {
		minutes += 24 * 60
	}
	return minutes
}

func main() {
	in, _ := os.Open("11677.in")
	defer in.Close()
	out, _ := os.Create("11677.out")
	defer out.Close()

	var h1, m1, h2, m2 int
	for {
		if fmt.Fscanf(in, "%d%d%d%d", &h1, &m1, &h2, &m2); h1 == 0 && m1 == 0 && h2 == 0 && m2 == 0 {
			break
		}
		fmt.Fprintln(out, solve(h1, m1, h2, m2))
	}
}
