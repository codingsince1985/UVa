// UVa 642 - Word Amalgamation

package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func split(word string) map[byte]int {
	charMap := make(map[byte]int)
	for i := range word {
		charMap[word[i]]++
	}
	return charMap
}

func find(word string, dict map[string]map[byte]int) []string {
	var matched []string
	charMap := split(word)
here:
	for i, vi := range dict {
		if len(charMap) == len(vi) {
			for j, vj := range charMap {
				if vi[j] != vj {
					continue here
				}
			}
			matched = append(matched, i)
		}
	}
	sort.Strings(matched)
	return matched
}

func main() {
	in, _ := os.Open("642.in")
	defer in.Close()
	out, _ := os.Create("642.out")
	defer out.Close()

	var word string
	dict := make(map[string]map[byte]int)
	for {
		if fmt.Fscanf(in, "%s", &word); word == "XXXXXX" {
			break
		}
		dict[word] = split(word)
	}
	for {
		if fmt.Fscanf(in, "%s", &word); word == "XXXXXX" {
			break
		}
		matched := find(word, dict)
		if len(matched) == 0 {
			fmt.Fprintln(out, "NOT A VALID WORD")
		} else {
			fmt.Fprintln(out, strings.Join(matched, "\n"))
		}
		fmt.Fprintln(out, "******")
	}
}
