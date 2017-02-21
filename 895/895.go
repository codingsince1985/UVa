// UVa 895 - Word Problem

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func buildMap(word string) map[byte]int {
	wm := make(map[byte]int)
	for i := range word {
		wm[word[i]]++
	}
	return wm
}

func solve(chs string, wordMap map[string]map[byte]int) int {
	charMap := buildMap(chs)
	var count int
here:
	for _, wm := range wordMap {
		for c, n := range wm {
			if charMap[c] < n {
				continue here
			}
		}
		count++
	}
	return count
}

func main() {
	in, _ := os.Open("895.in")
	defer in.Close()
	out, _ := os.Create("895.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	wordMap := make(map[string]map[byte]int)
	for s.Scan() {
		if line = s.Text(); line == "#" {
			break
		}
		wordMap[line] = buildMap(line)
	}
	for s.Scan() {
		if line = s.Text(); line == "#" {
			break
		}
		fmt.Fprintln(out, solve(strings.Replace(line, " ", "", -1), wordMap))
	}
}
