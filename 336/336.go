// UVa 336 - A Node Too Far

package main

import (
	"fmt"
	"os"
)

func bfs(adj map[int][]int, n, t int) int {
	visited := make(map[int]bool)
	q, h, l := make([][2]int, 0), 0, 0
	q = append(q, [2]int{n, t}); l++
	for l != 0 {
		curr := q[h]; l--; h++
		visited[curr[0]] = true
		adjs := adj[curr[0]]
		if curr[1] > 0 {
			for _, v := range adjs {
				if _, ok := visited[v]; !ok {
					q = append(q, [2]int{v, curr[1] - 1}); l++
				}
			}
		}
	}
	return len(visited)
}

func buildAdjacency(adj map[int][]int, n1, n2 int) {
	if _, ok := adj[n1]; !ok {
		adj[n1] = make([]int, 0)
	}
	adj[n1] = append(adj[n1], n2)
}

func main() {
	in, _ := os.Open("336.in")
	defer in.Close()
	out, _ := os.Create("336.out")
	defer out.Close()

	var nc, n1, n2, count int
	var adj map[int][]int
	for {
		fmt.Fscanf(in, "%d", &nc)
		if nc == 0 {
			break
		}

		adj = make(map[int][]int)
		for i := 0; i < nc; i++ {
			fmt.Fscanf(in, "%d%d", &n1, &n2)
			buildAdjacency(adj, n1, n2)
			buildAdjacency(adj, n2, n1)
		}

		var n, t int
		for {
			fmt.Fscanf(in, "%d%d", &n, &t)
			if n == 0 && t == 0 {
				fmt.Fscanf(in, "\n")
				break
			}
			visited := bfs(adj, n, t)
			count ++
			fmt.Fprintf(out, "Case %d: %d nodes not reachable from node %d with TTL = %d.\n", count, len(adj) - visited, n, t)
		}
	}
}
