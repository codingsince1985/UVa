// UVa 148 - Anagram checker

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	words   []string
	visited map[string]bool
	charMap map[byte]int
	paths   [][]string
)

func getMap(line string) map[byte]int {
	m := make(map[byte]int)
	for i := range line {
		if line[i] != ' ' {
			m[line[i]]++
		}
	}
	return m
}

func contains(word string) bool {
	cm := getMap(word)
	for k, v1 := range cm {
		if v2, ok := charMap[k]; !ok || v2 < v1 {
			return false
		}
	}
	return true
}

func remove(word string) {
	cm := getMap(word)
	for k, v := range cm {
		charMap[k] -= v
	}
}

func add(word string) {
	cm := getMap(word)
	for k, v := range cm {
		charMap[k] += v
	}
}

func full() bool {
	for _, v := range charMap {
		if v != 0 {
			return false
		}
	}
	return true
}

func newPath(path []string) bool {
here:
	for _, p := range paths {
		if len(p) == len(path) {
			for i := range path {
				if path[i] != p[i] {
					continue here
				}
			}
			return false
		}
	}
	return true
}

func dfs(path []string) {
	if full() {
		cp := make([]string, len(path))
		copy(cp, path)
		sort.Strings(cp)
		if newPath(cp) {
			paths = append(paths, cp)
		}
		return
	}
	for _, word := range words {
		if !visited[word] && contains(word) {
			visited[word] = true
			remove(word)
			dfs(append(path, word))
			visited[word] = false
			add(word)
		}
	}
}

func solve(line string) {
	charMap = getMap(line)
	visited = make(map[string]bool)
	paths = nil
	for _, word := range strings.Fields(line) {
		visited[word] = true
	}
	dfs(nil)
}

func main() {
	in, _ := os.Open("148.in")
	defer in.Close()
	out, _ := os.Create("148.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	for s.Scan() {
		if line = s.Text(); line == "#" {
			break
		}
		words = append(words, line)
	}
	for s.Scan() {
		if line = s.Text(); line == "#" {
			break
		}
		solve(line)
		for _, path := range paths {
			fmt.Fprintf(out, "%s = %s\n", line, strings.Join(path, " "))
		}
	}
}
