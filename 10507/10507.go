// UVa 10507 - Waking up brain

package main

import (
	"fmt"
	"os"
)

type node struct{ area, year int }

func bfs(n int, matrix [][]bool, queue []node) int {
	var year int
	visited := make(map[int]bool)
	for ; len(queue) > 0; queue = queue[1:] {
		curr := queue[0]
		visited[curr.area] = true
		year = curr.year
		for i := range matrix {
			if !visited[i] {
				var count int
				for j := range matrix[i] {
					if visited[j] && matrix[i][j] {
						count++
					}
				}
				if count >= 3 {
					queue = append(queue, node{i, curr.year + 1})
				}
			}
		}
	}
	if len(visited) != n {
		return -1
	}
	return year
}

func initQueue(wakeup string, areaMap map[byte]int) []node {
	var queue []node
	for _, vw := range []byte(wakeup) {
		queue = append(queue, node{areaMap[vw], 0})
	}
	return queue
}

func main() {
	in, _ := os.Open("10507.in")
	defer in.Close()
	out, _ := os.Create("10507.out")
	defer out.Close()

	var n, m int
	var wakeup string
	var a1, a2 byte
	for {
		if _, err := fmt.Fscanf(in, "%d", &n); err != nil {
			break
		}
		areaMap := make(map[byte]int)
		matrix := make([][]bool, n)
		for i := range matrix {
			matrix[i] = make([]bool, n)
		}
		idx := 0
		for fmt.Fscanf(in, "%d\n%s", &m, &wakeup); m > 0; m-- {
			fmt.Fscanf(in, "%c%c\n", &a1, &a2)
			if _, ok := areaMap[a1]; !ok {
				areaMap[a1] = idx
				idx++
			}
			if _, ok := areaMap[a2]; !ok {
				areaMap[a2] = idx
				idx++
			}
			matrix[areaMap[a1]][areaMap[a2]], matrix[areaMap[a2]][areaMap[a1]] = true, true
		}
		if year := bfs(n, matrix, initQueue(wakeup, areaMap)); year != -1 {
			fmt.Fprintf(out, "WAKE UP IN, %d, YEARS\n", year)
		} else {
			fmt.Fprintln(out, "THIS BRAIN NEVER WAKES UP")
		}
		fmt.Fscanln(in)
	}
}
