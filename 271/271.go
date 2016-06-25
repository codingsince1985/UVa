// UVa 271 - Simply Syntax

package main

import (
	"fmt"
	"os"
)

func wellFormed(s string) bool {
	var stack []string
	ok := true
here:
	for i := len(s) - 1; i >= 0; i-- {
		l := len(stack)
		switch {
		case s[i] >= 'p' && s[i] <= 'z':
			stack = append(stack, string(s[i]))
		case s[i] == 'N':
			if l < 1 {
				ok = false
				break here
			}
			stack[l-1] = "N" + stack[l-1]
		case s[i] == 'C' || s[i] == 'D' || s[i] == 'E' || s[i] == 'I':
			if l < 2 {
				ok = false
				break here
			}
			stack[l-2] = string(s[i]) + stack[l-1] + stack[l-2]
			stack = stack[:l-1]
		}
	}
	return ok && len(stack) == 1
}

func main() {
	in, _ := os.Open("271.in")
	defer in.Close()
	out, _ := os.Create("271.out")
	defer out.Close()

	var s string
	for {
		if _, err := fmt.Fscanf(in, "%s", &s); err != nil {
			break
		}
		if wellFormed(s) {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}
}
