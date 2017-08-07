// UVa 11192 - Group Reverse

package main

import (
	"fmt"
	"os"
	"strings"
)

func reverse(tokens []string) []string {
	reversed := make([]string, len(tokens))
	for i, token := range tokens {
		for j := range token {
			reversed[i] = string(token[j]) + reversed[i]
		}
	}
	return reversed
}

func split(line string, n int) []string {
	length := len(line) / n
	tokens := make([]string, n)
	for i := range tokens {
		tokens[i] = line[i*length : (i+1)*length]
	}
	return tokens
}

func main() {
	in, _ := os.Open("11192.in")
	defer in.Close()
	out, _ := os.Create("11192.out")
	defer out.Close()

	var n int
	var line string
	for {
		if _, err := fmt.Fscanf(in, "%d%s", &n, &line); err != nil || n == 0 {
			break
		}
		fmt.Fprintln(out, strings.Join(reverse(split(line, n)), ""))
	}
}
