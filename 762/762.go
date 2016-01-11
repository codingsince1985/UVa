// UVa 762 - We Ship Cheap

package main

import (
	"fmt"
	"os"
)

func buildAdj(links map[string][]string, n1, n2 string) {
	if _, ok := links[n1]; !ok {
		links[n1] = make([]string, 0)
	}
	links[n1] = append(links[n1], n2)
}

type node struct {
	n    string
	path [][2]string
}

func bfs(links map[string][]string, n1, n2 string) [][2]string {
	queue := make([]node, 0)
	var h, l int
	visited := make(map[string]bool)

	visited[n1] = true
	queue = append(queue, node{n1, make([][2]string, 0)}); l++
	for l > 0 {
		curr := queue[h]; l--; h++
		if curr.n == n2 {
			return curr.path
		}
		next := links[curr.n]
		for _, v := range next {
			if _, ok := visited[v]; !ok {
				visited[v] = true
				path := curr.path;
				path = append(path, [2]string{curr.n, v})
				queue = append(queue, node{v, path}); l++
			}
		}
	}
	return nil
}

func main() {
	in, _ := os.Open("762.in")
	defer in.Close();
	out, _ := os.Create("762.out")
	defer out.Close()

	var ln int
	var n1, n2, src, dest string
	var links map[string][]string
	for {
		if _, err := fmt.Fscanf(in, "%d", &ln); err != nil {
			break
		}
		links = make(map[string][]string)
		for i := 0; i < ln; i++ {
			fmt.Fscanf(in, "%s%s", &n1, &n2)
			buildAdj(links, n1, n2)
			buildAdj(links, n2, n1)
		}
		fmt.Fscanf(in, "%s%s\n\n", &src, &dest)
		route := bfs(links, src, dest)
		if route == nil {
			fmt.Fprintln(out, "No route")
		} else {
			for _, v := range route {
				fmt.Fprintln(out, v[0], v[1])
			}
		}
		fmt.Fprintln(out)
	}
}
