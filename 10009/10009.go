// UVa 10009 - All Roads Lead Where?

package main

import (
	"fmt"
	"os"
)

type node struct {
	n    string
	path []byte
}

func bfs(links map[string][]string, fm, to string) []byte {
	visited := make(map[string]bool)
	queue, h, l := make([]node, 0), 0, 0

	visited[fm] = true
	queue = append(queue, node{fm, []byte{fm[0]}}); l++

	for l != 0 {
		curr := queue[h]; h++; l--
		adjs := links[curr.n]
		for _, v := range adjs {
			if _, ok := visited[v]; !ok {
				// need to clone slice, otherwise same underneath array (memory space) will be used
				newPath := make([]byte, len(curr.path))
				copy(newPath, curr.path)
				newPath = append(newPath, v[0])

				if v == to {
					return newPath
				}
				visited[v] = true
				queue = append(queue, node{v, newPath}); l++
			}
		}
	}
	return nil
}

func buildLink(links map[string][]string, fm, to string) {
	if _, ok := links[fm]; !ok {
		links[fm] = make([]string, 0)
	}
	links[fm] = append(links[fm], to)
}

func main() {
	in, _ := os.Open("10009.in")
	defer in.Close()
	out, _ := os.Create("10009.out")
	defer out.Close()

	var kase, m, n int
	var links map[string][]string
	var fm, to string
	fmt.Fscanf(in, "%d\n", &kase)
	for i := 0; i < kase; i++ {
		fmt.Fscanf(in, "\n%d%d", &m, &n)
		links = make(map[string][]string)
		for j := 0; j < m; j++ {
			fmt.Fscanf(in, "%s%s", &fm, &to)
			buildLink(links, fm, to)
			buildLink(links, to, fm)
		}
		for j := 0; j < n; j++ {
			fmt.Fscanf(in, "%s%s", &fm, &to)
			path := bfs(links, fm, to)
			for _, v := range path {
				fmt.Fprintf(out, "%c", v)
			}
			fmt.Fprintln(out)
		}
		fmt.Fprintln(out)
	}
}
