// UVa 628 - Passwords

package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	out  io.WriteCloser
	dict []string
	rule string
)

func dfs(level int, tokens []string) {
	if level == len(rule) {
		fmt.Fprintln(out, strings.Join(tokens, ""))
		return
	}
	if rule[level] == '#' {
		for i := range dict {
			dfs(level+1, append(tokens, dict[i]))
		}
	} else {
		for i := '0'; i <= '9'; i++ {
			dfs(level+1, append(tokens, string(i)))
		}
	}
}

func main() {
	in, _ := os.Open("628.in")
	defer in.Close()
	out, _ = os.Create("628.out")
	defer out.Close()

	var nd, nr int
	for {
		if _, err := fmt.Fscanf(in, "%d", &nd); err != nil {
			break
		}
		dict = make([]string, nd)
		for i := range dict {
			fmt.Fscanf(in, "%s", &dict[i])
		}
		fmt.Fprintln(out, "--")
		for fmt.Fscanf(in, "%d", &nr); nr > 0; nr-- {
			fmt.Fscanf(in, "%s", &rule)
			dfs(0, nil)
		}
	}
}
