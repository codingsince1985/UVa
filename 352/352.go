// UVa 352 - The Seasonal War

package main

import (
	"fmt"
	"os"
)

func dfs(m [][]byte, cells [][]int, x, y, d, count int) {
	if x < 0 || y < 0 || x >= d || y >= d {
		return
	}
	if m[x][y] != '1' || cells[x][y] != 0 {
		return
	}
	cells[x][y] = count
	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			dfs(m, cells, x + dx, y + dy, d, count)
		}
	}
}

func floodFill(m [][]byte) int {
	count := 0
	d := len(m)
	cells := make([][]int, d)
	for i := range cells {
		cells[i] = make([]int, d)
	}
	for i := range cells {
		for j := range cells[i] {
			if m[i][j] == '1' && cells[i][j] == 0 {
				count ++
				dfs(m, cells, i, j, d, count)
			}
		}
	}
	return count
}

func main() {
	in, _ := os.Open("352.in")
	defer in.Close()
	out, _ := os.Create("352.out")
	defer out.Close()

	var d, count int
	var m [][]byte
	var line string
	for {
		if _, err := fmt.Fscanf(in, "%d", &d); err != nil {
			break
		}
		count ++
		m = make([][]byte, d)
		for i := range m {
			m[i] = make([]byte, d)
			fmt.Fscanf(in, "%s", &line)
			for j := range m[i] {
				m[i][j] = line[j]
			}
		}
		fmt.Fprintf(out, "Image number %d contains %d war eagles.\n", count, floodFill(m))
	}
}
