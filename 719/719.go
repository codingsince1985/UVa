// UVa 719 - Glass Beads

package main

import (
	"fmt"
	"os"
)

func solve(line string) int {
	i, j := 0, 1
	var x, y int
	for k, l := 0, len(line); i < l && j < l && k < l; {
		if x = i + k; x >= l {
			x -= l
		}
		if y = j + k; y >= l {
			y -= l
		}
		if line[x] == line[y] {
			k++
		} else {
			if line[x] > line[y] {
				i += k + 1
			} else {
				j += k + 1
			}
			if i == j {
				j++
			}
			k = 0
		}
	}
	return i + 1
}

func main() {
	in, _ := os.Open("719.in")
	defer in.Close()
	out, _ := os.Create("719.out")
	defer out.Close()

	var n int
	var line string
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%s", &line)
		fmt.Fprintln(out, solve(line))
	}
}
