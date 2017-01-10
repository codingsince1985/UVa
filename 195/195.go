// UVa 195 - Anagram

package main

import (
	"fmt"
	"os"
	"sort"
)

type bytes []byte

var (
	chars bytes
	out   *os.File
)

func isLower(a byte) bool { return a >= 'a' && a <= 'z' }

func isUpper(a byte) bool { return !isLower(a) }

func toUpper(a byte) byte { return a - 'a' + 'A' }

func (a bytes) Len() int { return len(a) }

func (a bytes) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a bytes) Less(i, j int) bool {
	if isLower(a[i]) && isLower(a[j]) || isUpper(a[i]) && isUpper(a[j]) {
		return a[i] <= a[j]
	}
	if isLower(a[i]) {
		return toUpper(a[i]) <= a[j]
	}
	return a[i] <= toUpper(a[j])
}

func dfs(i int, word bytes, visited []bool) {
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
		chars = bytes(line)
		sort.Sort(chars)
		word := make(bytes, len(chars))
		visited := make([]bool, len(chars))
		dfs(0, word, visited)
	}
}
