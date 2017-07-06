// UVa 10393 - The One-Handed Typist

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

var fingerMap = map[int]map[byte]bool{
	1:  {'q': true, 'a': true, 'z': true},
	2:  {'w': true, 's': true, 'x': true},
	3:  {'e': true, 'd': true, 'c': true},
	4:  {'r': true, 'f': true, 'v': true, 't': true, 'g': true, 'b': true},
	5:  {},
	6:  {},
	7:  {'y': true, 'h': true, 'n': true, 'u': true, 'j': true, 'm': true},
	8:  {'i': true, 'k': true},
	9:  {'o': true, 'l': true},
	10: {'p': true},
}

func disabled(fingers []int) map[byte]bool {
	chars := make(map[byte]bool)
	for _, finger := range fingers {
		for k := range fingerMap[finger] {
			chars[k] = true
		}
	}
	return chars
}

func solve(fingers []int, words []string) []string {
	chars := disabled(fingers)
	var longest []string
	var max int
here:
	for _, word := range words {
		if len(word) < max {
			continue
		}
		for i := range word {
			if chars[word[i]] {
				continue here
			}
		}
		if len(word) > max {
			max = len(word)
			longest = nil
		}
		longest = append(longest, word)
	}
	sort.Strings(longest)
	return longest
}

func main() {
	in, _ := os.Open("10393.in")
	defer in.Close()
	out, _ := os.Create("10393.out")
	defer out.Close()

	var f, n int
	for {
		if _, err := fmt.Fscanf(in, "%d%d", &f, &n); err != nil {
			break
		}
		fingers := make([]int, f)
		for i := range fingers {
			fmt.Fscanf(in, "%d", &fingers[i])
		}
		words := make([]string, n)
		for i := range words {
			fmt.Fscanf(in, "%s", &words[i])
		}
		longest := solve(fingers, words)
		fmt.Fprintln(out, len(longest))
		fmt.Fprintln(out, strings.Join(longest, "\n"))
	}
}
