// UVa 10391 - Compound Words

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func binarySearch(w string, words []string) bool {
	return words[sort.Search(len(words), func(i int) bool { return words[i] >= w })] == w
}

func solve(words []string) []string {
	var matched []string
here:
	for _, word := range words {
		if len(word) > 1 {
			for i := 1; i < len(word); i++ {
				if w1, w2 := word[:i], word[i:]; binarySearch(w1, words) && binarySearch(w2, words) {
					matched = append(matched, word)
					continue here
				}
			}
		}
	}
	return matched
}

func main() {
	in, _ := os.Open("10391.in")
	defer in.Close()
	out, _ := os.Create("10391.out")
	defer out.Close()

	var words []string
	var word string
	for {
		if _, err := fmt.Fscanf(in, "%s", &word); err != nil {
			break
		}
		words = append(words, word)
	}
	fmt.Fprintln(out, strings.Join(solve(words), "\n"))
}
