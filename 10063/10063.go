// UVa 10063 - Knuth's Permutation

package main

import (
	"fmt"
	"os"
	"strings"
)

func dfs(chars string) []string {
	if l := len(chars); l > 1 {
		var next []string
		for _, p := range dfs(chars[:l-1]) {
			for i := 0; i <= len(p); i++ {
				next = append(next, p[:i]+string(chars[l-1])+p[i:])
			}
		}
		return next
	}
	return []string{chars}
}

func main() {
	in, _ := os.Open("10063.in")
	defer in.Close()
	out, _ := os.Create("10063.out")
	defer out.Close()

	var chars string
	for first := true; ; {
		if _, err := fmt.Fscanf(in, "%s", &chars); err != nil {
			break
		}
		if first {
			first = false
		} else {
			fmt.Fprintln(out)
		}
		fmt.Fprintln(out, strings.Join(dfs(chars), "\n"))
	}
}
