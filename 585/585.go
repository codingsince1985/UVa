// UVa 585 - Triangles

package main

import (
	"fmt"
	"os"
)

var (
	n     int
	paper [][]byte
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int { return a + b - min(a, b) }

func dfsUp(i, j, level int) int {
	if i > 0 && paper[i-1][j] == '-' && paper[i-1][j+1] == '-' && paper[i-1][j+2] == '-' {
		return min(dfsUp(i-1, j, level+1), min(dfsUp(i-1, j+1, level+1), dfsUp(i-1, j+2, level+1)))
	}
	return level * (2 * level) / 2
}

func dfsDown(i, j, level int) int {
	if i+1 < n && j < len(paper[i])-1 && j >= 2 && paper[i+1][j-2] == '-' && paper[i+1][j-1] == '-' && paper[i+1][j] == '-' {
		return min(dfsDown(i+1, j-2, level+1), min(dfsDown(i+1, j-1, level+1), dfsDown(i+1, j, level+1)))
	}
	return level * (2 * level) / 2
}

func solve() int {
	var largest int
	for i, row := range paper {
		for j, cell := range row {
			if cell == '-' {
				largest = max(largest, max(dfsUp(i, j, 1), dfsDown(i, j, 1)))
			}
		}
	}
	return largest
}

func main() {
	in, _ := os.Open("585.in")
	defer in.Close()
	out, _ := os.Create("585.out")
	defer out.Close()

	var line string
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		paper = make([][]byte, n)
		for i := range paper {
			fmt.Fscanf(in, "%s", &line)
			paper[i] = []byte(line)
		}
		fmt.Fprintf(out, "Triangle #%d\n", kase)
		fmt.Fprintf(out, "The largest triangle area is %d.\n\n", solve())
	}
}
