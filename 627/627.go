// UVa 627 - The Net

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	id   int
	path []string
}

func bfs(s, e int, routes map[int][]int) []string {
	visited := map[int]bool{s: true}
	for queue := []node{{s, []string{strconv.Itoa(s)}}}; len(queue) > 0; queue = queue[1:] {
		curr := queue[0]
		if curr.id == e {
			return curr.path
		}
		for _, v := range routes[curr.id] {
			if !visited[v] {
				visited[v] = true
				queue = append(queue, node{v, append(curr.path, strconv.Itoa(v))})
			}
		}
	}
	return nil
}

func main() {
	in, _ := os.Open("627.in")
	defer in.Close()
	out, _ := os.Create("627.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var n, m, r1, r2, t, id int
	for s.Scan() {
		fmt.Sscanf(s.Text(), "%d", &n)
		routes := make(map[int][]int)
		for i := 0; i < n && s.Scan(); i++ {
			tokens := strings.Split(s.Text(), "-")
			fmt.Sscanf(tokens[0], "%d", &id)
			routes[id] = nil
			if len(tokens[1]) > 0 {
				for _, token := range strings.Split(tokens[1], ",") {
					fmt.Sscanf(token, "%d", &t)
					routes[id] = append(routes[id], t)
				}
			}
		}
		fmt.Fprintln(out, "-----")
		s.Scan()
		for fmt.Sscanf(s.Text(), "%d", &m); m > 0 && s.Scan(); m-- {
			fmt.Sscanf(s.Text(), "%d%d", &r1, &r2)
			if path := bfs(r1, r2, routes); len(path) == 0 {
				fmt.Fprintln(out, "connection impossible")
			} else {
				fmt.Fprintln(out, strings.Join(path, " "))
			}
		}
	}
}
