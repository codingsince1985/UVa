// UVa 630 - Anagrams (II)

package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

var vocabulary [][2]string

func solve(out io.Writer, word string) {
	fmt.Fprintf(out, "Anagrams for: %s\n", word)
	var count int
	sw := sortWord(word)
	for _, w := range vocabulary {
		if sw == w[1] {
			count++
			fmt.Fprintf(out, "%3d) %s\n", count, w[0])
		}
	}
	if count == 0 {
		fmt.Fprintf(out, "No anagrams for: %s\n", word)
	}
}

func sortWord(word string) string {
	bytes := strings.Split(word, "")
	sort.Strings(bytes)
	return strings.Join(bytes, "")
}

func main() {
	in, _ := os.Open("630.in")
	defer in.Close()
	out, _ := os.Create("630.out")
	defer out.Close()

	var t, n int
	var word string
	first := true
	for fmt.Fscanf(in, "%d", &t); t > 0; t-- {
		fmt.Fscanf(in, "\n%d", &n)
		vocabulary = make([][2]string, n)
		for i := range vocabulary {
			fmt.Fscanf(in, "%s", &vocabulary[i][0])
			vocabulary[i][1] = sortWord(vocabulary[i][0])
		}
		sort.Slice(vocabulary, func(i, j int) bool { return vocabulary[i][0] < vocabulary[j][0] })
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		for {
			if fmt.Fscanf(in, "%s", &word); word == "END" {
				break
			}
			solve(out, word)
		}
	}
}
