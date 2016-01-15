// UVa 10305 - Ordering Tasks

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func dfs(links *map[int][]int, visited *[]bool, answer *[]string, n int) {
	adjs := (*links)[n]
	for _, v := range adjs {
		if !(*visited)[v] {
			dfs(links, visited, answer, v)
		}
	}
	(*visited)[n] = true
	(*answer) = append([]string{strconv.Itoa(n)}, (*answer)...)
}

func topoSort(links *map[int][]int, n int) [] string {
	visited := make([]bool, n + 1)
	answer := make([]string, 0)
	for i := 1; i <= n; i++ {
		if !visited[i] {
			dfs(links, &visited, &answer, i)
		}
	}
	return answer
}

func main() {
	in, _ := os.Open("10305.in")
	defer in.Close()
	out, _ := os.Create("10305.out")
	defer out.Close()

	var n, m int
	var n1, n2 int
	var links map[int][]int
	for {
		fmt.Fscanf(in, "%d%d", &n, &m)
		if n == 0 && m == 0 {
			break
		}
		links = make(map[int][]int)
		for i := 0; i < m; i++ {
			fmt.Fscanf(in, "%d%d", &n1, &n2)
			if _, ok := links[n1]; !ok {
				links[n1] = make([]int, 0)
			}
			links[n1] = append(links[n1], n2)
		}
		answer := topoSort(&links, n)
		fmt.Fprintln(out, strings.Join(answer, " "))
	}
}
