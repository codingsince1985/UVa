// UVa 409 - Excuses, Excuses!

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isAlphabetic(b byte) bool { return b >= 'A' && b <= 'Z' || b >= 'a' && b <= 'z' }

func toLower(b byte) byte {
	if b >= 'A' && b <= 'Z' {
		return b - 'A' + 'a'
	}
	return b
}

func split(line string) []string {
	var words []string
	var word string
	var started bool
	for i := range line {
		if !started && isAlphabetic(line[i]) {
			started = true
		}
		if started {
			if !isAlphabetic(line[i]) {
				words = append(words, word)
				word = ""
				started = false
			} else {
				word += string(toLower(line[i]))
			}
		}
	}
	return words
}

func main() {
	in, _ := os.Open("409.in")
	defer in.Close()
	out, _ := os.Create("409.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var kase, k, e int
	for s.Scan() {
		r := strings.NewReader(s.Text())
		fmt.Fscanf(r, "%d%d", &k, &e)
		keywords := make(map[string]bool)
		for i := 0; i < k; i++ {
			s.Scan()
			keywords[s.Text()] = true
		}
		excuses := make([]string, e)
		scores := make([]int, e)
		var max int
		for i := range excuses {
			s.Scan()
			excuses[i] = s.Text()
			for _, word := range split(excuses[i]) {
				if keywords[word] {
					scores[i]++
				}
			}
			if scores[i] > max {
				max = scores[i]
			}
		}
		kase++
		fmt.Fprintf(out, "Excuse Set #%d\n", kase)
		for i, excuse := range excuses {
			if scores[i] == max {
				fmt.Fprintln(out, excuse)
			}
		}
		fmt.Fprintln(out)
	}
}
