// UVa 454 - Anagrams

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var keepLetter = func(r rune) rune {
	if r >= rune('a') && r <= rune('z') || r >= rune('A') && r <= rune('Z') {
		return r
	}
	return -1
}

func solve(out *os.File, words []string) {
	wordMap := make(map[string]string)
	for _, word := range words {
		s := strings.Split(strings.Map(keepLetter, word), "")
		sort.Strings(s)
		wordMap[word] = strings.Join(s, "")
	}

	sort.Strings(words)
	for i := 0; i < len(words)-1; i++ {
		for j := i + 1; j < len(words); j++ {
			if wordMap[words[i]] == wordMap[words[j]] {
				fmt.Fprintf(out, "%s = %s\n", words[i], words[j])
			}
		}
	}
}

func main() {
	in, _ := os.Open("454.in")
	defer in.Close()
	out, _ := os.Create("454.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	s.Scan()
	kase, _ := strconv.Atoi(s.Text())
	s.Scan()
	s.Text()
	var words []string
	var line string
	first := true
	for ; kase > 0; kase-- {
		for s.Scan() {
			if line = s.Text(); line == "" {
				break
			}
			words = append(words, line)
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		solve(out, words)
		words = nil
	}
}
