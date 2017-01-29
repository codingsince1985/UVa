// UVa 10815 - Andy's First Dictionary

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func isAlphabet(b byte) bool { return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' }

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b - 'A' + 'a'
	}
	return b
}

func main() {
	in, _ := os.Open("10815.in")
	defer in.Close()
	out, _ := os.Create("10815.out")
	defer out.Close()

	var b byte
	var word []byte
	wordMap := make(map[string]bool)
	var started bool
	for {
		if _, err := fmt.Fscanf(in, "%c", &b); err != nil {
			break
		}
		switch {
		case isAlphabet(b):
			if !started {
				word = nil
				started = true
			}
			word = append(word, toLower(b))
		case started:
			wordMap[string(word)] = true
			started = false
		}
	}
	var words []string
	for k := range wordMap {
		words = append(words, k)
	}
	sort.Strings(words)
	fmt.Fprintln(out, strings.Join(words, "\n"))
}
