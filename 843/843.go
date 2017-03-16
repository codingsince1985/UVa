// UVa 843 - Crypt Kicker

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var (
	out               *os.File
	dictionary, words []string
)

func sortWords(words []string) []string {
	sort.Slice(words, func(i, j int) bool {
		l1, l2 := len(words[i]), len(words[j])
		if l1 == l2 {
			return distinct(words[i]) > distinct(words[j])
		}
		return l1 > l2
	})
	return words
}

func distinct(word string) int {
	m := make(map[byte]bool)
	for i := range word {
		m[word[i]] = true
	}
	return len(m)
}

func getMap(w1, w2 string) (bool, map[byte]byte, map[byte]byte) {
	if len(w1) != len(w2) {
		return false, nil, nil
	}
	m1, m2 := make(map[byte]byte), make(map[byte]byte)
	for i := range w1 {
		if c, ok := m1[w2[i]]; ok && c != w1[i] {
			return false, nil, nil
		}
		if c, ok := m2[w1[i]]; ok && c != w2[i] {
			return false, nil, nil
		}
		m1[w2[i]] = w1[i]
		m2[w1[i]] = w2[i]
	}
	return true, m1, m2
}

func noConflict(m, nm map[byte]byte) bool {
	for k, v := range nm {
		if c, ok := m[k]; ok && c != v {
			return false
		}
	}
	return true
}

func decode(words []string, m map[byte]byte) bool {
	for _, word := range words {
		for i := range word {
			if _, ok := m[word[i]]; !ok {
				return false
			}
		}
	}
	return true
}

func backtracking(level int, m1, m2 map[byte]byte) map[byte]byte {
	if level == len(dictionary) {
		return m1
	}
	for _, word := range words {
		if ok, nm1, nm2 := getMap(dictionary[level], word); ok {
			if noConflict(m1, nm1) && noConflict(m2, nm2) {
				diff1 := make(map[byte]byte)
				for k, v := range nm1 {
					if _, ok := m1[k]; !ok {
						diff1[k] = v
						m1[k] = v
					}
				}
				diff2 := make(map[byte]byte)
				for k, v := range nm2 {
					if _, ok := m2[k]; !ok {
						diff2[k] = v
						m2[k] = v
					}
				}
				if decode(words, m1) {
					return m1
				}
				if m := backtracking(level+1, m1, m2); m != nil {
					return m
				}
				for k := range diff1 {
					delete(m1, k)
				}
				for k := range diff2 {
					delete(m2, k)
				}
			}
		}
	}
	return nil
}

func solve(line string) {
	sortWords(dictionary)
	words = strings.Split(line, " ")
	sortWords(words)
	m := backtracking(0, make(map[byte]byte), make(map[byte]byte))
	for i := range line {
		if line[i] != ' ' {
			if c, ok := m[line[i]]; ok {
				fmt.Fprintf(out, "%c", c)
			} else {
				fmt.Fprint(out, "*")
			}
		} else {
			fmt.Fprint(out, " ")
		}
	}
	fmt.Fprintln(out)
}

func main() {
	in, _ := os.Open("843.in")
	defer in.Close()
	out, _ = os.Create("843.out")
	defer out.Close()

	var n int
	fmt.Fscanf(in, "%d", &n)
	dictionary = make([]string, n)
	for i := range dictionary {
		fmt.Fscanf(in, "%s", &dictionary[i])
	}

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		solve(s.Text())
	}
}
