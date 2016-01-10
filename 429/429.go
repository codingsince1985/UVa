// UVa 429 - Word Transformation

package main

import (
	"fmt"
	"os"
)

func sameLenAndOneCharDiff(w1, w2 string) bool {
	if len(w1) != len(w2) {
		return false
	}
	count := 0
	for i := 0; i < len(w1); i++ {
		if (w1[i] != w2[i]) {
			count ++
		}
	}
	return count == 1
}

func buildLink(dict *map[string][]string, word string) {
	(*dict)[word] = make([]string, 0)
	for k, v := range (*dict) {
		if k == word {
			continue
		}
		if sameLenAndOneCharDiff(k, word) {
			(*dict)[k] = append(v, word)
			(*dict)[word] = append((*dict)[word], k)
		}
	}
}

type node struct {
	w string
	n int
}

func bfs(dict map[string][]string, fm, to string) int {
	visited := make(map[string]bool)
	queue := make([]node, 0)
	h, l := 0, 0

	visited[fm] = true
	queue = append(queue, node{fm, 0}); l++

	for l != 0 {
		curr := queue[h]; h++; l--
		adjs := dict[curr.w]
		for _, v := range adjs {
			if v == to {
				return curr.n + 1
			}
			if _, ok := visited[v]; !ok {
				visited[v] = true
				queue = append(queue, node{v, curr.n + 1}); l++
			}
		}
	}
	return -1
}

func main() {
	in, _ := os.Open("429.in")
	defer in.Close()
	out, _ := os.Create("429.out")
	defer out.Close()

	var n int
	var word, fm, to string
	var dict map[string][]string
	fmt.Fscanf(in, "%d\n\n", &n)
	for i := 0; i < n; i++ {
		dict = make(map[string][]string)
		for {
			fmt.Fscanf(in, "%s", &word)
			if word == "*" {
				break
			}
			buildLink(&dict, word)
		}
		for {
			if _, err := fmt.Fscanf(in, "%s%s\n", &fm, &to); err != nil {
				break
			}
			fmt.Fprintln(out, fm, to, bfs(dict, fm, to))
		}
	}
}
