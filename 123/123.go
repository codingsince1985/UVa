// UVa 123 - Searching Quickly

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var out *os.File

func outputLine(title []string, idx int) {
	newTitle := make([]string, len(title))
	copy(newTitle, title)
	newTitle[idx] = strings.ToUpper(newTitle[idx])
	fmt.Fprintln(out, strings.Join(newTitle, " "))
}

func output(titles [][]string, keywordMap map[string]bool) {
	var keywords []string
	for keyword := range keywordMap {
		keywords = append(keywords, keyword)
	}
	sort.Strings(keywords)
	for _, keyword := range keywords {
		for _, title := range titles {
			for idx, word := range title {
				if keyword == word {
					outputLine(title, idx)
				}
			}
		}
	}
}

func main() {
	in, _ := os.Open("123.in")
	defer in.Close()
	out, _ = os.Create("123.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	ignores := make(map[string]bool)
	var word string
	for s.Scan() {
		if word = s.Text(); word == "::" {
			break
		}
		ignores[word] = true
	}

	keywordMap := make(map[string]bool)
	var titles [][]string
	for s.Scan() {
		r := strings.NewReader(s.Text())
		var title []string
		for {
			if _, err := fmt.Fscanf(r, "%s", &word); err != nil {
				break
			}
			word = strings.ToLower(word)
			title = append(title, word)
			if !ignores[word] {
				keywordMap[word] = true
			}
		}
		titles = append(titles, title)
	}
	output(titles, keywordMap)
}
