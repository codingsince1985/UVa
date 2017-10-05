// UVa 10150 - Doublets

package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	words   []string
	wordMap map[string]int
	matrix  [][]bool
)

func doublet(w1, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}
	var count int
	for i := range w1 {
		if w1[i] != w2[i] {
			count++
		}
	}
	return count == 1
}

func buildGraph(words []string) {
	wordMap = make(map[string]int)
	l := len(words)
	matrix = make([][]bool, l)
	for i := range matrix {
		matrix[i] = make([]bool, l)
		wordMap[words[i]] = i
	}
	for i := 0; i < l-1; i++ {
		for j := i + 1; j < l; j++ {
			d := doublet(words[i], words[j])
			matrix[i][j], matrix[j][i] = d, d
		}
	}
}

func bfs(w1, w2 string) []string {
	visited := map[string]bool{w1: true}
	for queue := [][]string{{w1}}; len(queue) != 0; queue = queue[1:] {
		curr := queue[0]
		last := curr[len(curr)-1]
		if last == w2 {
			return curr
		}
		for i, isDoublet := range matrix[wordMap[last]] {
			if next := words[i]; isDoublet && !visited[next] {
				visited[next] = true
				queue = append(queue, append(curr, next))
			}
		}
	}
	return nil
}

func main() {
	in, _ := os.Open("10150.in")
	defer in.Close()
	out, _ := os.Create("10150.out")
	defer out.Close()

	var w1, w2 string
	first := true
	for {
		var w string
		if fmt.Fscanf(in, "%s", &w); w == "" {
			break
		}
		words = append(words, w)
	}
	buildGraph(words)
	for {
		if _, err := fmt.Fscanf(in, "%s%s", &w1, &w2); err != nil {
			break
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		if path := bfs(w1, w2); path == nil {
			fmt.Fprintln(out, "No solution.")
		} else {
			fmt.Fprintln(out, strings.Join(path, "\n"))
		}
	}
}
