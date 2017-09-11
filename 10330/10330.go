// UVa 10330 - Power Transmission

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
		for queue := []int{s}; len(queue) > 0 && capacity[t] == 0; {
			curr := queue[0]
			queue = queue[1:]
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
	in, _ := os.Open("10330.in")
	defer in.Close()
	out, _ := os.Create("10330.out")
	defer out.Close()

	var num, capacity, m, n1, n2, b, d int
	for {
		if _, err := fmt.Fscanf(in, "%d", &num); err != nil {
			break
		}
		n = 2*num + 1
		matrix := make([][]int, n+1)
		for i := range matrix {
			matrix[i] = make([]int, n+1)
		}
		for i := 1; i <= num; i++ {
			fmt.Fscanf(in, "%d", &capacity)
			matrix[i][i+num] = capacity
		}
		for fmt.Fscanf(in, "%d", &m); m > 0; m-- {
			fmt.Fscanf(in, "%d%d%d", &n1, &n2, &capacity)
			matrix[n1+num][n2] = capacity
		}
		fmt.Fscanf(in, "%d%d", &b, &d)
		for ; b > 0; b-- {
			fmt.Fscanf(in, "%d", &n1)
			matrix[0][n1] = math.MaxInt32
		}
		for ; d > 0; d-- {
			fmt.Fscanf(in, "%d", &n1)
			matrix[n1+num][n] = math.MaxInt32
		}
		fmt.Fprintln(out, edmondsKarp(0, n, matrix))
	}
}
