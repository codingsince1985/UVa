// UVa 10102 - The path in the colored field

package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func manhattanDistance(i, j int, grid [][]byte) int {
	minDistance := math.MaxInt32
	for x, vx := range grid {
		for y, vy := range vx {
			if vy == '3' {
				minDistance = min(minDistance, abs(i-x)+abs(j-y))
			}
		}
	}
	return minDistance
}

func solve(m int, grid [][]byte) int {
	var mins []int
	for x, vx := range grid {
		for y, vy := range vx {
			if vy == '1' {
				mins = append(mins, manhattanDistance(x, y, grid))
			}
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(mins)))
	return mins[0]
}

func main() {
	in, _ := os.Open("10102.in")
	defer in.Close()
	out, _ := os.Create("10102.out")
	defer out.Close()

	var m int
	var row string
	for {
		if _, err := fmt.Fscanf(in, "%d", &m); err != nil {
			break
		}
		grid := make([][]byte, m)
		for i := range grid {
			fmt.Fscanf(in, "%s", &row)
			grid[i] = []byte(row)
		}
		fmt.Fprintln(out, solve(m, grid))
	}
}
