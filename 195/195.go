// UVa 195 - Anagram

package main

import (
	"fmt"
	"os"
	"sort"
)

var (
	chars []byte
	out   *os.File
)

func isLower(a byte) bool { return a >= 'a' && a <= 'z' }

func isUpper(a byte) bool { return !isLower(a) }

func toUpper(a byte) byte { return a - 'a' + 'A' }

func dfs(i int, word []byte, visited []bool) {
	if i == len(chars) {
		fmt.Fprintln(out, string(word))
	} else {
		for j := range chars {
			if !visited[j] {
				if j > 0 && chars[j] == chars[j-1] && !visited[j-1] {
					continue
				}
				visited[j] = true
				word[i] = chars[j]
				dfs(i+1, word, visited)
				visited[j] = false
			}
		}
	}
}

func main() {
	in, _ := os.Open("195.in")
	defer in.Close()
	out, _ = os.Create("195.out")
	defer out.Close()

	var n int
	var line string
	for fmt.Fscanf(in, "%d", &n); n > 0; n-- {
		fmt.Fscanf(in, "%s", &line)
		chars = []byte(line)
		sort.Slice(chars, func(i, j int) bool {
			if isLower(chars[i]) && isLower(chars[j]) || isUpper(chars[i]) && isUpper(chars[j]) {
				return chars[i] <= chars[j]
			}
			if isLower(chars[i]) {
				return toUpper(chars[i]) <= chars[j]
			}
			return chars[i] <= toUpper(chars[j])
		})
		word := make([]byte, len(chars))
		visited := make([]bool, len(chars))
		dfs(0, word, visited)
	}
}
