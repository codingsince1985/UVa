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

var links map[string][]string

func bfs(fm, to string) []byte {
	visited := make(map[string]bool)
	visited[fm] = true
	var queue []node
	queue = append(queue, node{fm, []byte{fm[0]}})

	for len(queue) != 0 {
		curr := queue[0]
		queue = queue[1:]
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
				queue = append(queue, node{v, newPath})
			}
		}
	}
	return nil
}

func main() {
	in, _ := os.Open("10009.in")
	defer in.Close()
	out, _ := os.Create("10009.out")
	defer out.Close()

	var kase, m, n int
	var fm, to string
	fmt.Fscanf(in, "%d\n", &kase)
	for kase > 0 {
		fmt.Fscanf(in, "\n%d%d", &m, &n)
		links = make(map[string][]string)
		for i := 0; i < m; i++ {
			fmt.Fscanf(in, "%s%s", &fm, &to)
			links[fm] = append(links[fm], to)
			links[to] = append(links[to], fm)
		}
		for i := 0; i < n; i++ {
			fmt.Fscanf(in, "%s%s", &fm, &to)
			path := bfs(fm, to)
			for _, v := range path {
				fmt.Fprintf(out, "%c", v)
			}
			fmt.Fprintln(out)
		}
		kase--
		if kase > 0 {
			fmt.Fprintln(out)
		}
	}
}
