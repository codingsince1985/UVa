// UVa 10126 - Zipf's Law

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func isAlphabet(b byte) bool { return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' }

func parse(line string) []string {
	var b byte
	var word string
	var words []string
	for r := strings.NewReader(line); ; {
		if _, err := fmt.Fscanf(r, "%c", &b); err != nil {
			if word != "" {
				words = append(words, word)
			}
			break
		}
		if !isAlphabet(b) {
			if word != "" {
				words = append(words, word)
				word = ""
			}
		} else {
			word += string(b)
		}
	}
	return words
}

func count(words []string) map[string]int {
	m := make(map[string]int)
	for _, word := range words {
		m[strings.ToLower(word)]++
	}
	return m
}

func get(frequency map[string]int, n int) []string {
	var words []string
	for k, v := range frequency {
		if v == n {
			words = append(words, k)
		}
	}
	sort.Strings(words)
	return words
}

func main() {
	in, _ := os.Open("10126.in")
	defer in.Close()
	out, _ := os.Create("10126.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	first := true
	for s.Scan() {
		n, _ := strconv.Atoi(s.Text())
		frequency := make(map[string]int)
		for s.Scan() {
			if line = s.Text(); line == "EndOfText" {
				break
			}
			for k, v := range count(parse(line)) {
				frequency[k] += v
			}
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		if words := get(frequency, n); len(words) > 0 {
			fmt.Fprintln(out, strings.Join(words, "\n"))
		} else {
			fmt.Fprintln(out, "There is no such word.")
		}
	}
}
