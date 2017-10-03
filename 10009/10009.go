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
	visited := map[string]bool{fm: true}
	for queue := []node{{fm, []byte{fm[0]}}}; len(queue) > 0; queue = queue[1:] {
		curr := queue[0]
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
	for fmt.Fscanf(in, "%d\n", &kase); kase > 0; kase-- {
		links = make(map[string][]string)
		fmt.Fscanf(in, "\n%d%d", &m, &n)
		for ; m > 0; m-- {
			fmt.Fscanf(in, "%s%s", &fm, &to)
			links[fm] = append(links[fm], to)
			links[to] = append(links[to], fm)
		}
		for ; n > 0; n-- {
			fmt.Fscanf(in, "%s%s", &fm, &to)
			fmt.Fprintln(out, string(bfs(fm, to)))
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
	}
}
