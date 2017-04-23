// UVa 808 - Bee Breeding

package main

import (
	"fmt"
	"os"
)

const max = 10000

type cell struct{ x, y int }

var grid = func() []cell {
	directions := [][2]int{{-1, 0}, {-1, 1}, {0, 1}, {1, 0}, {1, -1}, {0, -1}}
	grid := make([]cell, max+1)
	x, y, cnt := 0, 0, 1
	grid[cnt] = cell{x, y}
	for ring := 1; ; ring++ {
		y--
		if cnt++; cnt > max {
			break
		}
		grid[cnt] = cell{x, y}
		for i, direction := range directions {
			side := ring
			if i == 0 {
				side--
			}
			for ; side > 0; side-- {
				x += direction[0]
				y += direction[1]
				if cnt++; cnt > max {
					break
				}
				grid[cnt] = cell{x, y}
			}
		}
	}
	return grid
}()

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func distance(a, b int) int {
	xDiff := abs(grid[a].x - grid[b].x)
	yDiff := abs(grid[a].y - grid[b].y)
	if yDiff > xDiff {
		yDiff -= xDiff
	}
	return xDiff + yDiff
}

func main() {
	in, _ := os.Open("808.in")
	defer in.Close()
	out, _ := os.Create("808.out")
	defer out.Close()

	var a, b int
	for {
		if fmt.Fscanf(in, "%d%d", &a, &b); a == 0 && b == 0 {
			break
		}
		fmt.Fprintf(out, "The distance between cells %d and %d is %d.\n", a, b, distance(a, b))
	}
}
