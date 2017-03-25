// UVa 439 - Knight Moves

package main

import (
	"fmt"
	"os"
)

type cell struct{ x, y, steps int }

var visited [][]bool

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func bfs(from, to cell) int {
	visited = make([][]bool, 9)
	for i := range visited {
		visited[i] = make([]bool, 9)
	}

	queue := []cell{from}
	visited[from.x][from.y] = true
	for len(queue) > 0 {
		curr := queue[0]
		if curr.x == to.x && curr.y == to.y {
			return curr.steps
		}
		queue = queue[1:]
		for dx := -2; dx <= 2; dx++ {
			for dy := -2; dy <= 2; dy++ {
				newX, newY := curr.x+dx, curr.y+dy
				if abs(dx)+abs(dy) == 3 && newX >= 1 && newX <= 8 && newY >= 1 && newY <= 8 && !visited[newX][newY] {
					visited[newX][newY] = true
					queue = append(queue, cell{newX, newY, curr.steps + 1})
				}
			}
		}
	}
	return -1
}

func main() {
	in, _ := os.Open("439.in")
	defer in.Close()
	out, _ := os.Create("439.out")
	defer out.Close()

	var from, to string
	for {
		if _, err := fmt.Fscanf(in, "%s%s", &from, &to); err != nil {
			break
		}
		steps := bfs(cell{int(from[0]-'a') + 1, int(from[1] - '0'), 0}, cell{int(to[0]-'a') + 1, int(to[1] - '0'), 0})
		fmt.Fprintf(out, "To get from %s to %s takes %d knight moves.\n", from, to, steps)
	}
}
