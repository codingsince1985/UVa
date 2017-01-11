// UVa 10305 - Ordering Tasks

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var answer []string

func dfs(links map[int][]int, visited []bool, n int) {
	adjs := links[n]
	for _, v := range adjs {
		if !visited[v] {
			dfs(links, visited, v)
		}
	}
	visited[n] = true
	answer = append([]string{strconv.Itoa(n)}, answer...)
}

func topoSort(links map[int][]int, n int) {
	visited := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		if !visited[i] {
			dfs(links, visited, i)
		}
	}
}

func main() {
	in, _ := os.Open("10305.in")
	defer in.Close()
	out, _ := os.Create("10305.out")
	defer out.Close()

	var n, m, n1, n2 int
	for {
		if fmt.Fscanf(in, "%d%d", &n, &m); n == 0 && m == 0 {
			break
		}
		links := make(map[int][]int)
		for i := 0; i < m; i++ {
			fmt.Fscanf(in, "%d%d", &n1, &n2)
			links[n1] = append(links[n1], n2)
		}
		topoSort(links, n)
		fmt.Fprintln(out, strings.Join(answer, " "))
	}
}
