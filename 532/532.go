// UVa 532 - Dungeon Master

package main

import (
	"fmt"
	"os"
)

type (
	point struct{ l, r, c int }
	node  struct {
		point
		n int
	}
)

var (
	delta   = []int{-1, 0, 1}
	l, r, c int
	cubes   [][][]byte
	visited [][][]bool
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func inRange(nl, nr, nc int) bool {
	if nl < 0 || nr < 0 || nc < 0 || nl >= l || nr >= r || nc >= c {
		return false
	}
	return true
}

func bfs(s, e point) int {
	visited[s.l][s.r][s.c] = true
	for queue := []node{{s, 0}}; len(queue) > 0; {
		curr := queue[0]
		queue = queue[1:]
		for _, dl := range delta {
			for _, dr := range delta {
				for _, dc := range delta {
					nl, nr, nc := curr.l+dl, curr.r+dr, curr.c+dc
					if abs(dl)+abs(dr)+abs(dc) == 1 && inRange(nl, nr, nc) {
						newPoint := point{nl, nr, nc}
						if newPoint == e {
							return curr.n + 1
						}
						if !visited[nl][nr][nc] && cubes[nl][nr][nc] == '.' {
							visited[nl][nr][nc] = true
							queue = append(queue, node{newPoint, curr.n + 1})
						}
					}
				}
			}
		}
	}
	return -1
}

func main() {
	in, _ := os.Open("532.in")
	defer in.Close()
	out, _ := os.Create("532.out")
	defer out.Close()

	var s, e point
	var line string
	for {
		if fmt.Fscanf(in, "%d%d%d", &l, &r, &c); l == 0 {
			break
		}

		cubes = make([][][]byte, l)
		visited = make([][][]bool, l)
		for i := range cubes {
			cubes[i] = make([][]byte, r)
			visited[i] = make([][]bool, r)
			for j := range cubes[i] {
				cubes[i][j] = make([]byte, c)
				visited[i][j] = make([]bool, c)
				fmt.Fscanf(in, "%s", &line)
				chars := []byte(line)
				for k := range cubes[i][j] {
					switch chars[k] {
					case 'S':
						s = point{i, j, k}
					case 'E':
						e = point{i, j, k}
					}
					cubes[i][j][k] = chars[k]
				}
			}
			fmt.Fscanln(in)
		}
		if minute := bfs(s, e); minute == -1 {
			fmt.Fprintln(out, "Trapped!")
		} else {
			fmt.Fprintf(out, "Escaped in %d minute(s).\n", minute)
		}
	}
}
