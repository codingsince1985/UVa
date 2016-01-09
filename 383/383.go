// UVa 383 - Shipping Routes

package main

import (
	"fmt"
	"os"
)

type node struct {
	n string
	l int
}

func bfs(legs map[string][]string, src, dest string) int {
	visited := make(map[string]bool)
	queue := make([]node, 0)
	visited[src] = true
	h, l := 0, 0
	queue = append(queue, node{src, 0}); l++
	for l != 0 {
		curr := queue[h]; l--; h++
		adjs := legs[curr.n]
		for _, adj := range adjs {
			if adj == dest {
				return curr.l + 1
			}
			if _, ok := visited[adj]; !ok {
				visited[adj] = true
				queue = append(queue, node{adj, curr.l + 1}); l++
			}
		}
	}
	return -1
}

func buildAdjacency(legs map[string][]string, wh1, wh2 string) {
	if _, ok := legs[wh1]; !ok {
		legs[wh1] = make([]string, 0)
	}
	legs[wh1] = append(legs[wh1], wh2)
}

func main() {
	in, _ := os.Open("383.in")
	defer in.Close()
	out, _ := os.Create("383.out")
	defer out.Close()

	var kase, m, n, p int
	var wh []string
	var legs map[string][]string
	var wh1, wh2 string
	fmt.Fscanf(in, "%d", &kase)

	fmt.Fprintf(out, "SHIPPING ROUTES OUTPUT\n")

	for i := 0; i < kase; i++ {
		fmt.Fscanf(in, "%d%d%d", &m, &n, &p)
		wh = make([]string, m)
		for j := 0; j < m; j++ {
			fmt.Fscanf(in, "%s", &wh[j])
		}

		legs = make(map[string][]string)
		for j := 0; j < n; j++ {
			fmt.Fscanf(in, "%s%s", &wh1, &wh2)
			buildAdjacency(legs, wh1, wh2)
			buildAdjacency(legs, wh2, wh1)
		}
		fmt.Fprintf(out, "\nDATA SET  %d\n\n", i + 1)

		var size int
		var src, dest string
		for j := 0; j < p; j++ {
			fmt.Fscanf(in, "%d%s%s", &size, &src, &dest)
			l := bfs(legs, src, dest)
			if l == -1 {
				fmt.Fprintf(out, "NO SHIPMENT POSSIBLE\n")
			} else {
				fmt.Fprintf(out, "$%d\n", l * 100 * size)
			}
		}
	}
	fmt.Fprintln(out, "\nEND OF OUTPUT")
}
