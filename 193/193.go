// UVa 193 - Graph Coloring

package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	links map[int][]int
	black []bool // 1-based
	answer []string // 1-based
	n, max int
)

func ok(curr int) bool {
	adjs := links[curr]
	for _, v := range adjs {
		if black[v] {
			return false
		}
	}
	return true
}

func dfs(curr int) {
	if curr > n {
		count := 0
		for i := 1; i <= n; i++ {
			if black[i] {
				count ++
			}
		}
		if count > max {
			max = count
			answer = nil
			for i := 1; i <= n; i++ {
				if black[i] {
					answer = append(answer, strconv.Itoa(i))
				}
			}
		}
		return
	}

	black[curr] = false
	dfs(curr + 1)
	black[curr] = true
	if ok(curr) {
		dfs(curr + 1)
	}
	black[curr] = false
}

func main() {
	in, _ := os.Open("193.in")
	defer in.Close()
	out, _ := os.Create("193.out")
	defer out.Close()

	var m, k, n1, n2 int
	fmt.Fscanf(in, "%d", &m)
	for i := 0; i < m; i++ {
		fmt.Fscanf(in, "%d%d", &n, &k)
		black = make([]bool, n + 1)
		links = make(map[int][]int)
		for j := 0; j < k; j++ {
			fmt.Fscanf(in, "%d%d", &n1, &n2)
			links[n1] = append(links[n1], n2)
			links[n2] = append(links[n2], n1)
		}
		dfs(1)
		fmt.Fprintln(out, max)
		fmt.Fprintln(out, strings.Join(answer, " "))
	}
}
