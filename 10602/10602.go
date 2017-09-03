// UVa 10602 - Editor Nottoobad

package main

import (
	"fmt"
	"os"
	"strings"
)

func compare(w1, w2 string) int {
	var prefix int
	for i := 0; i < len(w1) && i < len(w2); i++ {
		if w1[i] != w2[i] {
			break
		}
		prefix++
	}
	return prefix
}

func solve(n int, words []string) (int, []string) {
	curr := words[0]
	order := []string{curr}
	count := len(curr)
	for visited := map[int]bool{0: true}; len(visited) != n; {
		max := -1
		var idx int
		for i, word := range words {
			if !visited[i] {
				if prefix := compare(curr, word); prefix > max {
					max, idx = prefix, i
				}
			}
		}
		count += len(words[idx]) - max
		visited[idx] = true
		curr = words[idx]
		order = append(order, curr)
	}
	return count, order
}

func main() {
	in, _ := os.Open("10602.in")
	defer in.Close()
	out, _ := os.Create("10602.out")
	defer out.Close()

	var t, n int
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		fmt.Fscanf(in, "%d", &n)
		words := make([]string, n)
		for i := range words {
			fmt.Fscanf(in, "%s", &words[i])
		}
		count, order := solve(n, words)
		fmt.Fprintf(out, "%d\n%s\n", count, strings.Join(order, "\n"))
	}
}
