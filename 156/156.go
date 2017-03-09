// UVa 156 - Ananagrams

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func sortWord(word string) string {
	var chs []int
	chars := []byte(strings.ToLower(word))
	for _, v := range chars {
		chs = append(chs, int(v))
	}
	sort.Ints(chs)
	chars = nil
	for _, v := range chs {
		chars = append(chars, byte(v))
	}
	return string(chars)
}

func main() {
	in, _ := os.Open("156.in")
	defer in.Close()
	out, _ := os.Create("156.out")
	defer out.Close()

	var word string
	sortWordMap := make(map[string][]string)
	for {
		if fmt.Fscanf(in, "%s", &word); word == "#" {
			break
		}
		sortWord(word)
		sortWordMap[sortWord(word)] = append(sortWordMap[sortWord(word)], word)
	}
	var ananagrams []string
	for _, v := range sortWordMap {
		if len(v) == 1 {
			ananagrams = append(ananagrams, v[0])
		}
	}
	sort.Strings(ananagrams)
	fmt.Fprintln(out, strings.Join(ananagrams, "\n"))
}
