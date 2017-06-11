// UVa 11233 - Deli Deli

package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	irregulars = make(map[string]string)
	eses       = []string{"o", "s", "ch", "sh", "x"}
	vowelsY    = []string{"ay", "ey", "iy", "oy", "uy"}
)

func consonantAndY(word string) bool {
	for _, vowelY := range vowelsY {
		if strings.HasSuffix(word, vowelY) {
			return false
		}
	}
	return true
}

func solve(word string) string {
	if p, ok := irregulars[word]; ok {
		return p
	}
	if consonantAndY(word) {
		return word[:len(word)-1] + "ies"
	}
	for _, es := range eses {
		if strings.HasSuffix(word, es) {
			return word + "es"
		}
	}
	return word + "s"
}

func main() {
	in, _ := os.Open("11233.in")
	defer in.Close()
	out, _ := os.Create("11233.out")
	defer out.Close()

	var l, n int
	var w1, w2 string
	for fmt.Fscanf(in, "%d%d", &l, &n); l > 0; l-- {
		fmt.Fscanf(in, "%s%s", &w1, &w2)
		irregulars[w1] = w2
	}
	for ; n > 0; n-- {
		fmt.Fscanf(in, "%s", &w1)
		fmt.Fprintln(out, solve(w1))
	}
}
