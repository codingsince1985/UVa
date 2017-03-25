// UVa 571 - Jugs

package main

import (
	"fmt"
	"os"
	"strings"
)

type (
	status struct{ a, b int }
	node   struct {
		status
		steps []string
	}
)

func buildStatus(a, b int) status { return status{a, b} }

func bfs(a, b, n int) []string {
	visited := make(map[status]bool)
	queue := []node{{status{0, 0}, nil}}
	visited[status{0, 0}] = true
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.a == n || curr.b == n {
			return append(curr.steps, "success")
		}
		if s := buildStatus(a, curr.b); curr.a == 0 && !visited[s] {
			visited[s] = true
			queue = append(queue, node{s, append(curr.steps, "fill A")})
		}
		if s := buildStatus(curr.a, b); curr.b == 0 && !visited[s] {
			visited[s] = true
			queue = append(queue, node{s, append(curr.steps, "fill B")})
		}
		if curr.a < a && curr.b > 0 {
			if s := buildStatus(curr.a+curr.b, 0); curr.b < (a-curr.a) && !visited[s] {
				visited[s] = true
				queue = append(queue, node{s, append(curr.steps, "pour B A")})
			}
			if s := buildStatus(a, curr.b-(a-curr.a)); curr.b >= (a-curr.a) && !visited[s] {
				visited[s] = true
				queue = append(queue, node{s, append(curr.steps, "pour B A")})
			}
		}
		if curr.a > 0 && curr.b < b {
			if s := buildStatus(0, curr.a+curr.b); curr.a < (b-curr.b) && !visited[s] {
				visited[s] = true
				queue = append(queue, node{s, append(curr.steps, "pour A B")})
			}
			if s := buildStatus(curr.a-(b-curr.b), b); curr.a >= (b-curr.b) && !visited[s] {
				visited[s] = true
				queue = append(queue, node{s, append(curr.steps, "pour A B")})
			}
		}
		if s := buildStatus(0, curr.b); curr.a > 0 && !visited[s] {
			visited[s] = true
			queue = append(queue, node{s, append(curr.steps, "empty A")})
		}
		if s := buildStatus(curr.a, 0); curr.b > 0 && !visited[s] {
			visited[s] = true
			queue = append(queue, node{s, append(curr.steps, "empty B")})
		}
	}
	return nil
}

func main() {
	in, _ := os.Open("571.in")
	defer in.Close()
	out, _ := os.Create("571.out")
	defer out.Close()

	var a, b, n int
	for {
		if _, err := fmt.Fscanf(in, "%d%d%d", &a, &b, &n); err != nil {
			break
		}
		fmt.Fprintln(out, strings.Join(bfs(a, b, n), "\n"))
	}
}
