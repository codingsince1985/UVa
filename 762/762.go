// UVa 762 - We Ship Cheap

package main

import (
	"fmt"
	"os"
)

type node struct {
	n    string
	path [][2]string
}

func bfs(links map[string][]string, n1, n2 string) [][2]string {
	queue := []node{{n1, nil}}
	visited := map[string]bool{n1: true}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.n == n2 {
			return curr.path
		}
		next := links[curr.n]
		for _, v := range next {
			if _, ok := visited[v]; !ok {
				visited[v] = true
				path := curr.path
				path = append(path, [2]string{curr.n, v})
				queue = append(queue, node{v, path})
			}
		}
	}
	return nil
}

func main() {
	in, _ := os.Open("762.in")
	defer in.Close()
	out, _ := os.Create("762.out")
	defer out.Close()

	var ln int
	var n1, n2, src, dest string
	first := true
	for {
		if _, err := fmt.Fscanf(in, "%d", &ln); err != nil {
			break
		}
		links := make(map[string][]string)
		for ; ln > 0; ln-- {
			fmt.Fscanf(in, "%s%s", &n1, &n2)
			links[n1] = append(links[n1], n2)
			links[n2] = append(links[n2], n1)
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		fmt.Fscanf(in, "%s%s\n\n", &src, &dest)
		if route := bfs(links, src, dest); route == nil {
			fmt.Fprintln(out, "No route")
		} else {
			for _, v := range route {
				fmt.Fprintln(out, v[0], v[1])
			}
		}
	}
}
