// UVa 743 - The MTM Machine

package main

import (
	"fmt"
	"os"
	"strings"
)

func produce(s string) string {
	if s == "2" {
		return s
	}
	return s[1:]
}

func associate(s string) string {
	if s == "" {
		return s
	}
	return s + "2" + s
}

func solve(s string) string {
	switch {
	case strings.Contains(s, "0"):
		return ""
	case s[0] == '3':
		return associate(solve(s[1:]))
	case len(s) > 1 && s[0] == '2':
		return produce(s)
	default:
		return ""
	}
}

func main() {
	in, _ := os.Open("743.in")
	defer in.Close()
	out, _ := os.Create("743.out")
	defer out.Close()

	var n int64
	for {
		if fmt.Fscanf(in, "%d", &n); n == 0 {
			break
		}
		if s := solve(fmt.Sprint(n)); s != "" {
			fmt.Fprintln(out, s)
		} else {
			fmt.Fprintln(out, "NOT ACCEPTABLE")
		}
	}
}
