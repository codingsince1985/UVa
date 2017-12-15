// UVa 11045 - My T-shirt suits me

package main

import (
	"fmt"
	"math"
	"os"
)

var sizeMap = map[string]int{"XXL": 1, "XL": 2, "L": 3, "M": 4, "S": 5, "XS": 6}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func edmondsKarp(s, t int, matrix [][]int) int {
	var sum int
	n := len(matrix)
	parent := make([]int, n)
	flow := make([][]int, n)
	for i := range flow {
		flow[i] = make([]int, n)
	}
	for {
		capacity := make([]int, n)
		capacity[s] = math.MaxInt32
		for queue := []int{s}; len(queue) > 0 && capacity[t] == 0; queue = queue[1:] {
			curr := queue[0]
			for i := 1; i < n; i++ {
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
	in, _ := os.Open("11045.in")
	defer in.Close()
	out, _ := os.Create("11045.out")
	defer out.Close()

	var kase, shirts, m int
	var s1, s2 string
	for fmt.Fscanf(in, "%d", &kase); kase > 0; kase-- {
		fmt.Fscanf(in, "%d%d", &shirts, &m)
		matrix := make([][]int, 12+m+2)
		for i := range matrix {
			matrix[i] = make([]int, 12+m+2)
		}
		for i := 1; i <= 6; i++ {
			matrix[0][i] = math.MaxInt32
			matrix[i][i+6] = shirts / 6
		}
		for i := 1; i <= m; i++ {
			fmt.Fscanf(in, "%s%s", &s1, &s2)
			matrix[sizeMap[s1]+6][12+i] = 1
			matrix[sizeMap[s2]+6][12+i] = 1
			matrix[12+i][12+m+1] = 1
		}
		if max := edmondsKarp(0, 12+m+1, matrix); max >= m {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
