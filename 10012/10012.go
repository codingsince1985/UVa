// UVa 10012 - How Big Is It?

package main

import (
	"fmt"
	"math"
	"os"
)

var (
	m       int
	length  float64
	visited []bool
	radius  []float64
)

func currentPos(x, y, r float64) float64 { return x + math.Sqrt((y+r)*(y+r)-(y-r)*(y-r)) }

func calculate(order []int) {
	pos := make([]float64, m)
	pos[0] = radius[order[0]]
	maxLength := 2 * radius[order[0]]
	for i := 1; i < m; i++ {
		maxPos := radius[order[i]]
		for j := 0; j < i; j++ {
			maxPos = max(maxPos, currentPos(pos[j], radius[order[j]], radius[order[i]]))
		}
		pos[i] = maxPos
		maxLength = max(maxLength, pos[i]+radius[order[i]])
	}
	if maxLength < length {
		length = maxLength
	}
}

func dfs(level int, order []int) {
	if level == m {
		calculate(order)
		return
	}
	for i := range radius {
		if !visited[i] {
			visited[i] = true
			dfs(level+1, append(order, i))
			visited[i] = false
		}
	}
}

func main() {
	in, _ := os.Open("10012.in")
	defer in.Close()
	out, _ := os.Create("10012.out")
	defer out.Close()

	var n int
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%d", &m)
		radius = make([]float64, m)
		for i := range radius {
			fmt.Fscanf(in, "%f", &radius[i])
		}
		visited = make([]bool, m)
		length = math.MaxFloat64
		dfs(0, nil)
		fmt.Fprintf(out, "%.3f\n", length)
	}
}
