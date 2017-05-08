// UVa 208 - Firetruck

package main

import (
	"fmt"
	"os"
)

const max = 21

var (
	out            *os.File
	matrix         [][]bool
	total, n, fire int
	visited        []bool
	streetCorners  []int
)

func unionFind(x int, f []int) int {
	if f[x] != x {
		f[x] = unionFind(f[x], f)
	}
	return f[x]
}

func clear() {
	matrix = make([][]bool, max)
	for i := range matrix {
		matrix[i] = make([]bool, max)
	}
}

func output(path []int) {
	for _, s := range path {
		fmt.Fprintf(out, " %d", s)
	}
	fmt.Fprintln(out)
}

func dfs(curr int, path []int) {
	if curr == fire {
		output(path)
		total++
		return
	}
	for _, s := range streetCorners {
		if !visited[s] && matrix[curr][s] {
			visited[s] = true
			dfs(s, append(path, s))
			visited[s] = false
		}
	}
}

func solve() {
	matrix = matrix[:n+1]
	for i := range matrix {
		matrix[i] = matrix[i][:n+1]
	}
	f := make([]int, n+1)
	for i := range f {
		f[i] = i
	}
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if matrix[i][j] {
				if s1, s2 := unionFind(i, f), unionFind(j, f); s1 != s2 {
					f[s1] = s2
				}
			}
		}
	}
	streetCorners = nil
	for i := 1; i <= n; i++ {
		if unionFind(i, f) == unionFind(fire, f) {
			streetCorners = append(streetCorners, i)
		}
	}
	visited = make([]bool, n+1)
	visited[1] = true
	total = 0
	dfs(1, []int{1})
	fmt.Fprintf(out, "There are %d routes from the firestation to streetcorner %d.\n", total, fire)
}

func main() {
	in, _ := os.Open("208.in")
	defer in.Close()
	out, _ = os.Create("208.out")
	defer out.Close()

	var from, to int
	for kase := 1; ; kase++ {
		if _, err := fmt.Fscanf(in, "%d", &fire); err != nil {
			break
		}
		clear()
		n = 0
		for {
			if fmt.Fscanf(in, "%d%d", &from, &to); from == 0 && to == 0 {
				break
			}
			matrix[from][to], matrix[to][from] = true, true
			if from > n {
				n = from
			}
			if to > n {
				n = to
			}
		}
		fmt.Fprintf(out, "CASE %d:\n", kase)
		solve()
	}
}
