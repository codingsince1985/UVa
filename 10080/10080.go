// UVa 10080 - Gopher II

package main

import (
	"fmt"
	"math"
	"os"
)

var (
	n, m, s, v     int
	gophers, holes []point
	matrix         [][]bool
	matched        []int
)

type point struct{ x, y float64 }

func distance(p1, p2 point) float64 {
	return math.Sqrt((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y))
}

func dfs(gopher int, visited []bool) bool {
	for i, near := range matrix[gopher] {
		if !near {
			continue
		}
		if !visited[i] {
			visited[i] = true
			if matched[i] == -1 || dfs(matched[i], visited) {
				matched[i] = gopher
				return true
			}
		}
	}
	return false
}

func hungarian() int {
	matched = make([]int, n)
	for i := range matched {
		matched[i] = -1
	}
	safe := 0
	for i := range gophers {
		if dfs(i, make([]bool, n)) {
			safe++
		}
	}
	return safe
}

func solve() int {
	matrix = make([][]bool, n)
	for i := range matrix {
		matrix[i] = make([]bool, m)
		for j := range matrix[i] {
			if distance(gophers[i], holes[j]) <= float64(s*v) {
				matrix[i][j] = true
			}
		}
	}
	return hungarian()
}

func main() {
	in, _ := os.Open("10080.in")
	defer in.Close()
	out, _ := os.Create("10080.out")
	defer out.Close()

	var x, y float64
	for {
		if _, err := fmt.Fscanf(in, "%d%d%d%d", &n, &m, &s, &v); err != nil {
			break
		}
		gophers = make([]point, n)
		for i := range gophers {
			fmt.Fscanf(in, "%f%f", &x, &y)
			gophers[i] = point{x, y}
		}
		holes = make([]point, m)
		for i := range holes {
			fmt.Fscanf(in, "%f%f", &x, &y)
			holes[i] = point{x, y}
		}
		fmt.Fprintln(out, n-solve())
	}
}
