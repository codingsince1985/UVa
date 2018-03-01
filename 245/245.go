// UVa 245 - Uncompress

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var words []string

func isAlphabet(b byte) bool { return b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' }

func isDigit(b byte) bool { return b >= '0' && b <= '9' }

func moveToFront(word string, idx int) {
	if idx != -1 {
		words = append(words[:idx], words[idx+1:]...)
	}
	words = append([]string{word}, words...)
}

func push(word string) {
	idx := -1
	for i, w := range words {
		if w == word {
			idx = i
		}
	}
	moveToFront(word, idx)
}

func uncompress(pos int) string {
	moveToFront(words[pos], pos)
	return words[0]
}

func parse(line string) string {
	var uncompressed strings.Builder
	var word, pos string
	var idx int
	for i := range line {
		if isAlphabet(line[i]) {
			word += string(line[i])
		} else {
			if len(word) > 0 {
				uncompressed.WriteString(word)
				push(word)
				word = ""
			}
			if isDigit(line[i]) {
				pos += string(line[i])
			} else {
				if len(pos) > 0 {
					fmt.Sscanf(pos, "%d", &idx)
					uncompressed.WriteString(uncompress(idx - 1))
					pos = ""
				}
				uncompressed.WriteByte(line[i])
			}
		}
	}
	if len(word) > 0 {
		uncompressed.WriteString(word)
		push(word)
	}
	if len(pos) > 0 {
		fmt.Sscanf(pos, "%d", &idx)
		uncompressed.WriteString(uncompress(idx - 1))
	}
	return uncompressed.String()
}

func main() {
	in, _ := os.Open("245.in")
	defer in.Close()
	out, _ := os.Create("245.out")
	defer out.Close()

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)

	var line string
	for s.Scan() {
		if line = s.Text(); line == "0" {
			break
		}
		fmt.Fprintln(out, parse(line))
	}
}
