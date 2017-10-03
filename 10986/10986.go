// UVa 10986 - Sending email

package main

import (
	"fmt"
	"math"
	"os"
)

type node struct{ server, latency int }

func spfa(n, s, t int, matrix [][]node) int {
	distance := make([]int, n)
	for i := range distance {
		distance[i] = math.MaxInt32
	}
	distance[s] = 0
	visited := map[int]bool{s: true}
	for queue := []int{s}; len(queue) > 0; queue = queue[1:] {
		curr := queue[0]
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

	var kase, m, s, t, s1, s2, l, n int
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
		if distance := spfa(n, s, t, matrix); distance == math.MaxInt32 {
			fmt.Fprintln(out, "unreachable")
		} else {
			fmt.Fprintln(out, distance)
		}
	}
}
