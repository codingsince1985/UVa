// UVa 10986 - Sending email

package main

import (
	"fmt"
	"math"
	"os"
)

var n int

type node struct{ server, latency int }

func spfa(matrix [][]node, s, t int) int {
	visited := make([]bool, n)
	distance := make([]int, n)
	for i := range distance {
		distance[i] = math.MaxInt32
	}
	distance[s], visited[s] = 0, true
	queue := make([]int, 1)
	queue[0] = s
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, to := range matrix[curr] {
			if distance[to.server] > distance[curr]+to.latency {
				distance[to.server] = distance[curr] + to.latency
				if !visited[to.server] {
					queue = append(queue, to.server)
					visited[to.server] = true
				}
			}
		}
	}
	return distance[t]
}

func main() {
	in, _ := os.Open("10986.in")
	defer in.Close()
	out, _ := os.Create("10986.out")
	defer out.Close()

	var kase, m, s, t, s1, s2, l int
	fmt.Fscanf(in, "%d", &kase)
	for i := 1; i <= kase; i++ {
		fmt.Fscanf(in, "%d%d%d%d", &n, &m, &s, &t)
		matrix := make([][]node, n)
		for ; m > 0; m-- {
			fmt.Fscanf(in, "%d%d%d", &s1, &s2, &l)
			matrix[s1] = append(matrix[s1], node{s2, l})
			matrix[s2] = append(matrix[s2], node{s1, l})
		}
		fmt.Fprintf(out, "Case #%d: ", i)
		if distance := spfa(matrix, s, t); distance == math.MaxInt32 {
			fmt.Fprintln(out, "unreachable")
		} else {
			fmt.Fprintln(out, distance)
		}
	}
}
