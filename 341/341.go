// UVa 341 - Non-Stop Travel

package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func spfa(n, start, end int, matrix [][]int) ([]string, int) {
	pre := make([]int, n)
	distance := make([]int, n)
	for i := range distance {
		distance[i] = math.MaxInt32
	}
	distance[start] = 0
	visited := map[int]bool{start: true}
	for queue := []int{start}; len(queue) > 0; {
		curr := queue[0]
		queue = queue[1:]
		for to, delay := range matrix[curr] {
			if delay != 0 && distance[to] > distance[curr]+delay {
				distance[to] = distance[curr] + delay
				pre[to] = curr
				if !visited[to] {
					queue = append(queue, to)
					visited[to] = true
				}
			}
		}
	}
	var path []string
	for i := end; i != start; i = pre[i] {
		path = append([]string{strconv.Itoa(i + 1)}, path...)
	}
	return append([]string{strconv.Itoa(start + 1)}, path...), distance[end]
}

func main() {
	in, _ := os.Open("341.in")
	defer in.Close()
	out, _ := os.Create("341.out")
	defer out.Close()

	var n, num, to, delay, start, end int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		matrix := make([][]int, n)
		for i := range matrix {
			matrix[i] = make([]int, n)
			for fmt.Fscanf(in, "%d", &num); num > 0; num-- {
				fmt.Fscanf(in, "%d%d", &to, &delay)
				matrix[i][to-1] = delay
			}
		}
		fmt.Fscanf(in, "%d%d", &start, &end)
		fmt.Fprintf(out, "Case %d: Path = ", kase)
		path, total := spfa(n, start-1, end-1, matrix)
		fmt.Fprintf(out, "%s; %d second delay\n", strings.Join(path, " "), total)
		fmt.Fscanln(in)
	}
}
