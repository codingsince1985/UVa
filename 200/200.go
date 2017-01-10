// UVa 200 - Rare Order

package main

import (
	"fmt"
	"os"
)

var (
	order []int
	done  bool
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

func dfs(matrix [][]bool, visited []bool, route []int) {
	if len(route) == len(matrix) {
		order = make([]int, len(route))
		copy(order, route)
		done = true
		return
	}
	for i := range matrix {
		if !visited[i] && matrix[route[len(route)-1]][i] {
			visited[i] = true
			dfs(matrix, visited, append(route, i))
			if done {
				return
			}
			visited[i] = false
		}
	}
}

func output(out *os.File, chars map[byte]int) {
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
			if matrix[j][first] == true {
				continue here
			}
		}
		break
	}

	visited := make([]bool, len(chars))
	visited[first] = true
	var route []int
	route = append(route, first)
	dfs(matrix, visited, route)
	output(out, chars)
}
