// UVa 11541 - Decoding

package main

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

var (
	char   = func(c rune) bool { return !unicode.IsLetter(c) }
	number = func(c rune) bool { return !unicode.IsNumber(c) }
)

func decode(line string) []string {
	numbers := strings.FieldsFunc(line, number)
	num := make([]int, len(numbers))
	for i, n := range numbers {
		fmt.Sscanf(n, "%d", &num[i])
	}
	token := make([]string, len(num))
	for i, c := range strings.FieldsFunc(line, char) {
		token[i] = strings.Repeat(c, num[i])
	}
	return token
}

func main() {
	in, _ := os.Open("11541.in")
	defer in.Close()
	out, _ := os.Create("11541.out")
	defer out.Close()

	var line string
	var t int
	fmt.Fscanf(in, "%d", &t)
	for kase := 1; kase <= t; kase++ {
		fmt.Fscanf(in, "%s", &line)
		fmt.Fprintf(out, "Case %d: %s\n", kase, strings.Join(decode(line), ""))
	}
}
