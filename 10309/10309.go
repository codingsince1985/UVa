// UVa 10309 - Turn the Lights Off

package main

import (
	"fmt"
	"math"
	"os"
)

var (
	minSteps int
	grid     [][]bool
)

func allOff() bool {
	for _, row := range grid {
		for _, cell := range row {
			if cell {
				return false
			}
		}
	}
	return true
}

func press(y, x int) {
	grid[y][x] = !grid[y][x]
	if y-1 >= 0 {
		grid[y-1][x] = !grid[y-1][x]
	}
	if x-1 >= 0 {
		grid[y][x-1] = !grid[y][x-1]
	}
	if y+1 < 10 {
		grid[y+1][x] = !grid[y+1][x]
	}
	if x+1 < 10 {
		grid[y][x+1] = !grid[y][x+1]
	}
}

func dfs(y, x, steps int) {
	if steps >= minSteps {
		return
	}
	if x == 10 {
		if y++; y == 10 {
			if allOff() && steps < minSteps {
				minSteps = steps
			}
			return
		}
		x = 0
	}
	if y == 0 {
		dfs(y, x+1, steps)
		press(y, x)
		dfs(y, x+1, steps+1)
		press(y, x)
	} else {
		if grid[y-1][x] {
			press(y, x)
			dfs(y, x+1, steps+1)
			press(y, x)
		} else {
			dfs(y, x+1, steps)
		}
	}
}

func main() {
	in, _ := os.Open("10309.in")
	defer in.Close()
	out, _ := os.Create("10309.out")
	defer out.Close()

	var name, line string
	for {
		if fmt.Fscanf(in, "%s", &name); name == "end" {
			break
		}
		grid = make([][]bool, 10)
		for i := range grid {
			grid[i] = make([]bool, 10)
			fmt.Fscanf(in, "%s", &line)
			for j := range line {
				if line[j] == 'O' {
					grid[i][j] = true
				}
			}
		}
		minSteps = math.MaxInt32
		dfs(0, 0, 0)
		fmt.Fprintf(out, "%s %d\n", name, minSteps)
	}
}
