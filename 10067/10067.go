// UVa 10067 - Playing with Wheels

package main

import (
	"fmt"
	"os"
)

type config struct {
	n     [4]int
	steps int
}

var (
	forbidden  map[[4]int]bool
	target     config
	directions = [2]int{-1, 1}
)

func rotate(n [4]int, i, step int) [4]int {
	n[i] += step
	if n[i] == -1 {
		n[i] = 9
	} else if n[i] == 10 {
		n[i] = 0
	}
	return n
}

func bfs(initial config) int {
	visited := make(map[[4]int]bool)
	var queue []config
	visited[initial.n] = true
	queue = append(queue, initial)
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.n == target.n {
			return curr.steps
		}
		for i := range curr.n {
			for _, direction := range directions {
				next := rotate(curr.n, i, direction)
				if !visited[next] && !forbidden[next] {
					visited[next] = true
					queue = append(queue, config{next, curr.steps + 1})
				}
			}
		}
	}
	return -1
}

func main() {
	in, _ := os.Open("10067.in")
	defer in.Close()
	out, _ := os.Create("10067.out")
	defer out.Close()

	var kase, f int
	var t [4]int
	fmt.Fscanf(in, "%d", &kase)
	for kase > 0 {
		fmt.Fscanf(in, "%d%d%d%d", &t[0], &t[1], &t[2], &t[3])
		initial := config{t, 0}
		fmt.Fscanf(in, "%d%d%d%d", &t[0], &t[1], &t[2], &t[3])
		target = config{t, -1}
		forbidden = make(map[[4]int]bool)
		fmt.Fscanf(in, "%d", &f)
		for i := 0; i < f; i++ {
			fmt.Fscanf(in, "%d%d%d%d", &t[0], &t[1], &t[2], &t[3])
			forbidden[t] = true
		}
		fmt.Fprintln(out, bfs(initial))
		kase--
		if kase > 0 {
			fmt.Fscanln(in)
		}
	}
}
