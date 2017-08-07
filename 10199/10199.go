// UVa 10199 - Tourist Guide

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	matrix               [][]bool
	n, t                 int
	isCutVertex, visited []bool
	disc, low, parent    []int
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func articulationPoint(curr, root int) {
	t++
	low[curr], disc[curr] = t, t
	visited[curr] = true
	children := 0
	for i := range matrix[curr] {
		if matrix[curr][i] {
			if !visited[i] {
				children++
				parent[i] = curr
				articulationPoint(i, root)
				low[curr] = min(low[curr], low[i])
				if curr == root && children > 1 || curr != root && low[i] >= disc[curr] {
					isCutVertex[curr] = true
				}
			} else {
				if parent[curr] != i {
					low[curr] = min(low[curr], disc[i])
				}
			}
		}
	}
}

func getCutVertex() []bool {
	isCutVertex = make([]bool, n)
	visited = make([]bool, n)
	disc = make([]int, n)
	low = make([]int, n)
	parent = make([]int, n)
	t = 0
	for i, v := range visited {
		if !v {
			articulationPoint(i, i)
		}
	}
	return isCutVertex
}

func solve(orderMap map[int]string) []string {
	var results []string
	for i, v := range getCutVertex() {
		if v {
			results = append(results, orderMap[i])
		}
	}
	sort.Strings(results)
	return results
}

func main() {
	in, _ := os.Open("10199.in")
	defer in.Close()
	out, _ := os.Create("10199.out")
	defer out.Close()

	var r int
	var c1, c2 string
	for kase := 1; ; kase++ {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		locationMap := make(map[string]int)
		orderMap := make(map[int]string)
		for i := 0; i < n; i++ {
			fmt.Fscanf(in, "%s", &c1)
			locationMap[c1] = i
			orderMap[i] = c1
		}
		matrix = make([][]bool, n)
		for i := range matrix {
			matrix[i] = make([]bool, n)
		}
		for fmt.Fscanf(in, "%d", &r); r > 0; r-- {
			fmt.Fscanf(in, "%s%s", &c1, &c2)
			matrix[locationMap[c1]][locationMap[c2]] = true
			matrix[locationMap[c2]][locationMap[c1]] = true
		}
		if kase > 1 {
			fmt.Fprintln(out)
		}
		results := solve(orderMap)
		fmt.Fprintf(out, "City map #%d: %d camera(s) found\n", kase, len(results))
		fmt.Fprintln(out, strings.Join(results, "\n"))
	}
}
