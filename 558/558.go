// UVa 558 - Wormholes

package main

import (
	"fmt"
	"math"
	"os"
)

var n int

type edge struct{ from, to, time int }

func bellmanFord(edges []edge) bool {
	distance := make([]int, n)
	for i := range distance {
		distance[i] = math.MaxInt32
	}
	distance[0] = 0
	for i := 0; i < n-1; i++ {
		for _, e := range edges {
			if distance[e.to] < distance[e.from]+e.time {
				distance[e.to] = distance[e.from] + e.time
			}
		}
	}
	for _, e := range edges {
		if distance[e.from]+e.time < distance[e.to] {
			return true
		}
	}
	return false
}

func main() {
	in, _ := os.Open("558.in")
	defer in.Close()
	out, _ := os.Create("558.out")
	defer out.Close()

	var c, m, s1, s2, t int
	for fmt.Fscanf(in, "%d", &c); c > 0; c-- {
		fmt.Fscanf(in, "%d%d", &n, &m)
		edges := make([]edge, m)
		for i := 0; i < m; i++ {
			fmt.Fscanf(in, "%d%d%d", &s1, &s2, &t)
			edges[i] = edge{s1, s2, t}
		}
		if bellmanFord(edges) {
			fmt.Fprintln(out, "possible")
		} else {
			fmt.Fprintln(out, "not possible")
		}
	}
}
