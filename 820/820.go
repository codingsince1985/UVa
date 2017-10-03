// UVa 820 - Internet Bandwidth

package main

import (
	"fmt"
	"math"
	"os"
)

var n int

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func edmondsKarp(s, t int, matrix [][]int) int {
	var sum int
	parent := make([]int, n+1)
	flow := make([][]int, n+1)
	for i := range flow {
		flow[i] = make([]int, n+1)
	}
	for {
		capacity := make([]int, n+1)
		capacity[s] = math.MaxInt32
		for queue := []int{s}; len(queue) > 0 && capacity[t] == 0; queue = queue[1:] {
			curr := queue[0]
			for i := 1; i <= n; i++ {
				if capacity[i] == 0 && matrix[curr][i] > flow[curr][i] {
					queue = append(queue, i)
					parent[i] = curr
					capacity[i] = min(capacity[curr], matrix[curr][i]-flow[curr][i])
				}
			}
		}
		if capacity[t] == 0 {
			break
		}
		for curr := t; curr != s; curr = parent[curr] {
			flow[parent[curr]][curr] += capacity[t]
			flow[curr][parent[curr]] -= capacity[t]
		}
		sum += capacity[t]
	}
	return sum
}

func main() {
	in, _ := os.Open("820.in")
	defer in.Close()
	out, _ := os.Create("820.out")
	defer out.Close()

	var s, t, c, n1, n2, bw int
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		fmt.Fscanf(in, "%d%d%d", &s, &t, &c)
		matrix := make([][]int, n+1)
		for i := range matrix {
			matrix[i] = make([]int, n+1)
		}
		for ; c > 0; c-- {
			fmt.Fscanf(in, "%d%d%d", &n1, &n2, &bw)
			matrix[n1][n2] += bw
			matrix[n2][n1] = matrix[n1][n2]
		}
		fmt.Fprintf(out, "Network %d\n", kase)
		fmt.Fprintf(out, "The bandwidth is %d.\n\n", edmondsKarp(s, t, matrix))
	}
}
