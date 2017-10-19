// UVa 314 - Robot

package main

import (
	"fmt"
	"os"
)

type node struct{ y, x, direction, duration int }

var (
	m, n         int
	turns        = []int{1, 3}
	steps        = []int{1, 2, 3}
	directionMap = map[string]int{"north": 0, "east": 1, "south": 2, "west": 3}
	directions   = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
)

func bfs(s, e node, store [][]bool) int {
	visited := make([][][4]bool, m+1)
	for i := range visited {
		visited[i] = make([][4]bool, n+1)
	}
	visited[s.y][s.x][s.direction] = true
	for queue := []node{s}; len(queue) > 0; queue = queue[1:] {
		curr := queue[0]
		if curr.y == e.y && curr.x == e.x {
			return curr.duration
		}
		for _, turn := range turns {
			if newDirection := (curr.direction + turn) % 4; !visited[curr.y][curr.x][newDirection] {
				visited[curr.y][curr.x][newDirection] = true
				queue = append(queue, node{curr.y, curr.x, newDirection, curr.duration + 1})
			}
		}
		for _, step := range steps {
			if newY, newX := curr.y+step*directions[curr.direction][0], curr.x+step*directions[curr.direction][1]; newY > 0 && newY <= m && newX > 0 && newX <= n && !visited[newY][newX][curr.direction] {
				if store[newY][newX] {
					break
				}
				visited[newY][newX][curr.direction] = true
				queue = append(queue, node{newY, newX, curr.direction, curr.duration + 1})
			}
		}
	}
	return -1
}

func main() {
	in, _ := os.Open("314.in")
	defer in.Close()
	out, _ := os.Create("314.out")
	defer out.Close()

	var tmp int
	var s, e node
	var direction string
	for {
		if fmt.Fscanf(in, "%d%d", &m, &n); m == 0 && n == 0 {
			break
		}
		store := make([][]bool, m+1)
		for i := range store {
			store[i] = make([]bool, n+1)
		}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if fmt.Fscanf(in, "%d", &tmp); tmp == 1 {
					store[i][j], store[i+1][j], store[i][j+1], store[i+1][j+1] = true, true, true, true
				}
			}
		}
		fmt.Fscanf(in, "%d%d", &s.y, &s.x)
		fmt.Fscanf(in, "%d%d", &e.y, &e.x)
		fmt.Fscanf(in, "%s", &direction)
		s.direction = directionMap[direction]
		fmt.Fprintln(out, bfs(s, e, store))
	}
}
