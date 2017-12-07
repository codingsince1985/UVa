// UVa 10503 - The dominoes solitaire

package main

import (
	"fmt"
	"os"
)

type piece struct{ l, r int }

var (
	n, right int
	pieces   []piece
	visited  map[int]bool
)

func backtracking(dots, curr int) bool {
	if curr == n {
		return dots == right
	}
	for i, p := range pieces {
		if !visited[i] {
			if p.l == dots {
				visited[i] = true
				if backtracking(p.r, curr+1) {
					return true
				}
				visited[i] = false
			}
			if p.r == dots {
				visited[i] = true
				if backtracking(p.l, curr+1) {
					return true
				}
				visited[i] = false
			}
		}
	}
	return false
}

func main() {
	in, _ := os.Open("10503.in")
	defer in.Close()
	out, _ := os.Create("10503.out")
	defer out.Close()

	var m, left, tmp int
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fscanf(in, "%d", &m)
		fmt.Fscanf(in, "%d%d", &tmp, &left)
		fmt.Fscanf(in, "%d%d", &right, &tmp)
		pieces = make([]piece, m)
		for i := range pieces {
			fmt.Fscanf(in, "%d%d", &pieces[i].l, &pieces[i].r)
		}
		visited = make(map[int]bool)
		if backtracking(left, 0) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
