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

	s := bufio.NewScanner(in)
	s.Split(bufio.ScanLines)
	for s.Scan() {
		words := strings.Split(s.Text(), " ")
		var newLine []string
		for _, v := range words {
			newLine = append(newLine, reverse(v))
		}
		fmt.Fprintln(out, strings.Join(newLine, " "))
	}
}
