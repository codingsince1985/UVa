// UVa 483 - Word Scramble

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reverse(word string) string {
	l := len(word)
	ret := make([]byte, l)
	for i := range ret {
		ret[i] = word[l-1-i]
	}
	return string(ret)
}

func main() {
	in, _ := os.Open("483.in")
	defer in.Close()
	out, _ := os.Create("483.out")
	defer out.Close()

	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := scanner.Text()
		lineReader := strings.NewReader(line)
		newLine := make([]string, 0)
		var word string
		for {
			if _, err := fmt.Fscanf(lineReader, "%s", &word); err != nil {
				break
			}
			newLine = append(newLine, reverse(word))
		}
		fmt.Fprintln(out, strings.Join(newLine, " "))
	}
}
