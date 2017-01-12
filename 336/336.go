// UVa 336 - A Node Too Far

package main

import (
	"fmt"
	"os"
)

func bfs(adj map[int][]int, n, t int) int {
	visited := make(map[int]bool)
	var queue [][2]int
	queue = append(queue, [2]int{n, t})
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		visited[curr[0]] = true
		adjs := adj[curr[0]]
		if curr[1] > 0 {
			for _, v := range adjs {
				if _, ok := visited[v]; !ok {
					queue = append(queue, [2]int{v, curr[1] - 1})
				}
			}
		}
	}
	return len(visited)
}

func main() {
	in, _ := os.Open("336.in")
	defer in.Close()
	out, _ := os.Create("336.out")
	defer out.Close()

	var nc, n1, n2, count int
	for {
		if fmt.Fscanf(in, "%d", &nc); nc == 0 {
			break
		}

		adj := make(map[int][]int)
		for ; nc > 0; nc-- {
			fmt.Fscanf(in, "%d%d", &n1, &n2)
			adj[n1] = append(adj[n1], n2)
			adj[n2] = append(adj[n2], n1)
		}

		var n, t int
		for {
			if fmt.Fscanf(in, "%d%d", &n, &t); n == 0 && t == 0 {
				fmt.Fscanln(in)
				break
			}
			visited := bfs(adj, n, t)
			count++
			fmt.Fprintf(out, "Case %d: %d nodes not reachable from node %d with TTL = %d.\n", count, len(adj)-visited, n, t)
		}
	}
}
