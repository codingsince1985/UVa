// UVa 200 - Rare Order

package main

import (
	"fmt"
	"io"
	"os"
)

var (
	order   []int
	visited map[int]bool
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func buildMatrix(chars map[byte]int, words []string) [][]bool {
	matrix := make([][]bool, len(chars))
	for i := range matrix {
		matrix[i] = make([]bool, len(chars))
	}
	for i := range words {
		if i+1 < len(words) {
			minLen := min(len(words[i]), len(words[i+1]))
			for j := 0; j < minLen; j++ {
				if words[i][j] != words[i+1][j] {
					matrix[chars[words[i][j]]][chars[words[i+1][j]]] = true
					break
				}
			}
		}
	}
	return matrix
}

func dfs(matrix [][]bool, route []int) bool {
	if len(route) == len(matrix) {
		order = make([]int, len(route))
		copy(order, route)
		return true
	}
	for i := range matrix {
		if !visited[i] && matrix[route[len(route)-1]][i] {
			visited[i] = true
			if dfs(matrix, append(route, i)) {
				return true
			}
			visited[i] = false
		}
	}
	return false
}

func output(out io.Writer, chars map[byte]int) {
	for _, v := range order {
		for k, idx := range chars {
			if v == idx {
				fmt.Fprintf(out, "%c", k)
				break
			}
		}
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("200.in")
	defer in.Close()
	out, _ := os.Create("200.out")
	defer out.Close()

	chars := make(map[byte]int)
	var words []string
	var word string
	for {
		if fmt.Fscanf(in, "%s", &word); word == "#" {
			break
		}
		for i := range word {
			if _, ok := chars[word[i]]; !ok {
				chars[word[i]] = len(chars)
			}
		}
		words = append(words, word)
	}
	matrix := buildMatrix(chars, words)

	first := 0
here:
	for first = range matrix {
		for j := range matrix[first] {
			if matrix[j][first] {
				continue here
			}
		}
		break
	}
	visited = map[int]bool{first: true}
	dfs(matrix, []int{first})
	output(out, chars)
}
