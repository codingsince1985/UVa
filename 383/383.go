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
	var queue []node
	visited[src] = true
	queue = append(queue, node{src, 0})
	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
		adjs := legs[curr.n]
		for _, adj := range adjs {
			if adj == dest {
				return curr.l + 1
			}
			if _, ok := visited[adj]; !ok {
				visited[adj] = true
				queue = append(queue, node{adj, curr.l + 1})
			}
		}
	}
	return -1
}

func main() {
	in, _ := os.Open("383.in")
	defer in.Close()
	out, _ := os.Create("383.out")
	defer out.Close()

	var kase, m, n, p int
	var wh1, wh2 string
	fmt.Fscanf(in, "%d", &kase)

	fmt.Fprint(out, "SHIPPING ROUTES OUTPUT\n")
	for i := 1; i <= kase; i++ {
		fmt.Fscanf(in, "%d%d%d", &m, &n, &p)
		wh := make([]string, m)
		for j := range wh {
			fmt.Fscanf(in, "%s", &wh[j])
		}

		legs := make(map[string][]string)
		for ; n > 0; n-- {
			fmt.Fscanf(in, "%s%s", &wh1, &wh2)
			legs[wh1] = append(legs[wh1], wh2)
			legs[wh2] = append(legs[wh2], wh1)
		}
		fmt.Fprintf(out, "\nDATA SET  %d\n\n", i)

		var size int
		var src, dest string
		for ; p > 0; p-- {
			fmt.Fscanf(in, "%d%s%s", &size, &src, &dest)
			if l := bfs(legs, src, dest); l == -1 {
				fmt.Fprint(out, "NO SHIPMENT POSSIBLE\n")
			} else {
				fmt.Fprintf(out, "$%d\n", l*100*size)
			}
		}
	}
	fmt.Fprintln(out, "\nEND OF OUTPUT")
}
