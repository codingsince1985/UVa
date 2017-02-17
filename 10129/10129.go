// UVa 10129 - Play on Words

package main

import (
	"fmt"
	"os"
)

func findStart(inDegree, outDegree map[int]int) int {
	start := -1
	for k, v := range inDegree {
		if v != outDegree[k] {
			if start == -1 && v == outDegree[k]+1 {
				start = k
			} else {
				return -1
			}
		}
	}
	return start
}

func dfs(current int, matrix [][]bool, visited map[int]bool) {
	visited[current] = true
	for i := range matrix[current] {
		if matrix[current][i] && !visited[i] {
			dfs(i, matrix, visited)
		}
	}
}

func solve(words []string) bool {
	inDegree, outDegree := make(map[int]int), make(map[int]int)
	matrix := make([][]bool, 26)
	for i := range matrix {
		matrix[i] = make([]bool, 26)
	}
	visited := make(map[int]bool)
	for _, word := range words {
		first, last := int(word[0]-'a'), int(word[len(word)-1]-'a')
		matrix[first][last], visited[first], visited[last] = true, false, false
		inDegree[first]++
		outDegree[last]++
	}
	if start := findStart(inDegree, outDegree); start != -1 {
		dfs(start, matrix, visited)
		for _, v := range visited {
			if !v {
				return false
			}
		}
		return true
	}
	return false
}

func main() {
	in, _ := os.Open("10129.in")
	defer in.Close()
	out, _ := os.Create("10129.out")
	defer out.Close()

	var t, n int
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		fmt.Fscanf(in, "%d", &n)
		words := make([]string, n)
		for i := range words {
			fmt.Fscanf(in, "%s", &words[i])
		}
		if solve(words) {
			fmt.Fprintln(out, "Ordering is possible.")
		} else {
			fmt.Fprintln(out, "The door cannot be opened.")
		}
	}
}
